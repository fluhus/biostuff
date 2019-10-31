// Package fastq deals with Fastq reading and writing.
package fastq

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

// Represents a single Fastq entry.
type Fastq struct {
	Name     []byte
	Sequence []byte
	Quals    []byte
}

// Returns a formatted representation of the entry, ready to be printed
// (no new line at the end).
func (f *Fastq) String() string {
	// BUG(amit): Use MarshalText for encoding.
	return fmt.Sprintf("@%s\n%s\n+\n%s", f.Name, f.Sequence, f.Quals)
}

// Read reads the next fastq entry from the reader.
// Returns a non-nil error if reading fails, or io.EOF if encountered end of
// file. When EOF is returned, no fastq is available. On error, the returned
// fastq will be nil.
func Read(reader *bufio.Reader) (*Fastq, error) {
	// Read name.
	name, err := reader.ReadBytes('\n')
	if err != nil {
		// If end of file and no read, we're done.
		if err == io.EOF {
			if len(name) == 0 {
				return nil, err
			} else {
				return nil, errors.New("fastq read: unexpected end of file")
			}

			// Not end of file.
		} else {
			return nil, errors.New("fastq read: " + err.Error())
		}
	}

	// Handle name.
	name = trimNewLines(name)
	if len(name) == 0 || name[0] != '@' {
		return nil, errors.New("fastq read: expected '@' at beginning of" +
			" line: \"" + string(name) + "\"")
	}

	// Trim '@'
	name = name[1:]

	// Read sequence
	seq, err := reader.ReadBytes('\n')
	if err != nil {
		// If end of file, report unexpected
		if err == io.EOF {
			return nil, errors.New("fastq read: unexpected end of file")

			// Not end of file
		} else {
			return nil, errors.New("fastq read: " + err.Error())
		}
	}

	seq = trimNewLines(seq)

	// Read plus
	plus, err := reader.ReadBytes('\n')
	if err != nil {
		// If end of file, report unexpected
		if err == io.EOF {
			return nil, errors.New("fastq read: unexpected end of file")

			// Not end of file
		} else {
			return nil, errors.New("fastq read: " + err.Error())
		}
	}

	// Handle plus
	plus = trimNewLines(plus)
	if len(plus) == 0 || plus[0] != '+' {
		return nil, errors.New("fastq read: expected '+' at beginning of" +
			" line: \"" + string(plus) + "\"")
	}

	// Read qualities
	quals, err := reader.ReadBytes('\n')
	if err != nil {
		// If end of file, ignore and report on next read
		if err != io.EOF {
			return nil, errors.New("fastq read: " + err.Error())
		}
	}

	// Handle qualities
	quals = trimNewLines(quals)
	if len(quals) != len(seq) {
		return nil, errors.New("fastq read: sequence and qualities have" +
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
