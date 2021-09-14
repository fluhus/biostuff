// Package fastq parses fastq files.
//
// This package uses the format described in:
// https://en.wikipedia.org/wiki/FASTQ_format
//
// This package does not validate sequence and quality characters.
package fastq

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// Fastq is a single sequence in a fastq file.
type Fastq struct {
	Name     []byte // Entry name (without the '@')
	Sequence []byte // Sequence as received
	Quals    []byte // Qualities as received
}

// Text returns the textual representation of f in fastq format.
// Includes a trailing new line.
func (f *Fastq) Text() []byte {
	n := 6 + len(f.Name) + len(f.Sequence) + len(f.Quals)
	buf := bytes.NewBuffer(make([]byte, 0, n))
	buf.WriteByte('@')
	buf.Write(f.Name)
	buf.WriteByte('\n')
	buf.Write(f.Sequence)
	buf.WriteString("\n+\n")
	buf.Write(f.Quals)
	buf.WriteByte('\n')
	if buf.Len() != n {
		panic(fmt.Sprintf("bad length: %v expected %v", buf.Len(), n))
	}
	return buf.Bytes()
}

// A Reader reads sequences from a fastq stream.
type Reader struct {
	s *bufio.Scanner
}

// NewReader returns a new fastq reader that reads from r.
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
