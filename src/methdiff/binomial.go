package main

// Handles the binomial difference test.

import (
	"fmt"
	"math"
)

// Performs an exact binomial test for the difference between the 2 given
// samples.
func bindiff(n1, k1, n2, k2 int) (pvalue float64) {
	
}

// Returns log( P(k | n,p) )
func logBinomial(n, k int, p float64) float64 {
	return logChoose(n, k) + (float64(k) * math.Log(p)) +
			(float64(n - k) * math.Log(1 - p))
}

// logFactorials[i] = log(i!)
var logFactorials = []float64{0}

// Returns log(n!).
func logFactorial(n int) float64 {
	if n < 0 {
		panic(fmt.Sprint("Bad n:", n))
	}
	
	for len(logFactorials) <= n {
		i := len(logFactorials)
		logFactorials = append(logFactorials, logFactorials[i-1] +
				math.Log(float64(i)))
	}
	
	return logFactorials[n]
}

// Returns log(n choose k).
func logChoose(n, k int) float64 {
	if n < 0 || k < 0 || n < k {
		panic(fmt.Sprintf("Bad parameters: n=%d k=%d", n, k))
	}
	
	return logFactorial(n) - logFactorial(k) - logFactorial(n - k)
}

