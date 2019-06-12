package main

// Handles the binomial difference test.

import (
	"fmt"
	"math"
)

// Enables assertions for debugging.
const assert = false

// Performs an exact (2-sided) binomial test for the difference between the 2
// given samples.
func bindiff(n1, k1, n2, k2 int) (pvalue float64) {
	// Check input.
	if assert && (n1 < k1 || k1 < 0) {
		panic(fmt.Sprintf("Bad n1, k1: %d, %d", n1, k1))
	}
	
	if assert && (n2 < k2 || k2 < 0) {
		panic(fmt.Sprintf("Bad n2, k2: %d, %d", n2, k2))
	}

	// n=0 or both k=0 mean probability of 1.
	if n1 == 0 || n2 == 0 || (k1 == 0 && k2 == 0) {
		return 1
	}

	// Calculate probabilities.
	p1 := float64(k1) / float64(n1)
	p2 := float64(k2) / float64(n2)
	p := float64(k1 + k2) / float64(n1 + n2)
	
	// Ensure p2 is the greater.
	if p1 > p2 {
		n1, n2 = n2, n1
		k1, k2 = k2, k1
		p1, p2 = p2, p1
	}
	
	// Difference in question.
	diff := p2 - p1
	
	// Go over differences.
	pvalue = 0.0
	
	// First side - p2 is greater.
	cdfk := 0
	cdf := binoPdf(n1, 0, p)
	
	for k2 = 0; k2 <= n2; k2++ {
		// New probability.
		p2 = float64(k2) / float64(n2)
		if p2 < diff { continue }
		k1 = int( math.Floor(float64(n1) * (p2 - diff)) )
		
		// Update CDF.
		for cdfk < k1 {
			cdfk++
			cdf += binoPdf(n1, cdfk, p)
		}

		pvalue += cdf * binoPdf(n2, k2, p)
	}
	
	// Second side - p1 is greater.
	cdfk = 0
	cdf = binoPdf(n2, 0, p)
	
	for k1 = 0; k1 <= n1; k1++ {
		// New probability.
		p1 = float64(k1) / float64(n1)
		if p1 <= diff { continue }
		k2 = int( math.Ceil(float64(n2) * (p1 - diff)) - 1)
		
		// Update CDF.
		for cdfk < k2 {
			cdfk++
			cdf += binoPdf(n2, cdfk, p)
		}

		pvalue += cdf * binoPdf(n1, k1, p)
	}

	if assert && math.IsNaN(pvalue) {
		panic(fmt.Sprintf("NaN p-value for n1=%d k1=%d n2=%d k2=%d.",
				n1, k1, n2, k2))
	}
	
	return
}

// Returns P(k | n,p).
func binoPdf(n, k int, p float64) float64 {
	// Edge cases, to prevent NaN.
	if p == 0 {
		if k == 0 {
			return 1
		} else {
			return 0
		}
	}
	if p == 1 {
		if k == n {
			return 1
		} else {
			return 0
		}
	}
	
	result := math.Exp( logChoose(n, k) + float64(k) * math.Log(p) +
			float64(n - k) * math.Log(1 - p) )
	
	if assert && math.IsNaN(result) {
		panic(fmt.Sprintf("NaN PDF for (n=%d k=%d p=%f).", n, k, p))
	}
	
	return result
}

// logFactorials[i] = log(i!)
var logFactorials = []float64{0}

// Returns log(n!).
func logFactorial(n int) float64 {
	if assert && n < 0 {
		panic(fmt.Sprint("Bad n:", n))
	}
	
	for len(logFactorials) <= n {
		i := len(logFactorials)
		logFactorials = append( logFactorials,
				logFactorials[i-1] + math.Log(float64(i)) )
	}
	
	if assert && math.IsNaN(logFactorials[n]) {
		panic(fmt.Sprintf("NaN in facotorial[%d].", n))
	}
	
	return logFactorials[n]
}

// Returns log(n choose k).
func logChoose(n, k int) float64 {
	if assert && (n < 0 || k < 0 || n < k) {
		panic(fmt.Sprintf("Bad parameters: n=%d k=%d", n, k))
	}
	
	result := logFactorial(n) - logFactorial(k) - logFactorial(n - k)

	if assert && math.IsNaN(result) {
		panic(fmt.Sprintf("NaN for (%d choose %d).",
				n, k))
	}
	
	return result
}


