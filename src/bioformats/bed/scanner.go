package bed

// Bed text scanning interface.

import (
	"bufio"
	"io"
)

// Scans bed entries from a stream. Ignores header if exists.
type Scanner struct {
	scanner *bufio.Scanner
	current *Bed
	err error
	first bool    // Is next line the first line
	stopped bool  // Did we stop scanning
}

// Returns a new scanner that reads from the given stream.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{ bufio.NewScanner(r), nil, nil, true, false }
}

// Returns the last entry parsed by Scan().
func (s *Scanner) Current() *Bed {
	return s.current
}

func (s *Scanner) Err() error {
	return s.err
}

// Scans the next line from a bed file. The parsed object can be retreived by
// calling Bed(). Returns true if and only if a line was successfully parsed.
// After returning false, the Err() method will return the relevant error,
// except in EOF where the error will be nil.
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
	
	s.current, s.err = Parse(s.scanner.Text())
	
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

