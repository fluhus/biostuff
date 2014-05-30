package learning

// General helper functions.

import (
	"fmt"
)

// Returns the dot product of x and y.
func dot(x, y []float64) float64 {
	// Check input
	if len(x) != len(y) {
		panic(fmt.Sprintf("inconsistent dimentions: x(%d) y(%d)",
				len(x), len(y)))
	}
	
	// Multiply
	result := 0.0
	for i := range x {
		result += x[i] * y[i]
	}
	return result
}

// Returns a new vector of the sum of x and y.
func add(x, y []float64) []float64 {
	// Check input
	if len(x) != len(y) {
		panic(fmt.Sprintf("inconsistent dimentions: x(%d) y(%d)",
				len(x), len(y)))
	}
	
	result := make([]float64, len(x))
	for i := range x {
		result[i] = x[i] + y[i]
	}
	
	return result
}

// Returns x multiplied by scalar a.
func multiplyScalar(x []float64, a float64) []float64 {
	result := make([]float64, len(x))
	for i := range x {
		result[i] = x[i] * a
	}
	
	return result
}

// Converts a slice of ints to a slice of floats.
func intsToFloats(ints []int) []float64 {
	result := make([]float64, len(ints))
	for i := range ints {
		result[i] = float64(ints[i])
	}
	return result
}
