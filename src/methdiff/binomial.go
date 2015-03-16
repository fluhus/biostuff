package main

// Handles the binomial difference test.

import (
	"fmt"
	"math"
)

// Performs an exact (2-sided) binomial test for the difference between the 2
// given samples.
func bindiff(n1, k1, n2, k2 int) (pvalue float64) {
	// Probabilities.
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
	result := 0.0
	
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

		result += cdf * binoPdf(n2, k2, p)
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

		result += cdf * binoPdf(n1, k1, p)
	}
	
	return result
}

// Returns P(k | n,p).
func binoPdf(n, k int, p float64) float64 {
	return choose(n, k) * math.Pow(p, float64(k)) *
			math.Pow((1 - p), float64(n - k))
}

// factorials[i] = i!
var factorials = []float64{1}

// Returns n!.
func factorial(n int) float64 {
	if n < 0 {
		panic(fmt.Sprint("Bad n:", n))
	}
	
	for len(factorials) <= n {
		i := len(factorials)
		factorials = append(factorials, factorials[i-1] * float64(i))
	}
	
	return factorials[n]
}

// Returns n choose k.
func choose(n, k int) float64 {
	if n < 0 || k < 0 || n < k {
		panic(fmt.Sprintf("Bad parameters: n=%d k=%d", n, k))
	}
	
	return factorial(n) / factorial(k) / factorial(n - k)
}


