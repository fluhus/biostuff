// Handles SAM files.
package sam

import (
	"os"
	"bufio"
	"tools"
	"strings"
	"strconv"
)

// *** SAM LINE ****************************************************************

// Number of expected fields in a SAM line.
const samFields = 11

// Represents a single SAM line.
// Contains the mandatory fields according to the SAM format spec.
type SamLine struct {
	Qname string   // Query name
	Flag  int      // Bitwise flag (??)
	Rname string   // Reference sequence name
	Pos   int      // Mapping position (1-based)
	Mapq  int      // Mapping quality
	Cigar string   // CIGAR string (??)
	Rnext string   // ??
	Pnext string   // ??
	Tlen  int      // Observed template length (??)
	Seq   string   // Sequence
	Qual  string   // Phred qualities (ASCII)
}

// *** SAM READER **************************************************************

// Reads sam lines from a file.
type SamReader struct {
	rd *bufio.Reader
}

// Returns a new SAM reader (1MB buffer).
// If file cannot be opened, returns nil.
func NewSamReader(name string) *SamReader {
	// Open file
	file, err := os.Open(name)
	if err != nil {return nil}

	return &SamReader{bufio.NewReaderSize(file, tools.Mega)}
}

// Returns the next line from a SAM file.
// Returns nil if the line is corrupt (bad format), a read error occured or
// reached EOF.
func (srd *SamReader) NextLine() *SamLine {
	var line []byte

	for {
		// Try to read line
		line, err := srd.rd.ReadBytes('\n')
		line = []byte(strings.Trim(string(line), "\r\n"))

		// Check for error
		if err != nil {return nil}

		// Break only if non-comment (starting with '@')
		if line[0] != '@' {
			break
		}
	}

	// Split to fields
	fields := strings.Fields(string(line))
	if len(fields) < samFields {return nil}

	// Generate result & assign fields
	Atoi := strconv.Atoi

	result := &SamLine{}

	// BUG( ) TODO check int parsing errors and alert about them
	result.Qname   = fields[0]
	result.Flag, _ = Atoi(fields[1])
	result.Rname   = fields[2]
	result.Pos, _  = Atoi(fields[3])
	result.Mapq, _ = Atoi(fields[4])
	result.Cigar   = fields[5]
	result.Rnext   = fields[6]
	result.Pnext   = fields[7]
	result.Tlen, _ = Atoi(fields[8])
	result.Seq     = fields[9]
	result.Qual    = fields[10]

	return result
}






