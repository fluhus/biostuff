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
	var curField, curRefField, curFeatureField string

	for line, err := range lines {
		if err != nil {
			return nil, err
		}

		m := lineRE.FindStringSubmatch(line)
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
					f := Feature{Type: mf[1], Fields: map[string]string{}}
					if mf[2] != "" {
						f.Fields[""] = mf[2]
					}
					e.Features = append(e.Features, f)
					curFeatureField = ""
				} else {
					// Existing feature.
					f := e.Features[len(e.Features)-1]
					if mf := featureFieldRE.FindStringSubmatch(line); mf != nil {
						curFeatureField = mf[1]
						f.Fields[curFeatureField] = mf[2]
					} else {
						f.Fields[curFeatureField] += " " + m[2]
					}
				}
			case "ORIGIN":
				seq := seqRE.ReplaceAllString(line, "")
				if seq == "" {
					return nil, fmt.Errorf("bad sequence line: %q", line)
				}
				origin.WriteString(seq)
			default:
				return nil, fmt.Errorf("bad line: %s", line)
			}
		}
	}

	e.Origin = origin.String()

	for _, f := range e.Features {
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
	Features    []Feature           `json:",omitempty"`
	Origin      string              `json:",omitempty"`
}

// Feature is an entry under FEATURES.
type Feature struct {
	Type   string            // source, CDS, gene...
	Fields map[string]string // Field name without '/' to value.
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
	prefixes := []string{
		"LOCUS", "DEFINITION", "ACCESSION", "VERSION", "DBLINK", "KEYWORDS",
		"SOURCE", "  ORGANISM", "REFERENCE", "FEATURES", "ORIGIN",
	}
	re := []byte("^(")
	for _, p := range prefixes {
		re = append(re, regexp.QuoteMeta(p)...)
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
