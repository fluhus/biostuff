// Deals with fasta parsing and representation.
package fasta

import (
	"bufio"
	"fmt"
	"io"
	"runtime/debug"
	"sort"
)

// Converts number to nucleotide.
var num2nuc = []byte{'A', 'C', 'G', 'T'}

// *** FASTA ENTRY ************************************************************

// A single immutable fasta sequence, stored in 2-bit representation.
type Entry struct {
	name     string // sequence name (row that starts with '>')
	sequence []byte // sequence in 2-bit format
	length   uint   // number of nucleotides
	nStarts  []uint // coordinates of starts of 'N' chunks
	nEnds    []uint // coordinates of ends of 'N' chunks (exclusive)
}

// Returns an empty fasta entry.
func newEntry() *Entry {
	return &Entry{"", nil, 0, nil, nil}
}

// Returns the number of nucleotides in this fasta entry.
func (f *Entry) Length() int {
	return int(f.length)
}

// Returns the name of the fasta entry.
func (f *Entry) Name() string {
	return f.name
}

// Returns the nucleotide at the given position.
func (f *Entry) At(position int) byte {
	uposition := uint(position)

	// Check if N.
	if f.isN(uposition) {
		return 'N'
	}

	// Extract nucleotide.
	num := (f.sequence[uposition/4] >> (uposition % 4 * 2) & 3)

	return num2nuc[num]
}

// Checks whether the given pos holds an 'N'.
func (f *Entry) isN(pos uint) bool {
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

// Appends a nucleotide to the fasta entry.
func (f *Entry) append(nuc byte) error {
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

// Extracts a subsequence from the fasta.
func (f *Entry) Subsequence(start, length int) []byte {
	// Check input.
	if length < 0 {
		panic(fmt.Sprintf("Bad subsequence length: %d", length))
	}
	if start < 0 {
		panic(fmt.Sprintf("Bad subsequence start: %d", start))
	}
	if start+length > f.Length() {
		panic(fmt.Sprint("Subsequence position exceeds sequence length: "+
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
func (f *Entry) String() string {
	return fmt.Sprintf("%s[%d]", f.name, f.Length())
}

// Reads a single fasta entry from a stream. Returns EOF only if nothing was
// read.
func ReadEntry(r *bufio.Reader) (*Entry, error) {
	// States of the reader.
	const (
		stateStart    = iota // beginning of input
		stateName            // reading name
		stateNewLine         // beginning of new line
		stateSequence        // reading sequence
	)

	// Result entry.
	var name []byte
	result := newEntry()

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

// Reads all fasta entries from the given stream, until EOF. Stream will be
// buffered inside the function.
func ReadFasta(r io.Reader) ([]*Entry, error) {
	buf := bufio.NewReader(r)

	var result []*Entry
	var fa *Entry
	var err error

	for fa, err = ReadEntry(buf); err == nil; fa, err =
		ReadEntry(buf) {
		result = append(result, fa)

		// Release unused memory, so that the program doesn't
		// consume twice the memory it really needs.
		debug.FreeOSMemory()
	}

	if err != io.EOF {
		return nil, err
	}

	return result, nil
}

// ***** SERIALIZABLE ENTRY ****************************************************

// A fasta entry with exported fields for serialization.
type SerializableEntry struct {
	Name     string
	Sequence []byte
	Length   uint
	NStarts  []uint
	NEnds    []uint
}

// Converts a fasta entry to a serializable one.
func ToSerializable(f *Entry) *SerializableEntry {
	return &SerializableEntry{f.name, f.sequence, f.length, f.nStarts,
		f.nEnds}
}

// Converts a serializable fasta entry to a regular one.
func FromSerializable(f *SerializableEntry) *Entry {
	return &Entry{f.Name, f.Sequence, f.Length, f.NStarts,
		f.NEnds}
}
