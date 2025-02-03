// Package fasta decodes and encodes fasta files.
//
// This package uses the format described in:
// https://en.wikipedia.org/wiki/FASTA_format
//
// This package does not validate sequence characters.
package fasta

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"iter"

	"github.com/fluhus/gostuff/aio"
)

const (
	textLineLen = 80
)

// Fasta is a single sequence in a fasta file.
type Fasta struct {
	Name     []byte // Sequence name (without the '>')
	Sequence []byte // Sequence
}

// MarshalText returns the textual representation of f in fasta format. Includes a
// trailing new line. Always includes a name line, even for empty names. Sequence
// gets broken down into lines of length 80.
func (f *Fasta) MarshalText() ([]byte, error) {
	n := 2 + len(f.Name) + len(f.Sequence) +
		(len(f.Sequence)+textLineLen-1)/textLineLen
	buf := bytes.NewBuffer(make([]byte, 0, n))
	f.Write(buf)
	if buf.Len() != n {
		panic(fmt.Sprintf("bad len: %v want %v", buf.Len(), n))
	}
	return buf.Bytes(), nil
}

// Write writes this entry in textual Fasta format to the given writer.
// Includes a trailing new line.
func (f *Fasta) Write(w io.Writer) error {
	if _, err := fmt.Fprintf(w, ">%s\n", f.Name); err != nil {
		return err
	}
	for i := 0; i < len(f.Sequence); i += textLineLen {
		to := min(i+textLineLen, len(f.Sequence))
		if _, err := fmt.Fprintf(w, "%s\n", f.Sequence[i:to]); err != nil {
			return err
		}
	}
	return nil
}

// Reader returns an iterator over fasta entries in a reader.
func Reader(r io.Reader) iter.Seq2[*Fasta, error] {
	return func(yield func(*Fasta, error) bool) {
		br := bufio.NewReader(r)
		buf := make([]byte, 4096)

		// States of the reader.
		const (
			stateStart    = iota // Beginning of input
			stateNewLine         // Beginning of new line
			stateName            // Middle of name
			stateSequence        // Middle sequence
		)
		state := stateStart
		result := &Fasta{}

		for {
			n, err := br.Read(buf)
			if err != nil && err != io.EOF {
				yield(nil, err)
				return
			}
			if n == 0 {
				if len(result.Sequence) > 0 || len(result.Name) > 0 {
					yield(result, nil)
				}
				return
			}

			for _, b := range buf[:n] {
				switch state {
				case stateStart:
					// '>' marks the name of the sequence.
					if b == '>' {
						state = stateName
					} else {
						// If no '>' then only sequence without name.
						state = stateSequence
						if b == '\n' || b == '\r' {
							state = stateNewLine
						} else {
							result.Sequence = append(result.Sequence, b)
						}
					}

				case stateSequence:
					if b == '\n' || b == '\r' {
						state = stateNewLine
					} else {
						result.Sequence = append(result.Sequence, b)
					}

				case stateName:
					if b == '\n' || b == '\r' {
						state = stateNewLine
					} else {
						result.Name = append(result.Name, b)
					}

				case stateNewLine:
					if b == '\n' || b == '\r' {
						// Nothing. Move on to the next line.
					} else if b == '>' {
						// New sequence => yield current.
						state = stateName
						if !yield(result, nil) {
							return
						}
						result = &Fasta{}
					} else {
						// Just more sequence.
						state = stateSequence
						result.Sequence = append(result.Sequence, b)
					}
				}
			}
		}
	}
}

// File returns an iterator over fasta entries in a file.
func File(file string) iter.Seq2[*Fasta, error] {
	return func(yield func(*Fasta, error) bool) {
		f, err := aio.Open(file)
		if err != nil {
			yield(nil, err)
			return
		}
		defer f.Close()
		for fa, err := range Reader(f) {
			if !yield(fa, err) {
				break
			}
		}
	}
}
