// Package fastq deals with Fastq reading and writing.
package fastq

// TODO(amit): Add writing.

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// Fastq represents a single Fastq entry.
type Fastq struct {
	Name     []byte // Entry name (without the '@')
	Sequence []byte // Sequence as received
	Quals    []byte // Qualities as received
}

// A Reader reads text from an input and returns Fastq objects.
type Reader struct {
	s *bufio.Scanner
}

// NewReader returns a new Fastq reader.
func NewReader(r io.Reader) *Reader {
	return &Reader{s: bufio.NewScanner(r)}
}

// Next reads the next fastq entry from the reader.
// Returns a non-nil error if reading fails, or io.EOF if encountered end of
// file. When EOF is returned, no fastq is available.
func (r *Reader) Next() (*Fastq, error) {
	// Read name.
	if !r.s.Scan() {
		if r.s.Err() == nil {
			return nil, io.EOF
		}
		return nil, fmt.Errorf("fastq read: %v", r.s.Err())
	}
	name := copyBytes(r.s.Bytes())

	// Handle name.
	if len(name) == 0 || name[0] != '@' {
		return nil, fmt.Errorf("fastq read: expected '@' at beginning of"+
			" line: %q", string(name))
	}
	name = name[1:]

	// Read sequence
	if !r.s.Scan() {
		if r.s.Err() == nil {
			return nil, io.ErrUnexpectedEOF
		}
		return nil, fmt.Errorf("fastq read: %v", r.s.Err())
	}
	seq := copyBytes(r.s.Bytes())

	// Read plus
	if !r.s.Scan() {
		if r.s.Err() == nil {
			return nil, io.ErrUnexpectedEOF
		}
		return nil, fmt.Errorf("fastq read: %v", r.s.Err())
	}
	plus := copyBytes(r.s.Bytes())
	if !bytes.Equal(plus, []byte("+")) {
		return nil, fmt.Errorf("fastq read: expected '+' at beginning of"+
			" line: %q", string(plus))
	}

	// Read qualities
	if !r.s.Scan() {
		if r.s.Err() == nil {
			return nil, io.ErrUnexpectedEOF
		}
		return nil, fmt.Errorf("fastq read: %v", r.s.Err())
	}
	quals := copyBytes(r.s.Bytes())
	if len(quals) != len(seq) {
		return nil, fmt.Errorf("fastq read: sequence and qualities have"+
			" different lengths: %v and %v", len(seq), len(quals))
	}

	return &Fastq{name, seq, quals}, nil
}

// Copies the given bytes to a newly allocated slice.
func copyBytes(src []byte) []byte {
	b := make([]byte, len(src))
	copy(b, src)
	return b
}
