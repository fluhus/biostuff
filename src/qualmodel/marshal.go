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
			text = append(text, ("# " + line)...)
		}
	}
	
	// Output lines
	test = append(text, fmt.Sprintf("%d\n", len(m.counts))...)
	
	
	return
}
