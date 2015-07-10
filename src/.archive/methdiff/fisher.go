package main

// Temporary file for Fisher's exact test.

import (
	"math"
)

func fisher(n1, k1, n2, k2 int) (pvalue float64) {
	// Convert to Wikipedia's notations.
	a := k1
	b := n1 - k1
	c := k2
	d := n2 - k2

	logP := logFactorial(a + b) + logFactorial(a + c) + logFactorial(d + b) +
			logFactorial(d + c) - logFactorial(a) - logFactorial(b) -
			logFactorial(c) - logFactorial(d) - logFactorial(a+b+c+d)

	return math.Exp(logP)
}
