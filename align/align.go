// Package align provides functionality for aligning sequences.
//
// # Alignment Steps
//
// Alignments are represented as Steps. Each step when aligning two sequences can be
// either taking a character from each sequence and aligning them together, or
// taking only one character and aligning it with a gap. The match step covers the
// case of a mismatch as well.
//
// For example, the alignment of "blablab" and "blrblbr":
//
//	blablab-
//	|| || |
//	blrbl-br
//
// Can be represented in steps as:
//
//	[match, match, match, match, match, deletion, match, insertion]
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
		panic(fmt.Sprintf("pair (%d %s, %d %s) is not in the substitution-matrix",
			a, charOrGap(a), b, charOrGap(b)))
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
