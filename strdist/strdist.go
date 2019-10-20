// String metric functions.
package strdist

import (
	"fmt"

	"github.com/fluhus/golgi/seqtools"
)

// TODO(amit): Add tests.
// TODO(amit): Remove bigram functions?

// Computes the edit distance for 2 byte arrays.
func EditDistance(s1, s2 []byte) int {
	// Lengths
	m, n := len(s1), len(s2)

	// Create computing buffers
	prev := make([]int, m+1)
	next := make([]int, m+1)

	// Let's go!
	for i := 0; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			// First column
			if i == 0 {
				next[j] = j
				continue
			}

			// First row
			if j == 0 {
				next[j] = i
				continue
			}

			// Not first
			insertion := prev[j] + 1
			deletion := next[j-1] + 1
			replacement := prev[j-1]

			// If mismatch, replacement adds an operation
			if s1[j-1] != s2[i-1] {
				replacement++
			}

			// Select the minimal out of the 3
			next[j] = minInt(replacement, insertion, deletion)
		}

		// Next becomes prev
		next, prev = prev, next
	}

	return prev[m]
}

// Computes the edit distance for 2 strings.
func EditDistanceStrings(s1, s2 string) int {
	return EditDistance([]byte(s1), []byte(s2))
}

// Returns the n-gram distance between 2 DNA sequences
func NgramDistance(n int, s1, s2 []byte) int {
	// Get n-gram vectors
	v1 := seqtools.NgramVector(n, s1)
	v2 := seqtools.NgramVector(n, s2)

	// Calculate distance
	result := 0
	for i := range v1 {
		if v1[i] < v2[i] {
			result += v2[i] - v1[i]
		} else {
			result += v1[i] - v2[i]
		}
	}

	// Return half, so that each difference counts as 1
	return result / 2
}

// Returns the n-gram distance between 2 DNA sequences
func NgramDistanceStrings(n int, s1, s2 string) int {
	return NgramDistance(n, []byte(s1), []byte(s2))
}

// Calculates bi-gram distance between 2 DNA sequences.
func BigramDistanceStrings(s1, s2 string) int {
	return NgramDistance(2, []byte(s1), []byte(s2))
}

// Calculates bi-gram distance between 2 DNA sequences.
func BigramDistance(s1, s2 []byte) int {
	return NgramDistance(2, s1, s2)
}

// Returns the Hamming distance between 2 byte arrays.
func HammingDistance(s1, s2 []byte) int {
	// String lengths
	l1, l2 := len(s1), len(s2)

	// Make sure s1 is the shorter
	if l1 > l2 {
		s1, s2 = s2, s1
		l1, l2 = l2, l1
	}

	// Start with the difference
	result := l2 - l1

	// Add 1 for each non matching character
	for i := range s1 {
		if s1[i] != s2[i] {
			result++
		}
	}

	return result
}

// Returns the Hamming distance between 2 strings.
func HammingDistanceStrings(s1, s2 string) int {
	return HammingDistance([]byte(s1), []byte(s2))
}

// Returns the maximal out of input ints.
func maxInt(a ...int) int {
	// Zero arguments not allowed
	if len(a) == 0 {
		// TODO(amit): Return min int.
		panic("must have at least 1 argument")
	}

	// Start with first element
	result := a[0]

	// Check the others
	for i := 1; i < len(a); i++ {
		if a[i] > result {
			result = a[i]
		}
	}

	return result
}

// Returns the minimal out of input ints.
func minInt(a ...int) int {
	// Zero arguments not allowed
	if len(a) == 0 {
		// TODO(amit): Return max int.
		panic("must have at least 1 argument")
	}

	// Start with first element
	result := a[0]

	// Check the others
	for i := 1; i < len(a); i++ {
		if a[i] < result {
			result = a[i]
		}
	}

	return result
}

// Holds alignment scores for Blast distance calculation.
// The scores should be high for "bad" operations (mismatches, gaps), and
// low for matches.
type BlastScores struct {
	Match     int
	Mismatch  int
	GapOpen   int
	GapExtend int
}

// Returns the default Blast alignment scores.
func BlastDefaultScores() BlastScores {
	return BlastScores{-2, 3, 5, 2}
}

// For Blast distance calculation.
// Holds a score and the step that led to it (match, gap1, gap2).
type blastBlock struct {
	score int
	step  int
}

// Implements the Blast algorithm for 2 byte arrays.
// The scores should be high for "bad" operations (mismatches, gaps), and
// low for matches.
func BlastDistance(s1 []byte, s2 []byte, scores BlastScores) int {
	// Step type constants
	const mm = 0 // Match/mismatch
	const g1 = 1 // Gap in sequence 1
	const g2 = 2 // Gap in sequence 2

	// Make sure s1 is the shorter
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	// Record lengths
	m, n := len(s1), len(s2)

	// Dynamic calculation arrays
	prev := make([]blastBlock, m+1)
	next := make([]blastBlock, m+1)

	// Calculate!
	for col := 0; col < n+1; col++ {
		for row := 0; row < m+1; row++ {
			// First block
			if col == 0 && row == 0 {
				next[row] = blastBlock{0, mm}
				continue
			}

			// First column
			if col == 0 {
				// Check gap type
				// If extension
				if next[row-1].step == g2 {
					next[row] =
						blastBlock{next[row-1].score + scores.GapExtend, g2}

					// If opening
				} else {
					next[row] =
						blastBlock{next[row-1].score + scores.GapOpen, g2}
				}

				continue
			}

			// First row
			if row == 0 {
				// Check gap type
				// If extension
				if prev[row].step == g1 {
					next[row] =
						blastBlock{prev[row].score + scores.GapExtend, g1}

					// If opening
				} else {
					next[row] =
						blastBlock{prev[row].score + scores.GapOpen, g1}
				}

				continue
			}

			// Not first in anything
			var gap1, gap2, match int

			// Gap 1 score
			if prev[row].step == g1 {
				gap1 = prev[row].score + scores.GapExtend
			} else {
				gap1 = prev[row].score + scores.GapOpen
			}

			// Gap 2 score
			if next[row-1].step == g2 {
				gap2 = next[row-1].score + scores.GapExtend
			} else {
				gap2 = next[row-1].score + scores.GapOpen
			}

			// Match / mismatch score
			if s1[row-1] == s2[col-1] {
				match = prev[row-1].score + scores.Match
			} else {
				match = prev[row-1].score + scores.Mismatch
			}

			// Pick the minimal
			minScore := minInt(gap1, gap2, match)

			// Check which one
			switch minScore {
			default:
				panic(fmt.Sprintf("Oh no! invalid minScore: %d", minScore))
			case match:
				next[row] = blastBlock{match, mm}
			case gap1:
				next[row] = blastBlock{gap1, g1}
			case gap2:
				next[row] = blastBlock{gap2, g2}
			}
		}

		// Switch prev & next
		prev, next = next, prev
	}

	return prev[m].score
}

// Implements the Blast algorithm for 2 strings.
// The scores should be high for "bad" operations (mismatches, gaps), and
// low for matches.
func BlastDistanceStrings(s1 string, s2 string, scores BlastScores) int {
	return BlastDistance([]byte(s1), []byte(s2), scores)
}

// Implements the Blast algorithm for 2 byte arrays.
// The scores should be high for "bad" operations (mismatches, gaps), and
// low for matches.
// *** LOCAL ***
func LocalBlastDistance(s1 []byte, s2 []byte, scores BlastScores) int {
	// Step type constants
	const mm = 0 // Match/mismatch
	const g1 = 1 // Gap in sequence 1
	const g2 = 2 // Gap in sequence 2

	// Record lengths
	m, n := len(s1), len(s2)

	// Dynamic calculation matrix
	mat := make([][]blastBlock, m+1)
	for i := range mat {
		mat[i] = make([]blastBlock, n+1)
	}

	// Best score recorder
	bestScore := 0

	// Calculate!
	for col := 0; col < n+1; col++ {
		for row := 0; row < m+1; row++ {
			// First block
			if col == 0 && row == 0 {
				mat[row][col] = blastBlock{0, mm}
				continue
			}

			// First column
			if col == 0 {
				// Check gap type
				// If extension
				if mat[row-1][col].step == g2 {
					mat[row][col] =
						blastBlock{mat[row-1][col].score + scores.GapExtend, g2}

					// If opening
				} else {
					mat[row][col] =
						blastBlock{mat[row-1][col].score + scores.GapOpen, g2}
				}

				continue
			}

			// First row
			if row == 0 {
				// Check gap type
				// If extension
				if mat[row][col-1].step == g1 {
					mat[row][col] =
						blastBlock{mat[row][col-1].score + scores.GapExtend, g1}

					// If opening
				} else {
					mat[row][col] =
						blastBlock{mat[row][col-1].score + scores.GapOpen, g1}
				}

				continue
			}

			// Not first in anything
			var gap1, gap2, match, newmatch int

			// Gap 1 score
			if mat[row][col-1].step == g1 {
				gap1 = mat[row][col-1].score + scores.GapExtend
			} else {
				gap1 = mat[row][col-1].score + scores.GapOpen
			}

			// Gap 2 score
			if mat[row-1][col].step == g2 {
				gap2 = mat[row-1][col].score + scores.GapExtend
			} else {
				gap2 = mat[row-1][col].score + scores.GapOpen
			}

			// Match / mismatch score
			if s1[row-1] == s2[col-1] {
				match = mat[row-1][col-1].score + scores.Match
				newmatch = scores.Match
			} else {
				match = mat[row-1][col-1].score + scores.Mismatch
				newmatch = scores.Mismatch
			}

			// Pick the minimal
			minScore := minInt(gap1, gap2, match, newmatch)

			// Check which one
			switch minScore {
			default:
				panic(fmt.Sprintf("Oh no! invalid minScore: %d", minScore))
			case match:
				mat[row][col] = blastBlock{match, mm}
			case newmatch:
				mat[row][col] = blastBlock{newmatch, mm}
			case gap1:
				mat[row][col] = blastBlock{gap1, g1}
			case gap2:
				mat[row][col] = blastBlock{gap2, g2}
			}

			// Update new best score
			if mat[row][col].score < bestScore {
				bestScore = mat[row][col].score
			}
		}
	}

	return bestScore
}
