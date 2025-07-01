// Package fastq decodes and encodes fastq files.
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

	"github.com/fluhus/gostuff/snm"
)

// Fastq is a single sequence in a fastq file.
type Fastq struct {
	Name     []byte // Entry name (without the '@').
	Sequence []byte // Sequence.
	Quals    []byte // Qualities as ASCII characters.
}

// MarshalText returns the textual representation of f in fastq format.
// Includes a trailing new line.
func (f *Fastq) MarshalText() ([]byte, error) {
	n := 6 + len(f.Name) + len(f.Sequence) + len(f.Quals)
	buf := bytes.NewBuffer(make([]byte, 0, n))
	f.Write(buf)
	if buf.Len() != n {
		panic(fmt.Sprintf("bad length: %v expected %v", buf.Len(), n))
	}
	return buf.Bytes(), nil
}

// Write writes this entry in textual Fastq format to the given writer.
// Includes a trailing new line.
func (f *Fastq) Write(w io.Writer) error {
	_, err := fmt.Fprintf(w, "@%s\n%s\n+\n%s\n", f.Name, f.Sequence, f.Quals)
	return err
}

// A reader reads sequences from a fastq stream.
type reader struct {
	s *bufio.Scanner
}

// Returns a new fastq reader that reads from r.
func newReader(r io.Reader) *reader {
	return &reader{s: bufio.NewScanner(r)}
}

// Reads the next fastq entry from the reader.
// Returns a non-nil error if reading fails, or io.EOF if encountered end of
// file. When EOF is returned, no fastq is available.
func (r *reader) read() (*Fastq, error) {
	// Read name.
	if !r.s.Scan() {
		if r.s.Err() == nil {
			return nil, io.EOF
		}
		return nil, fmt.Errorf("fastq read: %v", r.s.Err())
	}
	name := snm.TightClone(r.s.Bytes())

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
	seq := snm.TightClone(r.s.Bytes())

	// Read plus
	if !r.s.Scan() {
		if r.s.Err() == nil {
			return nil, io.ErrUnexpectedEOF
		}
		return nil, fmt.Errorf("fastq read: %v", r.s.Err())
	}
	plus := r.s.Bytes()
	if !bytes.HasPrefix(plus, []byte("+")) {
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
	quals := snm.TightClone(r.s.Bytes())
	if len(quals) != len(seq) {
		return nil, fmt.Errorf("fastq read: sequence and qualities have"+
			" different lengths: %v and %v", len(seq), len(quals))
	}

	return &Fastq{name, seq, quals}, nil
}
