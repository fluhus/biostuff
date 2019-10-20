// Handles bed file representation and parsing.
package bed

import (
	"fmt"
	"strconv"
	"strings"
)

// A simple genomic region notation.
type Bed struct {
	Chr   string
	Start int
	End   int
}

// Parses a single bed line. Keeps chromosome, start and end in the bed object.
// All other fields are returned in a string array. Returns a non-nil error if
// couldn't parse.
func Parse(s string) (*Bed, []string, error) {
	// Split
	fields := strings.Split(s, "\t")
	if len(fields) < 3 {
		return nil, nil, fmt.Errorf("Bad number of fields: %d, expected"+
			" at least 3", len(fields))
	}

	result := &Bed{}

	var err error
	result.Chr = fields[0]
	result.Start, err = strconv.Atoi(fields[1])
	if err != nil {
		return nil, nil, err
	}
	result.End, err = strconv.Atoi(fields[2])
	if err != nil {
		return nil, nil, err
	}

	return result, fields[3:], nil
}

// Returns a string representation of the bed entry.
// Mainly for debugging.
func (b *Bed) String() string {
	return fmt.Sprintf("{%s|%d|%d}", b.Chr, b.Start, b.End)
}
