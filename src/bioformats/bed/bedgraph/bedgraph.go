// Handles bed-graph file representation and parsing.
package bedgraph

import (
	"strings"
	"fmt"
	"strconv"
)

// A simple genomic region notation.
type BedGraph struct {
	Chr string
	Start int
	End int
	Value float64
}

// Parses a single bed-graph line. Keeps chromosome, start and end in the bed-
// graph object. All other fields are returned in a string array. Returns a
// non-nil error if couldn't parse.
func Parse(s string) (*BedGraph, []string, error) {
	// Split
	fields := strings.Split(s, "\t")
	if len(fields) < 4 {
		return nil, nil, fmt.Errorf("Bad number of fields: %d, expected" +
				" at least 4", len(fields))
	}
	
	result := &BedGraph{}
	
	var err error
	result.Chr = fields[0]
	result.Start, err = strconv.Atoi(fields[1])
	if err != nil { return nil, nil, err }
	result.End, err = strconv.Atoi(fields[2])
	if err != nil { return nil, nil, err }
	result.Value, err = strconv.ParseFloat(fields[3], 64)
	if err != nil { return nil, nil, err }
	
	return result, fields[4:], nil
}

// Returns a string representation of the bed-graph entry.
// Mainly for debugging.
func (b *BedGraph) String() string {
	return fmt.Sprintf("{%s|%d|%d|%f}", b.Chr, b.Start, b.End, b.Value)
}

