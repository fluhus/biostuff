// Package fasta deals with fasta parsing and representation.
//
// Input Format
//
// The package uses the format described in:
// https://en.wikipedia.org/wiki/FASTA_format
//
// A valid fasta can have plain bases:
//
//  AAAAAATTTTTTCCCCCCGGGGGG
//
// Or have names separating the sequences, starting with '>':
//
//  >sequence1
//  AAAAAATTTTTTCCCCCCGGGGGG
//  >sequence2
//  AAAAAATTTTTTCCCCCCGGGGGG
//
// Output Format
//
// The package is case insensitive. 'A' and 'a' are equivalent. The output
// of all functions that return bases is in upper case.
//
// Memory Footprint
//
// Fasta sequences are represented in a 2-bit format. The size of a sequence
// with n bases should be n/4 bytes, plus 2 integers for each sequence of N's.
package fasta

import (
	"bufio"
	"io"
)

// Fasta is a single sequence in a fasta file.
type Fasta struct {
	Name     []byte // Sequence name (without the '>')
	Sequence []byte // Sequence
}

// A Reader reads sequences from a fasta stream.
type Reader struct {
	r *bufio.Reader
}

// NewReader returns a new reader for the given input.
func NewReader(r io.Reader) *Reader {
	return &Reader{bufio.NewReader(r)}
}

// Next reads a single fasta sequence from a stream. Returns EOF only if
// nothing was read.
func (r *Reader) Next() (*Fasta, error) {
	// States of the reader.
	const (
		stateStart    = iota // Beginning of input
		stateNewLine         // Beginning of new line
		stateName            // Middle of name
		stateSequence        // Middle sequence
	)

	// Start reading.
	result := &Fasta{}
	state := stateStart
	var b byte
	var err error
	readAnything := false

loop:
	for b, err = r.r.ReadByte(); err == nil; b, err = r.r.ReadByte() {
		readAnything = true
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
				// New sequence => done reading.
				r.r.UnreadByte()
				break loop
			} else {
				// Just more sequence.
				state = stateSequence
				result.Sequence = append(result.Sequence, b)
				if err != nil {
					break loop
				}
			}
		}
	}

	// Return EOF only if encountered before reading anything.
	if !readAnything {
		return nil, err
	}
	// EOF will be returned on the next call to read.
	if err != nil && err != io.EOF {
		return nil, err
	}

	return result, nil
}
