// Package fastq deals with Fastq reading and writing.
package fastq

import (
	"bufio"
	"fmt"
	"io"
)

// Fastq represents a single Fastq entry.
type Fastq struct {
	Name     []byte // Entry name (without the '@')
	Sequence []byte // Sequence as received
	Quals    []byte // Qualities as received
}

// BytesReader is anything that has the ReadBytes method.
type BytesReader interface {
	ReadBytes(byte) ([]byte, error)
}

// Next reads the next fastq entry from the reader.
// Returns a non-nil error if reading fails, or io.EOF if encountered end of
// file. When EOF is returned, no fastq is available. On error, the returned
// fastq will be nil.
func Next(reader BytesReader) (*Fastq, error) {
	// Read name.
	name, err := reader.ReadBytes('\n')
	if err != nil {
		// If end of file and no read, we're done.
		if err == io.EOF {
			if len(name) == 0 {
				return nil, err
			}
			return nil, io.ErrUnexpectedEOF
		}
		// Not end of file.
		return nil, fmt.Errorf("fastq read: %v", err)
	}

	// Handle name.
	name = trimNewLines(name)
	if len(name) == 0 || name[0] != '@' {
		return nil, fmt.Errorf("fastq read: expected '@' at beginning of"+
			" line: %q", string(name))
	}

	// Trim '@'
	name = name[1:]

	// Read sequence
	seq, err := reader.ReadBytes('\n')
	if err != nil {
		// If end of file, report unexpected
		if err == io.EOF {
			return nil, fmt.Errorf("fastq read: unexpected end of file")

		}
		// Not end of file
		return nil, fmt.Errorf("fastq read: %v", err)
	}

	seq = trimNewLines(seq)

	// Read plus
	plus, err := reader.ReadBytes('\n')
	if err != nil {
		// If end of file, report unexpected
		if err == io.EOF {
			return nil, io.ErrUnexpectedEOF

		}
		// Not end of file
		return nil, fmt.Errorf("fastq read: %v", err)
	}

	// Handle plus
	plus = trimNewLines(plus)
	if len(plus) == 0 || plus[0] != '+' {
		return nil, fmt.Errorf("fastq read: expected '+' at beginning of"+
			" line: %q", string(plus))
	}

	// Read qualities
	quals, err := reader.ReadBytes('\n')
	if err != nil {
		// If end of file, ignore and report on next read
		if err != io.EOF {
			return nil, fmt.Errorf("fastq read: %v", err)
		}
	}

	// Handle qualities
	quals = trimNewLines(quals)
	if len(quals) != len(seq) {
		return nil, fmt.Errorf("fastq read: sequence and qualities have" +
			" different lengths")
		// TODO(amit): should I include more details in the error message?
	}

	// Finally done!
	return &Fastq{name, seq, quals}, nil
}

// trimNewLines Trims new-line and carriage-return from both ends of the slice.
func trimNewLines(b []byte) []byte {
	start := 0
	end := len(b)
	for i, v := range b {

		if v == '\n' || v == '\r' {
			// If encountered new lines up to here.
			if i == start {
				start++
			}
		} else {
			end = i + 1 // end will point to the last non-new-line
		}

	}
	return b[start:end]
}

// ForEach reads all the sequences from the given fastq stream, until EOF.
// Calls f on each sequence. If f returns a non-nil error, iteration is stopped
// and the error is returned.
func ForEach(r io.Reader, f func(*Fastq) error) error {
	buf := bufio.NewReader(r)
	var fa *Fastq
	var err error

	for fa, err = Next(buf); err == nil; fa, err = Next(buf) {
		if err := f(fa); err != nil {
			return err
		}
	}
	if err != io.EOF {
		return err
	}

	return nil
}
