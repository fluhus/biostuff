package bedgraph

// Bed text scanning interface.

import (
	"bufio"
	"io"
)

// Scans bed-graph entries from a stream. Ignores header if exists.
type Scanner struct {
	scanner *bufio.Scanner
	bed *BedGraph
	err error
	text string      // The parsed line as is
	fields []string  // Rest of the bed-graph line (extra fields)
	first bool       // Is next line the first line
	stopped bool     // Did we stop scanning
}

// Returns a new scanner that reads from the given stream.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{ bufio.NewScanner(r), nil, nil, "", nil, true, false }
}

// Returns the last entry parsed by Scan().
func (s *Scanner) Bed() *BedGraph {
	return s.bed
}

func (s *Scanner) Fields() []string {
	return s.fields
}

func (s *Scanner) Text() string {
	return s.text
}

func (s *Scanner) Err() error {
	return s.err
}

// Scans the next line from a bed-graph file. The parsed object can be retreived
// by calling BedGraph(). Returns true if and only if a line was successfully
// parsed. After returning false, the Err() method will return the relevant
// error, except in EOF where the error will be nil.
func (s *Scanner) Scan() bool {
	// If already stopped
	if s.stopped {
		return false
	}

	// Scan next line
	if !s.scanner.Scan() {
		s.err = s.scanner.Err()
		s.stopped = true
		return false
	}
	
	s.text = s.scanner.Text()
	s.bed, s.fields, s.err = Parse(s.scanner.Text())
	
	// Parsing error
	if s.err != nil {
		// First line may be a header, so skip it if error
		if s.first {
			s.first = false
			return s.Scan()
		}
		
		// Not first -> scanning failed
		s.stopped = true
		return false
	}
	
	// No error
	s.first = false
	return true
}

