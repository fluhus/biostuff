// Package sam handles SAM files.
package sam

import (
	"bufio"
	"fmt"
	"io"

	"github.com/fluhus/gostuff/csvdec"
)

// TODO(amit): Add marshaling.

// A raw structure for the initial parsing using csvdec.
type samRaw struct {
	Qname          string
	Flag           int
	Rname          string
	Pos            int
	Mapq           int
	Cigar          string
	Rnext          string
	Pnext          int
	Tlen           int
	Seq            string
	Qual           string
	OptionalFields []string
}

// SAM represents a single SAM line.
type SAM struct {
	Qname          string                 // Query name
	Flag           int                    // Bitwise flag
	Rname          string                 // Reference sequence name
	Pos            int                    // Mapping position (1-based)
	Mapq           int                    // Mapping quality
	Cigar          string                 // CIGAR string
	Rnext          string                 // Ref. name of the mate/next read
	Pnext          int                    // Position of the mate/next read
	Tlen           int                    // Observed template length
	Seq            string                 // Sequence
	Qual           string                 // Phred qualities (ASCII)
	OptionalFields map[string]interface{} // Typed optional fields.
}

// Converts a raw SAM struct to an exported SAM struct.
func fromRaw(raw *samRaw) *SAM {
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
	result.OptionalFields = map[string]interface{}{}
	// TODO(amit): Complete parsing of optional fields.
	return result
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
	return fromRaw(raw), nil
}
