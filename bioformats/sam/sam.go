// Handles SAM files.
package sam

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// *** SAM LINE ****************************************************************

// Number of expected fields in a SAM line.
const samFields = 11

// Represents a single SAM line.
// Contains the mandatory fields according to the SAM format spec.
type Sam struct {
	Qname string // Query name
	Flag  int    // Bitwise flag (??)
	Rname string // Reference sequence name
	Pos   int    // Mapping position (1-based)
	Mapq  int    // Mapping quality
	Cigar string // CIGAR string (??)
	Rnext string // ??
	Pnext string // ??
	Tlen  int    // Observed template length (??)
	Seq   string // Sequence
	Qual  string // Phred qualities (ASCII)
}

// A string ready to be printed as a line in a sam file (no new line
// at the end).
func (s *Sam) String() string {
	return fmt.Sprintf("%s %d %s %d %d %s %s %s %d %s %s",
		s.Qname, s.Flag, s.Rname, s.Pos, s.Mapq, s.Cigar,
		s.Rnext, s.Pnext, s.Tlen, s.Seq, s.Qual)
}

// *** SAM READER **************************************************************

// Returns the next line from a SAM file.
// Returns the error that the read action had returned.
func ReadNext(in *bufio.Reader) (*Sam, error) {
	var line []byte
	var err error

	for {
		// Try to read line
		line, err = in.ReadBytes('\n')
		line = []byte(strings.Trim(string(line), "\r\n"))

		// Check for error (ignore if EOF and a line was read)
		if err != nil &&
			!(err == io.EOF && len(line) > 0) {
			return nil, err
		}

		// Break only if non-comment (starting with '@')
		if line[0] != '@' {
			break
		}
	}

	// Split to fields
	fields := strings.Fields(string(line))
	if len(fields) < samFields {
		return nil, errors.New(fmt.Sprintf("bad SAM line, only"+
			" %d fields (out of required %d)", len(fields), samFields))
	}

	// Generate result & assign fields
	atoi := strconv.Atoi

	result := &Sam{}

	// BUG( ) TODO check int parsing errors and alert about them

	result.Qname = fields[0]
	result.Flag, _ = atoi(fields[1])
	result.Rname = fields[2]
	result.Pos, _ = atoi(fields[3])
	result.Mapq, _ = atoi(fields[4])
	result.Cigar = fields[5]
	result.Rnext = fields[6]
	result.Pnext = fields[7]
	result.Tlen, _ = atoi(fields[8])
	result.Seq = fields[9]
	result.Qual = fields[10]

	return result, nil
}
