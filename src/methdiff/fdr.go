package main

// FDR correction for p-values.

import (
	"sort"
)

// Returns a slice of corrected q-values, respective to the input p-value slice.
func fdr(pvals []float64) []float64 {
	// Sort p-values.
	result := make([]float64, len(pvals))
	copy(result, pvals)
	sort.Sort(sort.Float64Slice(result))
	
	// Register q-values in a map, so that duplicates fall in the same entry.
	qvals := make(map[float64]float64)
	for i, v := range result {
		qval := v * float64(len(pvals)) / float64(i + 1)
		if qval > 1 { qval = 1 }
		qvals[v] = qval
	}
	
	// Assign q-values in result array.
	for i, v := range pvals {
		result[i] = qvals[v]
	}
	
	return result
}


