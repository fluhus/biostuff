// Deals with Fastq reading and writing.
package fastq

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"

	"github.com/fluhus/biostuff/seqtools"
)

// Represents a single Fastq entry.
type Fastq struct {
	Id       []byte
	Sequence []byte
	Quals    []byte
}

// Returns a formatted representation of the entry, ready to be printed
// (no new line at the end).
func (f *Fastq) String() string {
	return fmt.Sprintf("@%s\n%s\n+\n%s", f.Id, f.Sequence, f.Quals)
}

// Reads the next fastq entry from the reader.
// Returns a non-nil error if reading fails, or io.EOF if encountered end of
// file. When EOF is returned, no fastq is available. On error, the returned
// fastq will be nil.
func ReadNext(reader *bufio.Reader) (*Fastq, error) {
	// Read ID
	id, err := reader.ReadBytes('\n')
	if err != nil {
		// If end of file and no read, we're done
		if err == io.EOF {
			if len(id) == 0 {
				return nil, err
			} else {
				return nil, errors.New("fastq read error: unexpected end of file")
			}

			// Not end of file, bummer
		} else {
			return nil, errors.New("fastq read error: " + err.Error())
		}
	}

	// Handle ID
	id = trimNewLines(id)
	if len(id) == 0 || id[0] != '@' {
		return nil, errors.New("fastq read error: expected '@' at beginning of" +
			" line: \"" + string(id) + "\"")
	}

	// Trim '@'
	id = id[1:]

	// Read sequence
	seq, err := reader.ReadBytes('\n')
	if err != nil {
		// If end of file, report unexpected
		if err == io.EOF {
			return nil, errors.New("fastq read error: unexpected end of file")

			// Not end of file
		} else {
			return nil, errors.New("fastq read error: " + err.Error())
		}
	}

	seq = trimNewLines(seq)

	// Read plus
	plus, err := reader.ReadBytes('\n')
	if err != nil {
		// If end of file, report unexpected
		if err == io.EOF {
			return nil, errors.New("fastq read error: unexpected end of file")

			// Not end of file
		} else {
			return nil, errors.New("fastq read error: " + err.Error())
		}
	}

	// Handle plus
	plus = trimNewLines(plus)
	if len(plus) == 0 || plus[0] != '+' {
		return nil, errors.New("fastq read error: expected '+' at beginning of" +
			" line: \"" + string(plus) + "\"")
	}

	// Read qualities
	quals, err := reader.ReadBytes('\n')
	if err != nil {
		// If end of file, ignore and report on next read
		if err != io.EOF {
			return nil, errors.New("fastq read error: " + err.Error())
		}
	}

	// Handle qualities
	quals = trimNewLines(quals)
	if len(quals) != len(seq) {
		return nil, errors.New("fastq read error: sequence and qualities have" +
			" different lengths")
		// BUG( ) TODO should I include more details in the error message?
	}

	// Finally done!
	return &Fastq{id, seq, quals}, nil
}

// Used for different phred offsets.
type PhredOffset byte

const (
	Illumina18 PhredOffset = 33
)

// Applies single nucleotide errors according to the quality sequence.
// 'offset' is a positive value of the phred's offset.
//
// Modifies the sequence and cannot be undone!
func (f *Fastq) ApplyQuals(offset PhredOffset) {
	// Check offset
	if offset < 0 {
		panic(fmt.Sprint("bad offset:", offset))
	}

	// Check quality length
	if len(f.Quals) != len(f.Sequence) {
		panic(fmt.Sprintf("inconsistent sequence and quals lengths: %d, %d",
			len(f.Sequence), len(f.Quals)))
	}

	// Go over qualities
	for i := range f.Quals {
		// Extract real quality
		phred := f.Quals[i] - byte(offset)
		qual := math.Pow(10.0, float64(phred)/-10.0)

		// Mutate randomly
		if rand.Float64() < qual {
			originalChar := f.Sequence[i]
			for f.Sequence[i] == originalChar {
				f.Sequence[i] = seqtools.RandNuc()
			}
		}
	}
}

// Creates a FastQ quality sequence for
// the given nucleotide sequence.
func MakeQuals(sequence []byte) []byte {
	// BUG( ) I should replace this mock with a real quality generator.
	result := make([]byte, len(sequence))
	for i := range result {
		result[i] = 'I'
	}
	return result
}

// Trims new-line and carriage-return from both ends of the slice.
func trimNewLines(b []byte) []byte {
	start := 0
	end := len(b)
	for i, v := range b {

		if v == '\n' || v == '\r' {
			// If encountered new lines up to here
			if i == start {
				start++
			}
		} else {
			end = i + 1 // end will point to the last non-new-line
		}

	}

	return b[start:end]
}
