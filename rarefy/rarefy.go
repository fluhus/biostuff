// Package rarefy provides rarefaction curve generation.
package rarefy

import (
	"fmt"
	"math"
	"reflect"
	"slices"

	"github.com/fluhus/gostuff/bits"
	"github.com/fluhus/gostuff/gnum"
	"github.com/fluhus/gostuff/snm"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
)

// TODO(amit): Add support for min number of reads per species.

// Rarefy returns a rarefaction curve for the given read counts.
// Output is one slice of x values (number of reads)
// and one slice of corresponding y values (species count).
//
// readCounts has the number of reads per species.
// step is the x-axis interval length.
// nperms is the number of permutations to average on.
func Rarefy(readCounts map[string]int, step, nperms int) ([]int, []int) {
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
func rarefy[T constraints.Integer](readCounts map[string]int, step, nperms int) ([]int, []int) {
	assn := make([]T, 0, gnum.Sum(maps.Values(readCounts)))
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
