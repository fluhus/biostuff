package qualmodel

// This file handles marshaling and unmarshaling of quality models.

import (
	"strings"
	"fmt"
)

func (m *Model) MarshalText() (text []byte, err error) {
	// Output comment
	if m.comment != "" {
		split := strings.Split(m.comment, "\n")
		for _,line := range split {
			text = append(text, ("# " + line + "\n")...)
		}
	}
	
	// Output lines
	text = append(text, fmt.Sprintf("%d\n", len(m.counts))...)
	for i := range m.counts {
		// Output line length
		text = append(text, fmt.Sprint(len(m.counts[i]))...)
		
		// Add counts
		for j := range m.counts[i] {
			text = append(text, fmt.Sprintf(" %d", m.counts[i][j])...)
		}
		
		text = append(text, '\n')
	}
	
	return
}
