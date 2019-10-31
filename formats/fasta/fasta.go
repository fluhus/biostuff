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
	"fmt"
	"io"
	"sort"
)

// BUG(amit): Add an option to create a sequence.

// Converts number to nucleotide.
var num2nuc = []byte{'A', 'C', 'G', 'T'}

// Fasta is a single immutable fasta sequence.
type Fasta struct {
	name     string // sequence name (row that starts with '>')
	sequence []byte // sequence in 2-bit format
	length   uint   // number of nucleotides
	nStarts  []uint // coordinates of starts of 'N' chunks
	nEnds    []uint // coordinates of ends of 'N' chunks (exclusive)
}

// newFasta returns an empty fasta sequence.
func newFasta() *Fasta {
	return &Fasta{"", nil, 0, nil, nil}
}

// Len returns the number of nucleotides in this sequence.
func (f *Fasta) Len() int {
	return int(f.length)
}

// Name returns the name of the fasta, without the '>' prefix.
func (f *Fasta) Name() string {
	return f.name
}

// At returns the nucleotide at the given position.
func (f *Fasta) At(position int) byte {
	uposition := uint(position)

	// Check if N.
	if f.isN(uposition) {
		return 'N'
	}

	// Extract nucleotide.
	num := (f.sequence[uposition/4] >> (uposition % 4 * 2) & 3)

	return num2nuc[num]
}

// isN checks whether the given position holds an 'N'.
func (f *Fasta) isN(pos uint) bool {
	if len(f.nStarts) == 0 {
		return false
	}

	i := sort.Search(len(f.nStarts), func(j int) bool {
		return f.nStarts[j] > pos
	}) - 1

	if i == -1 {
		i = 0
	}
	return pos >= f.nStarts[i] && pos < f.nEnds[i]
}

// appends adds a nucleotide to the fasta sequence.
func (f *Fasta) append(nuc byte) error {
	var num uint
	switch nuc {
	case 'a', 'A':
		num = 0
	case 'c', 'C':
		num = 1
	case 'g', 'G':
		num = 2
	case 't', 'T':
		num = 3
	case 'n', 'N':
		num = 4
	default:
		num = 5
	}

	// If unknown nucleotide.
	if num == 5 {
		return fmt.Errorf("Bad nucleotide: " + string([]byte{nuc}))
	}

	// If 'N'.
	if num == 4 {
		num = 0

		// Start a new chunk?
		if len(f.nEnds) == 0 || f.nEnds[len(f.nEnds)-1] < f.length {
			// Yes.
			f.nStarts = append(f.nStarts, f.length)
			f.nEnds = append(f.nEnds, f.length+1)
		} else {
			// No.
			f.nEnds[len(f.nEnds)-1] = f.length + 1
		}
	}

	// Append an extra byte.
	if f.length%4 == 0 {
		f.sequence = append(f.sequence, 0)
	}

	// Set bits.
	f.sequence[f.length/4] |= byte(num << (f.length % 4 * 2))

	f.length++

	return nil
}

// Subsequence extracts a subsequence from the fasta.
func (f *Fasta) Subsequence(start, length int) []byte {
	if length < 0 {
		panic(fmt.Sprintf("Bad subsequence length: %d", length))
	}
	if start < 0 {
		panic(fmt.Sprintf("Bad subsequence start: %d", start))
	}
	if start+length > f.Len() {
		panic(fmt.Sprint("Subsequence position exceeds sequence length: "+
			"start %d, length %d.", start, length))
	}
	// BUG(amit): Improve performance of Subsequence.

	// Generate result.
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = f.At(start + i)
	}

	return result
}

// Sequence returns the entire sequence.
func (f *Fasta) Sequence() []byte {
	return f.Subsequence(0, f.Len())
}

// String returns a string representation of a fasta, for debugging.
// Format: name[length]
func (f *Fasta) String() string {
	return fmt.Sprintf("%s[%d]", f.name, f.Len())
}

// Read reads a single fasta sequence from a stream. Returns EOF only if
// nothing was read.
func Read(r io.ByteScanner) (*Fasta, error) {
	// States of the reader.
	const (
		stateStart    = iota // beginning of input
		stateName            // reading name
		stateNewLine         // beginning of new line
		stateSequence        // reading sequence
	)

	var name []byte
	result := newFasta()

	// Start reading.
	state := stateStart
	var b byte
	var err error
	readAnything := false

loop:
	for b, err = r.ReadByte(); err == nil; b, err = r.ReadByte() {
		readAnything = true
		switch state {
		case stateStart:
			// '>' marks the name of the sequence.
			if b == '>' {
				state = stateName

				// If no '>' then only sequence without name.
			} else {
				state = stateSequence
				if b == '\n' || b == '\r' {
					state = stateNewLine
				} else {
					err = result.append(b)
					if err != nil {
						break loop
					}
				}
			}

		case stateSequence:
			if b == '\n' || b == '\r' {
				state = stateNewLine
			} else {
				err = result.append(b)
				if err != nil {
					break loop
				}
			}

		case stateName:
			if b == '\n' || b == '\r' {
				state = stateNewLine
			} else {
				name = append(name, b)
			}

		case stateNewLine:
			if b == '\n' || b == '\r' {
				// Nothing. Move on to the next line.
			} else if b == '>' {
				// New sequence => done reading.
				r.UnreadByte()
				break loop
			} else {
				// Just more sequence.
				state = stateSequence
				err = result.append(b)
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

	result.name = string(name)

	// Reallocate sequence to take less memory.
	newSequence := make([]byte, len(result.sequence))
	copy(newSequence, result.sequence)
	result.sequence = newSequence

	newStarts := make([]uint, len(result.nStarts))
	copy(newStarts, result.nStarts)
	result.nStarts = newStarts

	newEnds := make([]uint, len(result.nEnds))
	copy(newEnds, result.nEnds)
	result.nEnds = newEnds

	return result, nil
}

// ReadAll reads all fasta entries from the given stream, until EOF. Stream
// will be buffered inside the function.
func ReadAll(r io.Reader) ([]*Fasta, error) {
	buf := bufio.NewReader(r)

	var result []*Fasta
	var fa *Fasta
	var err error

	for fa, err = Read(buf); err == nil; fa, err =
		Read(buf) {
		result = append(result, fa)
	}

	if err != io.EOF {
		return nil, err
	}

	return result, nil
}
