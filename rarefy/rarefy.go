// Package rarefy provides rarefaction curve generation.
package rarefy

import (
	"fmt"
	"iter"
	"math"
	"reflect"
	"slices"

	"github.com/fluhus/gostuff/bits"
	"github.com/fluhus/gostuff/gnum"
	"github.com/fluhus/gostuff/snm"
	"golang.org/x/exp/constraints"
)

// TODO(amit): Add support for min number of reads per species.

const (
	// Use new calculation that doesn't do simulations,
	// instead it calculates the expected numbers directly.
	useNonIterativeRarefy = true

	// Chunk assignment before iterating over it.
	useChunkedIteration = true
)

// Rarefy returns a rarefaction curve for the given read counts.
// Output is one slice of x values (number of reads)
// and one slice of corresponding y values (species count).
//
// readCounts has the number of reads per species.
// step is the x-axis interval length.
// nperms is the number of permutations to average on.
func Rarefy(readCounts []int, step, nperms int) ([]int, []int) {
	if step < 1 {
		panic(fmt.Sprintf("bad step: %v", step))
	}
	if useNonIterativeRarefy {
		return rarefy2(readCounts, step)
	}
	// Dispatch the smallest possible integer, to reduce memory footprint.
	if len(readCounts) < 1<<8 {
		return rarefy[uint8](readCounts, step, nperms)
	}
	if len(readCounts) < 1<<16 {
		return rarefy[uint16](readCounts, step, nperms)
	}
	if len(readCounts) < 1<<32 {
		return rarefy[uint32](readCounts, step, nperms)
	}
	return rarefy[uint64](readCounts, step, nperms)
}

// Generic implementation.
func rarefy[T constraints.Integer](readCounts []int, step, nperms int) ([]int, []int) {
	assn := make([]T, 0, gnum.Sum(readCounts))
	i := T(0)
	for _, c := range readCounts {
		for range c {
			assn = append(assn, i)
		}
		i++
		if i == 0 { // Assertion that integer size is enough.
			panic(fmt.Sprintf("species count overflow: counting %v species with %v",
				len(readCounts), reflect.TypeOf(i)))
		}
	}
	snm.Shuffle(assn)

	found := make([]byte, (len(readCounts)+7)/8)
	var xx, yy, yyy []int
	for p := range nperms {
		clear(found)
		yy = yy[:0]
		if useChunkedIteration {
			i := 0
			for chunk := range slices.Chunk(assn, step) {
				for _, a := range chunk {
					bits.Set(found, a, true)
				}
				i += len(chunk)
				if p == 0 {
					xx = append(xx, i)
				}
				// TODO(amit): Instead of Sum I can keep a counter.
				// Should be faster.
				yy = append(yy, bits.Sum(found))
			}
		} else {
			for i, a := range assn {
				bits.Set(found, a, true)
				if (i+1)%step == 0 {
					if p == 0 {
						xx = append(xx, i+1)
					}
					// TODO(amit): Instead of Sum I can keep a counter.
					// Should be faster.
					yy = append(yy, bits.Sum(found))
				}
			}
			if p == 0 {
				xx = append(xx, len(assn))
			}
			yy = append(yy, bits.Sum(found))
		}
		if p == 0 {
			yyy = make([]int, len(yy))
		}
		gnum.Add(yyy, yy)

		if nperms > 1 {
			assn = chunkShuffle(assn)
		}
	}
	for i, y := range yyy {
		yyy[i] = gnum.Idiv(y, nperms)
	}
	return xx, yyy
}

// Shuffles a slice by chunks, rather than by single elements.
// Faster than regular shuffle and random enough for this use-case.
func chunkShuffle[T any](a []T) []T {
	// Using sqrt(n) chunks.
	nchunks := int(math.Round(math.Sqrt(float64(len(a)))))
	chunks := make([][]T, nchunks)
	n := len(a)
	for i := range nchunks {
		from, to := i*n/nchunks, (i+1)*n/nchunks
		chunks[i] = a[from:to]
	}
	snm.Shuffle(chunks)
	return slices.Concat(chunks...)
}

// A non-random implementation.
// Calculates expected species counts directly instead
// of simulating.
func rarefy2(readCounts []int, step int) ([]int, []int) {
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
