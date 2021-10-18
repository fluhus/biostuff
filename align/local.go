package align

import "fmt"

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
