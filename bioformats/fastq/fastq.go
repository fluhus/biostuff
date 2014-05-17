// Deals with Fastq reading and writing.
package fastq

import (
	"io"
	"os"
	"fmt"
	"bufio"
	"bytes"
	"errors"
)

// Enables assertions.
const assert = true;

// Reports if assertions are enabled.
func init() {
	if assert {
		fmt.Fprintln(os.Stderr, "*** package fastq: assertions enabled ***")
	}
}

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

// Returns a byte representation of the given score, for the Fastq format.
func phred(score int) byte {
	// Score <=> byte offset (see Fastq spec in Wikipedia)
	// BUG( ) Is this offset valid for all sequencing machines?
	const offset = 64

	return byte(score + offset)
}

// Returns a qual string for the given sequence.
func MakeQuals(sequence []byte) []byte {
	result := make([]byte, len(sequence))

	// Generate values
	for i := range result {
		result[i] = phred(60) // Arbitrary high value, to isolate mapping
		                      // from quality considerations
	}

	return result
}
// Bug( ) TODO make a more realistic quality generation algorithm.

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
	id = bytes.Trim(id, "\n\r")
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
	
	seq = bytes.Trim(seq, "\n\r")
	
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
	plus = bytes.Trim(plus, "\n\r")
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
	quals = bytes.Trim(quals, "\n\r")
	if len(quals) != len(seq) {
		return nil, errors.New("fastq read error: sequence and qualities have" +
				" different lengths")
		// BUG( ) TODO should I include more details in the error message?
	}
	
	// Finally done!
	return &Fastq{ id, seq, quals }, nil
}



