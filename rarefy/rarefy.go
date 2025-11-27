// Package rarefy provides rarefaction curve generation.
package rarefy

import (
	"fmt"
	"iter"
	"math"

	"github.com/fluhus/gostuff/gnum"
	"github.com/fluhus/gostuff/snm"
)

// TODO(amit): Add support for min number of reads per species.

// Rarefy returns a rarefaction curve for the given read counts.
// Output is one slice of x values (number of reads)
// and one slice of corresponding y values (species count).
//
// readCounts has the number of reads per species.
// step is the x-axis interval length.
func Rarefy(readCounts []int, step int) ([]int, []int) {
	if step < 1 {
		panic(fmt.Sprintf("bad step: %v", step))
	}

	sum := gnum.Sum(readCounts)
	var xx, yy []int
	lfAll := gnum.LogFactorial(sum)
	lfCounts := snm.SliceToSlice(readCounts, func(i int) float64 {
		return gnum.LogFactorial(sum - i)
	})
	for x := range steps(sum, step) {
		xx = append(xx, x)
		y := 0.0
		lfX := gnum.LogFactorial(sum - x)
		for i, c := range readCounts {
			p := 0.0
			if sum >= c+x {
				// Hypergeometric probability for no species c in the first
				// x reads.
				p = math.Exp(lfCounts[i] + lfX - lfAll - gnum.LogFactorial(sum-x-c))
			}
			// 1-p because we want the inverse probability (at least 1).
			y += 1 - p
		}
		yy = append(yy, int(math.Round(y)))
	}
	return xx, yy
}

// Yields multiples of step that are less than or equal
// to sum, and also sum itself.
func steps(sum, step int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := step; i <= sum; i += step {
			if !yield(i) {
				return
			}
		}
		if sum%step != 0 {
			yield(sum)
		}
	}
}
