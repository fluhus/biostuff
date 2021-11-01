// Package sam parses SAM files.
//
// This package uses the format described in:
// https://en.wikipedia.org/wiki/SAM_(file_format)
package sam

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"sort"
	"strconv"

	"github.com/fluhus/gostuff/csvdec"
)

// TODO(amit): Add writing.

// A raw structure for the initial parsing using csvdec.
type samRaw struct {
	Qname string
	Flag  int
	Rname string
	Pos   int
	Mapq  int
	Cigar string
	Rnext string
	Pnext int
	Tlen  int
	Seq   string
	Qual  string
	Extra []string
}

// SAM is a single line (alignment) in a SAM file.
type SAM struct {
	Qname string                 // Query name
	Flag  int                    // Bitwise flag
	Rname string                 // Reference sequence name
	Pos   int                    // Mapping position (1-based)
	Mapq  int                    // Mapping quality
	Cigar string                 // CIGAR string
	Rnext string                 // Ref. name of the mate/next read
	Pnext int                    // Position of the mate/next read
	Tlen  int                    // Observed template length
	Seq   string                 // Sequence
	Qual  string                 // Phred qualities (ASCII)
	Tags  map[string]interface{} // Typed optional tags.
}

// MarshalText returns the textual representation of s in SAM format.
// Includes a trailing new line.
func (s *SAM) MarshalText() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	fmt.Fprintf(buf, "%s\t%d\t%s\t%d\t%d\t%s\t%s\t%d\t%d\t%s\t%s",
		s.Qname, s.Flag, s.Rname, s.Pos,
		s.Mapq, s.Cigar, s.Rnext, s.Pnext,
		s.Tlen, s.Seq, s.Qual,
	)
	for _, tag := range tagsToText(s.Tags) {
		fmt.Fprintf(buf, "\t%s", tag)
	}
	buf.WriteByte('\n')
	return buf.Bytes(), nil
}

// Converts a raw SAM struct to an exported SAM struct.
func fromRaw(raw *samRaw) (*SAM, error) {
	result := &SAM{}
	result.Qname = raw.Qname
	result.Flag = raw.Flag
	result.Rname = raw.Rname
	result.Pos = raw.Pos
	result.Mapq = raw.Mapq
	result.Cigar = raw.Cigar
	result.Rnext = raw.Rnext
	result.Pnext = raw.Pnext
	result.Tlen = raw.Tlen
	result.Seq = raw.Seq
	result.Qual = raw.Qual
	result.Tags = map[string]interface{}{}
	var err error
	result.Tags, err = parseTags(raw.Extra)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// A Reader reads and parses SAM lines.
type Reader struct {
	r *bufio.Reader
	d *csvdec.Decoder
	h bool // Indicates that we are done reading the header.
}

// NewReader returns a new SAM reader that reads from r.
func NewReader(r io.Reader) *Reader {
	var b *bufio.Reader
	switch r := r.(type) {
	case *bufio.Reader:
		b = r
	default:
		b = bufio.NewReader(r)
	}
	d := csvdec.NewDecoder(b)
	d.Comma = '\t'
	d.FieldsPerRecord = -1 // Allow variable number of fields.
	return &Reader{b, d, false}
}

// NextHeader returns the next header line as a raw string, including the '@'.
// Returns EOF when out of header lines, then Next can be called for the
// data lines.
func (r *Reader) NextHeader() (string, error) {
	if r.h {
		panic("Cannot read header after reading alignments.")
	}
	b, err := r.r.ReadByte()
	if err != nil {
		return "", err
	}
	if b != '@' {
		r.r.UnreadByte()
		r.h = true
		return "", io.EOF
	}
	line, err := r.r.ReadBytes('\n')
	if err != nil {
		if err == io.EOF {
			if len(line) == 0 {
				return "", io.ErrUnexpectedEOF
			}
		} else {
			return "", err
		}
	}
	line = line[:len(line)-1]
	if len(line) == 0 {
		return "", fmt.Errorf("encountered an empty header line")
	}
	return "@" + string(line), nil
}

// Read returns the next SAM line.
func (r *Reader) Read() (*SAM, error) {
	for !r.h {
		r.NextHeader()
	}
	raw := &samRaw{}
	err := r.d.Decode(raw)
	if err != nil {
		return nil, err
	}
	return fromRaw(raw)
}

// Returns a map from tag name to its parsed (typed) value.
func parseTags(values []string) (map[string]interface{}, error) {
	result := make(map[string]interface{}, len(values))
	for _, f := range values {
		parts, err := splitTag(f)
		if err != nil {
			return nil, err
		}
		switch parts[1] {
		case "A":
			if len(parts[2]) != 1 {
				return nil, fmt.Errorf("illegal value for tag type %v: %q, "+
					"want a single character",
					parts[1], parts[2])
			}
			result[parts[0]] = parts[2][0]
		case "i":
			x, err := strconv.Atoi(parts[2])
			if err != nil {
				return nil, fmt.Errorf("illegal value for tag type %v: %q, "+
					"want an integer",
					parts[1], parts[2])
			}
			result[parts[0]] = x
		case "f":
			x, err := strconv.ParseFloat(parts[2], 64)
			if err != nil {
				return nil, fmt.Errorf("illegal value for tag type %v: %q, "+
					"want a number",
					parts[1], parts[2])
			}
			result[parts[0]] = x
		case "Z":
			result[parts[0]] = parts[2]
		case "H":
			x, err := hex.DecodeString(parts[2])
			if err != nil {
				return nil, fmt.Errorf("illegal value for tag type %v: %q, "+
					"want a hexadecimal sequence",
					parts[1], parts[2])
			}
			result[parts[0]] = x
		case "B":
			// TODO(amit): Not implemented yet. Treating like string for now.
			result[parts[0]] = parts[2]
		default:
			return nil, fmt.Errorf("unrecognized tag type: %v, in tag %v",
				parts[1], f)
		}
	}
	return result, nil
}

// Splits a SAM tag by colon. Used instead of strings.SpliN for performance.
func splitTag(tag string) ([3]string, error) {
	colon1, colon2 := -1, -1
	for i, c := range tag {
		if c == ':' {
			if colon1 == -1 {
				colon1 = i
			} else {
				colon2 = i
				break
			}
		}
	}
	var result [3]string
	if colon2 == -1 {
		return result, fmt.Errorf("tag doesn't have at least 3 colons: %q", tag)
	}
	result[0] = tag[:colon1]
	result[1] = tag[colon1+1 : colon2]
	result[2] = tag[colon2+1:]
	return result, nil
}

// Returns the given tags in SAM format, sorted and tab-separated.
func tagsToText(tags map[string]interface{}) []string {
	texts := make([]string, 0, len(tags))
	for tag, val := range tags {
		texts = append(texts, tagToText(tag, val))
	}
	sort.Strings(texts)
	return texts
}

// Returns the SAM format representation of the given tag.
func tagToText(tag string, val interface{}) string {
	switch val := val.(type) {
	case byte:
		return tag + ":A:" + strconv.Itoa(int(val))
	case int:
		return tag + ":i:" + strconv.Itoa(val)
	case float64:
		return tag + ":f:" + strconv.FormatFloat(val, 'e', -1, 64)
	case string:
		return tag + ":Z:" + val
	case []byte:
		return tag + ":H:" + hex.EncodeToString(val)
	default:
		panic(fmt.Sprintf("unsupported type for value %v", val))
	}
}
