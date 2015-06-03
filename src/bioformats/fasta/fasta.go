// Deals with fasta parsing and representation.
package fasta

import (
	"fmt"
	"bufio"
	"io"
)

// Converts number to nucleotide.
var num2nuc = []byte{'A', 'C', 'G', 'T'}

// Converts nucleotide to number.
var nuc2num = map[byte]uint {'A':0, 'C':1, 'G':2, 'T':3, 'N':4, 'a':0, 'c':1, 'g':2, 't':3, 'n':4}


// *** FASTA ENTRY ************************************************************

// A single immutable fasta sequence, stored in 2-bit representation.
type FastaEntry struct {
	name     string            // sequence name (row that starts with '>')
	sequence []byte            // sequence in 2-bit format
	length   uint              // number of nucleotides
	isN      map[uint]struct{} // coordinates of 'N' nucleotides
}

// Returns an empty fasta entry.
func newFastaEntry() *FastaEntry {
	return &FastaEntry{"", nil, 0, make(map[uint]struct{})}
}

// Returns the number of nucleotides in this fasta entry.
func (f *FastaEntry) Length() int {
	return int(f.length)
}

// Returns the name of the fasta entry.
func (f *FastaEntry) Name() string {
	return f.name
}

// Returns the nucleotide at the given position.
func (f *FastaEntry) At(position int) byte {
	uposition := uint(position)

	// Check if N.
	if _,n := f.isN[uposition]; n {
		return 'N'
	}

	// Extract nucleotide.
	num := (f.sequence[uposition / 4] >> (uposition % 4 * 2) & 3)
	
	return num2nuc[num]
}

// Appends a nucleotide to the fasta entry.
func (f *FastaEntry) append(nuc byte) error {
	num, ok := nuc2num[nuc]

	// If unknown nucleotide.
	if !ok {
		return fmt.Errorf("Bad nucleotide: " + string([]byte{nuc}))
	}

	// If 'N'.
	if num == 4 {
		num = 0
		f.isN[f.length] = struct{}{}
	}

	// Append an extra byte.
	if f.length % 4 == 0 {
		f.sequence = append(f.sequence, 0)
	}

	// Set bits.
	f.sequence[f.length / 4] |= byte( num << (f.length % 4 * 2) )

	f.length++

	return nil
}

// Extracts a subsequence from the fasta.
func (f *FastaEntry) Subsequence(start, length int) []byte {
	// Check input.
	if length < 0 {
		panic(fmt.Sprint("Bad subsequence length: %d", length))
	}
	if start < 0 {
		panic(fmt.Sprint("Bad subsequence start: %d", start))
	}
	if start + length > f.Length() {
		panic(fmt.Sprint("Subsequence position exceeds sequence length: " +
				"start %d, length %d.", start, length))
	}

	// Generate result.
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = f.At(start + i)
	}

	return result
}

// String representation of an entry. Format: name[length]
func (f *FastaEntry) String() string {
	return fmt.Sprintf("%s[%d]", f.name, f.Length())
}

// Reads a single fasta entry from a stream. Returns EOF only if nothing was
// read.
func ReadFastaEntry(r *bufio.Reader) (*FastaEntry, error) {
	// States of the reader.
	const (
		stateStart = iota  // beginning of input
		stateName          // reading name
		stateNewLine       // beginning of new line
		stateSequence      // reading sequence
	)
	
	// Result entry.
	var name []byte
	result := newFastaEntry()
	
	// Start reading.
	state := stateStart
	var b byte
	var err error
	readAnything := false
	
	loop: for b, err = r.ReadByte(); err == nil; b, err = r.ReadByte() {
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
					if err != nil { break loop }
				}
			}
			
		case stateSequence:
			if b == '\n' || b == '\r' {
				state = stateNewLine
			} else {
				err = result.append(b)
				if err != nil { break loop }
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
				if err != nil { break loop }
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
	
	return result, nil
}

// Reads all fasta entries from the given stream, until EOF. Stream will be
// buffered inside the function.
func ReadFasta(r io.Reader) ([]*FastaEntry, error) {
	buf := bufio.NewReader(r)
	
	var result []*FastaEntry
	var fa *FastaEntry
	var err error

	for fa, err = ReadFastaEntry(buf); err == nil; fa, err =
			ReadFastaEntry(buf) {
		result = append(result, fa)
	}

	if err != io.EOF {
		return nil, err
	}

	return result, nil
}
