// Package sam handles SAM files.
package sam

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"strings"

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

// SAM represents a single SAM line.
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

// Reader reads and parses SAM lines.
type Reader struct {
	r *bufio.Reader
	d *csvdec.Decoder
	h bool // Indicates that we are done reading the header.
}

// NewReader returns a new SAM reader.
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

// Next returns the next SAM line.
func (r *Reader) Next() (*SAM, error) {
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
	result := map[string]interface{}{}
	for _, f := range values {
		parts := strings.SplitN(f, ":", 3)
		if len(parts) < 3 {
			return nil, fmt.Errorf("tag doesn't have at least 3 colons: %v", f)
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
