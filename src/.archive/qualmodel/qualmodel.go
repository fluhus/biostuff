// Handles fastq quality modeling.
package qualmodel

import (
	"fmt"
	"math/rand"
)

// Models quality scores for fastq simulation.
type Model struct {
	counts  [][]int  // [x][y] x: position along the read, y: quality score
	comment string   // for data documentation
}

// Creates a new model, according to the given counts.
//
// counts: 1st index is position along the read, 2nd index is the quality
// score. 'counts' will be copied, so modifications after calling 'New' will
// not affect the created model.
func New(counts [][]int) *Model {
	// Copy counts
	countsCopy := make([][]int, len(counts))
	
	for i := range counts {
		countsCopy[i] = make([]int, len(counts[i]))
		copy(countsCopy[i], counts[i])
	}
	
	return &Model{countsCopy, ""}
}

// Creates a new model, according to the given counts and with a comment
// that can be read later.
//
// counts: 1st index is position along the read, 2nd index is the quality
// score. 'counts' will be copied, so modifications after calling 'new' will
// not affect the created model.
//
// comment: may include any character and new lines.
func NewWithComment(counts [][]int, comment string) *Model {
	result := New(counts)
	result.comment = comment
	
	return result
}

// Returns the length of the read that fits this model.
func (m *Model) Len() int {
	return len(m.counts)
}

// Returns a random quality score at the given position.
// Panics if position is out of read-length bounds.
func (m *Model) Qual(position int) int {
	// Check input
	if position < 0 || position >= len(m.counts) {
		panic(fmt.Sprint("bad position:", position))
	}
	
	// Create quality
	sum := 0
	selected := -1
	counts := m.counts[position]
	for i := range counts {
		// Nothing to select (prevent division by 0)
		if sum == 0 && counts[i] == 0 {
			continue
		}
		
		sum += counts[i]
		if rand.Float64() < float64(counts[i]) / float64(sum) {
			selected = i
		}
	}
	
	// No selectable quality
	if sum == 0 {
		panic(fmt.Sprint("no non zero elements at position", position))
	}
	
	return selected
}

// Generates an array of qualities, in length of the model.
func (m *Model) Quals() []int {
	result := make([]int, len(m.counts))
	
	for i := range result {
		result[i] = m.Qual(i)
	}
	
	return result
}

// Returns this model's comment.
func (m *Model) Comment() string {
	return m.comment
}





