/* ****************************************************************************
 * Many useful functions.
 * ****************************************************************************/

// Useful functions for general purposes.
package tools

import (
	"math/rand"
	"time"
)

const Mega = 1048576

// Sets the random's seed to current time.
// Named after the good old VB function... =]
func Randomize() {
	rand.Seed(int64(time.Now().UnixNano()))
}

// Returns the index of the maximal element out of input ints.
func ArgMaxInt(a ...int) int {
	// Zero arguments not allowed
	if len(a) == 0 {
		panic("Must have at least 1 argument")
	}

	// Start with first element
	result := 0

	// Check the others
	for i := 1; i < len(a); i++ {
		if a[i] > a[result] {
			result = i
		}
	}

	return result
}

// Returns the index of the minimal element out of input ints.
func ArgMinInt(a ...int) int {
	// Zero arguments not allowed
	if len(a) == 0 {
		panic("Must have at least 1 argument")
	}

	// Start with first element
	result := 0

	// Check the others
	for i := 1; i < len(a); i++ {
		if a[i] < a[result] {
			result = i
		}
	}

	return result
}

// Returns the index of the maximal element out of input ints.
// If several maximums exist, will pick one randomly.
func ArgMaxIntR(a ...int) int {
	// Find maximum
	m := MaxInt(a...)

	// Go over array and pick one max randomly
	result := -1
	found := float64(0)
	for i := range a {
		if a[i] == m {
			found++
			if rand.Float64() < (1.0 / found) {
				result = i
			}
		}
	}

	return result
}

// Returns the index of the minimal element out of input ints.
// If several minimums exist, will pick one randomly.
func ArgMinIntR(a ...int) int {
	// Find minimum
	m := MinInt(a...)

	// Go over array and pick one max randomly
	result := -1
	found := float64(0)
	for i := range a {
		if a[i] == m {
			found++
			if rand.Float64() < (1.0 / found) {
				result = i
			}
		}
	}

	return result
}

// Returns the maximal out of input ints.
func MaxInt(a ...int) int {
	return a[ArgMaxInt(a...)]
}

// Returns the minimal out of input ints.
func MinInt(a ...int) int {
	return a[ArgMinInt(a...)]
}

// Time of last call to Tic().
var tic time.Time

// Records current time, for later calling Toc().
func Tic() {
	tic = time.Now()
}

// Returns the duration elapsed since last call of Tic().
func Toc() time.Duration {
	return time.Now().Sub(tic)
}
