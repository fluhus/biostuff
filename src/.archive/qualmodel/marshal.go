package qualmodel

// This file handles marshaling and unmarshaling of quality models.

import (
	"encoding/json"
)

// Used for marshaling.
type modelMarshaler struct {
	Comment string
	Counts  [][]int
}

// Marshals a model into JSON.
func (m *Model) MarshalText() ([]byte, error) {
	marshaler := modelMarshaler{m.comment, m.counts}
	return json.MarshalIndent(marshaler, "", "\t")
}

// Unmarshals a model from JSON. Will not change if err != nil.
func (m *Model) UnmarshalText(data []byte) error {
	marshaler := &modelMarshaler{}
	err := json.Unmarshal(data, marshaler)
	
	if err != nil {
		return err
	}
	
	m.counts = marshaler.Counts
	m.comment = marshaler.Comment
	return nil
}

