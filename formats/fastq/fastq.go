// Package fastq deals with Fastq reading and writing.
package fastq

import (
	"fmt"
	"io"
)

// Fastq represents a single Fastq entry.
type Fastq struct {
	Name     []byte
	Sequence []byte
	Quals    []byte
}

// BytesReader is anything that has the ReadBytes method.
type BytesReader interface {
	ReadBytes(byte) ([]byte, error)
}

// Read reads the next fastq entry from the reader.
// Returns a non-nil error if reading fails, or io.EOF if encountered end of
// file. When EOF is returned, no fastq is available. On error, the returned
// fastq will be nil.
func Read(reader BytesReader) (*Fastq, error) {
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
