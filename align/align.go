// Package align provides functionality for aligning sequences.
package align

import "fmt"

// A Step is an alignment of one character from each sequence.
//
// For example, the alignment of "blablab" and "blrblbr":
//  blablab-
//  || || |
//  blrbl-br
// Can be represented in steps as:
//  [match, match, match, match, match, deletion, match, insertion]
type Step byte

// Possible values of a step.
const (
	Match     Step = 1 // A character of A is aligned with a character of B.
	Deletion  Step = 2 // A character of A is aligned with a gap.
	Insertion Step = 3 // A character of B is aligned with a gap.
)

// String returns the name of the step.
func (s Step) String() string {
	switch s {
	case Match:
		return "match"
	case Deletion:
		return "deletion"
	case Insertion:
		return "insertion"
	default:
		panic(fmt.Sprintf("unknown step value: %d", s))
	}
}

// Gap is the byte value for a gap in a substitution matrix.
const Gap = 255

// A SubstitutionMatrix is a map from character pair p to the score of aligning p[0]
// with p[1]. Character 255 is reserved for gap. The pair [Gap,Gap] represents the
// general cost of opening a gap (gap-open). A matrix may be asymmetrical.
type SubstitutionMatrix map[[2]byte]int

// Get returns the value of aligning a with b. Panics if the pair [a,b] is not in the
// matrix.
func (m SubstitutionMatrix) Get(a, b byte) int {
	s, ok := m[[2]byte{a, b}]
	if !ok {
		panic(fmt.Sprintf("pair (%v,%v) is not in the substitution-matrix",
			a, b))
	}
	return s
}

// Symmetrical returns a copy of m. The copy also contains each pair flipped, mapped
// to the pair's original value.
func (m SubstitutionMatrix) Symmetrical() SubstitutionMatrix {
	result := SubstitutionMatrix{}
	for k, v := range m {
		result[k] = v
		flip := [2]byte{k[1], k[0]}
		if k[0] != k[1] {
			if v2, ok := m[flip]; ok && v2 != v {
				panic(fmt.Sprintf("conflicting pair: (%v,%v)=%v, (%v,%v)=%v",
					k[0], k[1], v, k[1], k[0], v2))
			}
		}
		result[flip] = v
	}
	return result
}

// A single element in the dynamic-programming table of the alignment algorithm.
type block struct {
	score int
	step  Step
}

// Global performs global alignment on a and b and finds the highest scoring
// alignment. Returns the steps where each step is on a single character, relating
// to a. Time and space complexities are O(len(a)*len(b)).
//
// Uses the Needleman-Wunsch algorithm.
func Global(a, b []byte, m SubstitutionMatrix) ([]Step, int) {
	an, bn := len(a)+1, len(b)+1
	blocks := make([]block, an*bn)
	for i := range blocks {
		ai, bi := i/bn, i%bn

		// Edges of the matrix.
		if ai == 0 && bi == 0 {
			continue
		}
		if ai == 0 {
			blocks[i].step = Insertion
			blocks[i].score = blocks[i-1].score + m.Get(Gap, b[bi-1])
			if bi == 1 { // New gap
				blocks[i].score += m.Get(Gap, Gap)
			}
			continue
		}
		if bi == 0 {
			blocks[i].step = Deletion
			blocks[i].score = blocks[i-bn].score + m.Get(a[ai-1], Gap)
			if ai == 1 { // New gap
				blocks[i].score += m.Get(Gap, Gap)
			}
			continue
		}

		// Middle of matrix. Calculate the score of each possible step.
		mch := blocks[i-bn-1].score + m.Get(a[ai-1], b[bi-1])
		del := blocks[i-bn].score + m.Get(a[ai-1], Gap)
		if blocks[i-bn].step != Deletion {
			del += m.Get(Gap, Gap)
		}
		ins := blocks[i-1].score + m.Get(Gap, b[bi-1])
		if blocks[i-1].step != Insertion {
			ins += m.Get(Gap, Gap)
		}
		blocks[i] = decideOnStep(mch, del, ins)
	}

	return traceAlignmentSteps(blocks, bn)
}

// Returns a block with the highest scoring step.
func decideOnStep(mch, del, ins int) block {
	if mch >= del && mch >= ins {
		return block{score: mch, step: Match}
	} else if del >= ins {
		return block{score: del, step: Deletion}
	} else {
		return block{score: ins, step: Insertion}
	}
}

// Reproduces the alignment steps that lead to the final highest score.
// Returns the steps and their score.
func traceAlignmentSteps(blocks []block, bn int) ([]Step, int) {
	steps := make([]Step, 0, bn+len(blocks)/bn)
	i := len(blocks) - 1
	for i > 0 {
		steps = append(steps, blocks[i].step)
		switch blocks[i].step {
		case Match:
			i -= bn + 1
		case Deletion:
			i -= bn
		case Insertion:
			i -= 1
		}
	}
	if i < 0 {
		panic(fmt.Sprintf("bad i: %v, expected 0", i))
	}
	// Reverse steps.
	for i := 0; i < len(steps)/2; i++ {
		steps[i], steps[len(steps)-1-i] = steps[len(steps)-1-i], steps[i]
	}
	return steps, blocks[len(blocks)-1].score
}
