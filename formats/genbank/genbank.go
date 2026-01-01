// Package genbank provides parsing for GenBank data.
//
// Based on the documentation found here:
// https://www.ncbi.nlm.nih.gov/genbank/samplerecord/
//
// # A word of caution
//
// The GenBank format is more intuitive than it is formal.
// Meaning, it is easy to understand for a human looking at it,
// but it is difficult to come up with one set of rules
// for a parser to follow, due to inconsistency between different entries
// and even between fields in the same entry.
//
// Since each line needs to be treated according to what "makes sense"
// at that particular instance, a parser needs to do its best
// to mimic human intuition about that data.
// This is prone to errors.
// Therefore, try to check that the output of this parser makes sense
// on your specific input,
// and please report any mistakes or problems on the [GitHub page].
//
// [GitHub page]: https://github.com/fluhus/biostuff/issues
package genbank

import (
	"fmt"
	"io"
	"iter"
	"regexp"
	"strings"

	"github.com/fluhus/gostuff/iterx"
	"github.com/fluhus/gostuff/snm"
)

const (
	// Use specialized splitting functions rather than regex,
	// to improve performance.
	useSplitFunc = true

	// Use new specialized splitting logic to further reduce
	// regex usage.
	useSuperSplitFunc = true

	// Ignore unknown fields, rather than returning an error.
	tolerateUnknownFields = true
)

// Reader iterates over GenBank entries in a reader.
func Reader(r io.Reader) iter.Seq2[*GenBank, error] {
	return func(yield func(*GenBank, error) bool) {
		i := 0
		for e := range splitToEntries(removeEmpty(iterx.LinesReader(r), &i)) {
			et, err := linesToEntry(e)
			if !yield(et, withLineNumber(err, i)) {
				return
			}
		}
	}
}

// File iterates over GenBank entries in a file.
func File(file string) iter.Seq2[*GenBank, error] {
	return func(yield func(*GenBank, error) bool) {
		i := 0
		for e := range splitToEntries(removeEmpty(iterx.LinesFile(file), &i)) {
			et, err := linesToEntry(e)
			if !yield(et, withLineNumber(err, i)) {
				return
			}
		}
	}
}

// Splits a line iterator into separate line iterators,
// each iterates over the lines of a single entry.
func splitToEntries(lines iter.Seq2[string, error]) iter.Seq[iter.Seq2[string, error]] {
	return func(yield func(iter.Seq2[string, error]) bool) {
		pull, stop := iter.Pull2(lines)
		defer stop()
		done := false
		for {
			// Check that there are lines before yielding a new iterator.
			first, firstErr, ok := pull()
			if !ok {
				return
			}
			if !yield(func(yield func(string, error) bool) {
				if !yield(first, firstErr) {
					return
				}
				for {
					line, err, ok := pull()
					if !ok { // EOF
						done = true
						return
					}
					if line == "//" { // End of entry
						return
					}
					if !yield(line, err) { // Entry data
						done = true
						return
					}
				}
			}) {
				return
			}
			if done {
				return
			}
		}
	}
}

// Returns an iterator over the same lines, excluding the empty ones.
func removeEmpty(lines iter.Seq2[string, error], i *int) iter.Seq2[string, error] {
	return func(yield func(string, error) bool) {
		for s, err := range lines {
			*i++
			if err == nil && s == "" {
				continue
			}
			if !yield(s, err) {
				return
			}
		}
	}
}

// Parses the lines of a single entry.
func linesToEntry(lines iter.Seq2[string, error]) (*GenBank, error) {
	e := &GenBank{}
	origin := &strings.Builder{}
	comment := &strings.Builder{}
	var curField, curRefField, curFeatureField string
	var m, mff []string

	for line, err := range lines {
		if err != nil {
			return nil, err
		}

		if useSplitFunc {
			m = splitLine(line, m)
		} else {
			m = lineRE.FindStringSubmatch(line)
		}
		if m == nil {
			return nil, fmt.Errorf("could not parse line: %q", line)
		}

		if m[1] != "" {
			curField = m[1]
		}

		switch m[1] {
		case "LOCUS":
			// TODO(amit): Check if field is already populated?
			// if e.Locus != "" {
			// 	return nil, fmt.Errorf("found two locus: %q, %q", line, e.Locus)
			// }
			e.Locus = m[2]
		case "DEFINITION":
			e.Definition = m[2]
		case "ACCESSION":
			e.Accessions = strings.Split(m[2], " ")
		case "VERSION":
			e.Version = m[2]
		case "DBLINK":
			e.DBLink = append(e.DBLink, m[2])
		case "KEYWORDS":
			e.Keywords = m[2]
		case "SOURCE":
			e.Source = m[2]
		case "  ORGANISM":
			e.Organism = m[2]
		case "REFERENCE":
			curRefField = ""
			ref := map[string]string{}
			if m[2] != "" {
				ref[""] = m[2]
			}
			e.References = append(e.References, ref)
		case "FEATURES":
			// Just change current field.
		case "ORIGIN":
			// Just change current field.
		case "COMMENT":
			comment.WriteString(m[2])
		case "CONTIG":
			// Ignore
			// TODO(amit): Figure this out.
		case "":
			switch curField {
			case "LOCUS":
				e.Locus += " " + m[2]
			case "DEFINITION":
				e.Definition += " " + m[2]
			case "ACCESSION":
				e.Accessions = append(e.Accessions,
					strings.Split(m[2], " ")...)
			case "VERSION":
				e.Version += " " + m[2]
			case "DBLINK":
				e.DBLink = append(e.DBLink, m[2])
			case "KEYWORDS":
				e.Keywords += " " + m[2]
			case "SOURCE":
				e.Source += " " + m[2]
			case "  ORGANISM":
				if e.OrganismTax == "" {
					e.OrganismTax = m[2]
				} else {
					e.OrganismTax += " " + m[2]
				}
			case "REFERENCE":
				ref := e.References[len(e.References)-1]
				if mref := refRE.FindStringSubmatch(line); mref != nil {
					curRefField = mref[1]
					ref[curRefField] = mref[2]
				} else {
					ref[curRefField] += " " + m[2]
				}
			case "FEATURES":
				if mf := featureRE.FindStringSubmatch(line); mf != nil {
					// New feature.
					f := &Feature{Type: mf[1], rawFields: newSBMap()}
					if mf[2] != "" {
						f.rawFields.write("", mf[2])
					}
					e.Features = append(e.Features, f)
					curFeatureField = ""
				} else {
					// Existing feature.
					f := e.Features[len(e.Features)-1]
					if useSplitFunc {
						mff = splitFeatureField(line, mff)
					} else {
						mff = featureFieldRE.FindStringSubmatch(line)
					}
					if mff != nil {
						curFeatureField = mff[1]
						f.rawFields.write(curFeatureField, mff[2])
					} else {
						f.rawFields.write(curFeatureField, " ", m[2])
					}
				}
			case "ORIGIN":
				seq := seqRE.ReplaceAllString(line, "")
				if seq == "" {
					return nil, fmt.Errorf("bad sequence line: %q", line)
				}
				origin.WriteString(seq)
			case "COMMENT":
				comment.WriteByte('\n')
				comment.WriteString(m[2])
			case "CONTIG":
				// Ignore
				// TODO(amit): Figure this out.
			default:
				if !tolerateUnknownFields {
					return nil, fmt.Errorf("bad line under %q: %q",
						curField, line)
				}
			}
		default:
			if !tolerateUnknownFields {
				return nil, fmt.Errorf("bad line: %q", line)
			}
		}
	}

	e.Origin = origin.String()
	e.Comment = comment.String()

	for _, f := range e.Features {
		f.convertRawFields()
		for k := range f.Fields {
			// Some feature values are wrapped in quotes, remove them.
			f.Fields[k] = strings.Trim(f.Fields[k], "\"")
			// Translation is special, it is not a sentence so no need
			// for the added spaces between lines.
			if k == "translation" {
				f.Fields[k] = strings.ReplaceAll(f.Fields[k], " ", "")
			}
		}
	}

	// TODO(amit): Remove redundant point?
	// e.OrganismTax = strings.TrimRight(e.OrganismTax, ".")

	// TODO(amit): Check for empty fields?

	return e, nil
}

// GenBank is a single GenBank entry.
type GenBank struct {
	Locus       string
	Definition  string
	Accessions  []string
	Version     string
	DBLink      []string `json:",omitempty"`
	Keywords    string
	Source      string
	Organism    string
	OrganismTax string
	References  []map[string]string `json:",omitempty"`
	Features    []*Feature          `json:",omitempty"`
	Origin      string              `json:",omitempty"`
	Comment     string              `json:",omitempty"`
}

// Feature is an entry under FEATURES.
type Feature struct {
	Type      string            // source, CDS, gene...
	Fields    map[string]string // Field name without '/' to value.
	rawFields *sbMap            // Maps field name to value builder.
}

func (f *Feature) convertRawFields() {
	if f.rawFields == nil {
		panic("attempt to convert with nil raw")
	}
	if f.Fields != nil {
		panic("attempt to convert with non-nil Fields")
	}
	f.Fields = f.rawFields.strings()
	f.rawFields = nil
}

// Maps field name to a string builder builder.
type sbMap struct {
	m map[string]*strings.Builder
}

// Returns a new map.
func newSBMap() *sbMap {
	return &sbMap{map[string]*strings.Builder{}}
}

// Appends the strings in v to the key k.
func (m *sbMap) write(k string, v ...string) {
	s := m.m[k]
	if s == nil {
		s = &strings.Builder{}
		m.m[k] = s
	}
	for _, x := range v {
		s.WriteString(x)
	}
}

// Converts from string builders to a map of strings.
func (m *sbMap) strings() map[string]string {
	return snm.MapToMap(m.m, func(k string, v *strings.Builder) (string, string) {
		return k, v.String()
	})
}

var (
	lineRE         *regexp.Regexp
	refRE          = regexp.MustCompile(`^   ?(\S+)\s+(.*)`)
	featureRE      = regexp.MustCompile(`^     (\S+)\s+(.*)`)
	featureFieldRE = regexp.MustCompile(`^     \s+/(\w+)(?:=(.*))?`)
	seqRE          = regexp.MustCompile(`[\d\s]+`)
)

// Populates lineRE with expected field names.
func init() {
	prefixes := []string{"  ORGANISM", "\\S+"}
	re := []byte("^(")
	for _, p := range prefixes {
		re = append(re, p...)
		// There is an extra | at the end which captures the empty string.
		re = append(re, '|')
	}
	re = append(re, `)(?:\s+(.*))?`...)
	lineRE = regexp.MustCompile(string(re))
}

// Adds line number if the error is non-nil.
func withLineNumber(err error, i int) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("line %d: %w", i, err)
}

// Splits a line into a field name and a value,
// the way lineRE would.
func splitLine(line string, reuse []string) []string {
	if useSuperSplitFunc {
		if line == "" {
			return append(reuse, "", "", "")
		}

		// Extract prefix.
		prefix := line
		if firstWordIs(line, "  ORGANISM") {
			prefix = line[:10]
		} else {
			for i, c := range line {
				if isWhitespace(c) {
					prefix = line[:i]
					break
				}
			}
		}

		// Extract suffix.
		suffix := ""
		for i, c := range line[len(prefix):] {
			if !isWhitespace(c) {
				suffix = line[len(prefix)+i:]
				break
			}
		}
		return append(reuse[:0], "", prefix, suffix)
	}
	if !lineRE.MatchString(line) {
		return nil
	}
	var minSpaceI int
	if strings.HasPrefix(line, "  ORGANISM") {
		minSpaceI = 2
	}
	preSpace := true
	var s1, s2 string
	for i, c := range line {
		if preSpace && i >= minSpaceI && isWhitespace(c) {
			s1 = line[:i]
			preSpace = false
		}
		if !preSpace && !isWhitespace(c) {
			s2 = line[i:]
			break
		}
	}
	if preSpace {
		s1 = line
	}
	return append(reuse[:0], "", s1, s2)
}

// Splits a feature-field line into a field name and a value,
// the way featureFieldRE would.
func splitFeatureField(line string, reuse []string) []string {
	if useSuperSplitFunc {
		// Remove first whitespaces.
		if !strings.HasPrefix(line, "      ") {
			return nil
		}
		from := 0
		for i, c := range line {
			if isWhitespace(c) {
				from = i + 1
			} else {
				break
			}
		}
		line = line[from:]

		// Feature has to start with '/'.
		if !strings.HasPrefix(line, "/") {
			return nil
		}
		line = line[1:]

		// Extract feature name.
		upto := 0
		for i, c := range line {
			if wordChars[c] {
				upto = i + 1
			} else {
				break
			}
		}
		prefix := line[:upto]
		suffix := line[upto:]
		if suffix == "" {
			return append(reuse[:0], "", prefix, "")
		}
		if suffix[0] != '=' {
			return nil
		}
		return append(reuse[:0], "", prefix, suffix[1:])
	}

	if !featureFieldRE.MatchString(line) {
		return nil
	}
	const (
		preWord = iota
		inWord
		postWord
	)
	state := preWord
	var i1, i2, i3, i4 int
	line = line[4:]
	for i, c := range line {
		switch state {
		case preWord:
			if !isWhitespace(c) {
				if c != '/' {
					// Should not happen if matched regex.
					panic(fmt.Sprintf("bad feature line: %q", line))
				}
				state = inWord
				i1, i2 = i+1, i+1
			}
		case inWord:
			if wordChars[c] {
				i2 = i + 1
			} else {
				if c != '=' {
					// I am not sure what to do here.
					// Some files don't have the '=' consistently but some do,
					// so not sure when/how to enforce this.
					if false {
						// Should not happen if matched regex.
						panic(fmt.Sprintf(
							"bad feature line: wanted '=' but found %q: %q",
							c, line))
					}
					return nil
				}
				state = postWord
				i3, i4 = i+1, i+1
			}
		case postWord:
			i4 = i + 1
		}
	}
	s1, s2 := line[i1:i2], line[i3:i4]
	return append(reuse[:0], "", s1, s2)
}

type char interface {
	byte | rune
}

// Returns whether x is a space or a tab.
func isWhitespace[C char](x C) bool {
	return x == ' ' || x == '\t'
}

// Characters that match \w regex.
var wordChars = initWordChars()

// Creates a slice with true for word characters (\w)
// and false for anything else.
func initWordChars() []bool {
	b := make([]bool, 256)
	for _, rng := range []string{"az", "AZ", "09", "__"} {
		for i := rng[0]; i <= rng[1]; i++ {
			b[i] = true
		}
	}
	return b
}

// Checks that s starts with the word, followed by a whitespace
// or end of string.
func firstWordIs(s, word string) bool {
	if !strings.HasPrefix(s, word) {
		return false
	}
	if len(s) == len(word) {
		return true
	}
	return isWhitespace(s[len(word)])
}
