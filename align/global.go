package align

import "fmt"

// Global performs global alignment on a and b and finds the highest scoring
// alignment. Returns the steps relating to a, and the alignment score.
// Time and space complexities are O(len(a)*len(b)).
//
// Uses the Needleman-Wunsch algorithm.
func Global(a, b []byte, m SubstitutionMatrix) (steps []Step, score float64) {
	s := m.toArray()
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
			blocks[i].score = blocks[i-1].score + s.get(Gap, b[bi-1])
			if bi == 1 { // New gap
				blocks[i].score += s.get(Gap, Gap)
			}
			continue
		}
		if bi == 0 {
			blocks[i].step = Deletion
			blocks[i].score = blocks[i-bn].score + s.get(a[ai-1], Gap)
			if ai == 1 { // New gap
				blocks[i].score += s.get(Gap, Gap)
			}
			continue
		}

		// Middle of matrix. Calculate the score of each possible step.
		mch := blocks[i-bn-1].score + s.get(a[ai-1], b[bi-1])
		del := blocks[i-bn].score + s.get(a[ai-1], Gap)
		if blocks[i-bn].step != Deletion {
			del += s.get(Gap, Gap)
		}
		ins := blocks[i-1].score + s.get(Gap, b[bi-1])
		if blocks[i-1].step != Insertion {
			ins += s.get(Gap, Gap)
		}
		blocks[i] = decideOnStep(mch, del, ins)
	}

	return traceAlignmentSteps(blocks, bn)
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
