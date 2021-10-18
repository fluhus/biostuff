// Package align provides functionality for aligning sequences.
//
// Alignment Steps
//
// Alignments are represented as Steps. Each step when aligning two sequences can be
// either taking a character from each sequence and aligning them together, or
// taking only one character and aligning it with a gap. The match step covers the
// case of a mismatch as well.
//
// For example, the alignment of "blablab" and "blrblbr":
//  blablab-
//  || || |
//  blrbl-br
// Can be represented in steps as:
//  [match, match, match, match, match, deletion, match, insertion]
package align

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

// A Step is an alignment of one character from each sequence.
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
type SubstitutionMatrix map[[2]byte]float64

// Get returns the value of aligning a with b. Either argument may have the value
// Gap. Panics if the pair [a,b] is not in the matrix.
func (m SubstitutionMatrix) Get(a, b byte) float64 {
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
	score float64
	step  Step
}

// Global performs global alignment on a and b and finds the highest scoring
// alignment. Returns the steps relating to a, and the alignment score.
// Time and space complexities are O(len(a)*len(b)).
//
// Uses the Needleman-Wunsch algorithm.
func Global(a, b []byte, m SubstitutionMatrix) (steps []Step, score float64) {
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

// Local performs local alignment on a and b and finds the highest scoring
// alignment. Returns the steps relating to a, ai and bi as the start positions of
// the local alignment in a and b respectively, and the alignment score.
// Time and space complexities are O(len(a)*len(b)).
//
// Uses the Smith-Waterman algorithm.
func Local(a, b []byte, m SubstitutionMatrix) (
	steps []Step, ai, bi int, score float64) {
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
			if blocks[i].score < 0 {
				blocks[i] = block{0, 0}
			}
			continue
		}
		if bi == 0 {
			blocks[i].step = Deletion
			blocks[i].score = blocks[i-bn].score + m.Get(a[ai-1], Gap)
			if ai == 1 { // New gap
				blocks[i].score += m.Get(Gap, Gap)
			}
			if blocks[i].score < 0 {
				blocks[i] = block{0, 0}
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
		if blocks[i].score < 0 {
			blocks[i] = block{0, 0}
		}
	}

	steps, i, score := traceAlignmentStepsLocal(blocks, bn)
	return steps, i/bn - 1, i%bn - 1, score
}

// Returns a block with the highest scoring step.
func decideOnStep(mch, del, ins float64) block {
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
func traceAlignmentSteps(blocks []block, bn int) ([]Step, float64) {
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

// Reproduces the alignment steps that lead to the final highest score.
// Returns the steps and their score.
func traceAlignmentStepsLocal(blocks []block, bn int) ([]Step, int, float64) {
	imax := argmax(blocks)
	steps := make([]Step, 0, bn+len(blocks)/bn)
	i := imax
	last := i
	for i > 0 {
		if blocks[i].score < 0 {
			panic(fmt.Sprintf("bad score at i=%d: %f", i, blocks[i].score))
		}
		if blocks[i].score == 0 {
			break
		}
		last = i
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
	// Check placed here and not in the begining to go allow checks.
	if blocks[imax].score == 0 {
		return nil, 0, 0
	}
	// Reverse steps.
	for i := 0; i < len(steps)/2; i++ {
		steps[i], steps[len(steps)-1-i] = steps[len(steps)-1-i], steps[i]
	}
	return steps, last, blocks[imax].score
}

// Returns the index of the highest score block.
func argmax(blocks []block) int {
	imax := 0
	for i, b := range blocks {
		if b.score > blocks[imax].score {
			imax = i
		}
	}
	return imax
}

// GoString implements the fmt.GoStringer interface.
func (m SubstitutionMatrix) GoString() string {
	buf := &strings.Builder{}
	fmt.Fprintln(buf, "SubstitutionMatrix{")
	var sorted [][]byte
	for k := range m {
		sorted = append(sorted, []byte{k[0], k[1]})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return bytes.Compare(sorted[i], sorted[j]) < 0
	})
	for _, k := range sorted {
		fmt.Fprintf(buf, "{%s,%s}:%v,\n",
			charOrGap(k[0]), charOrGap(k[1]), m.Get(k[0], k[1]))
	}
	fmt.Fprintln(buf, "}")
	return buf.String()
}

// Returns a quoted char, or the constant Gap for a gap.
func charOrGap(c byte) string {
	if c == Gap {
		return "Gap"
	}
	return fmt.Sprintf("%q", c)
}
