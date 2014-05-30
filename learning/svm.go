package learning

// Implementation of a soft SVM.

import (
	"fmt"
)

// A single soft SVM. Records learned data and provides classification
// functionality.
type Svm struct {
	theta  []float64   // Learning vector
	w      []float64   // Classification vector
	biased   bool      // Is this SVM biased
	lambda   float64   // Lambda pace parameter
	t        float64   // Number of learned samples so far
}

// Returns a new unbiased SVM with the given dimention and lambda.
func NewSvmUnbiased(dimention int, lambda float64) *Svm {
	// BUG( ) TODO check dimention
	return &Svm{ make([]float64, dimention),
			make([]float64, dimention), false, lambda, 0 }
}

// Returns a new biased SVM with the given dimention and lambda.
func NewSvmBiased(dimention int, lambda float64) *Svm {
	// BUG( ) TODO check dimention
	return &Svm{ make([]float64, dimention + 1),
			make([]float64, dimention + 1), true, lambda, 0 }
}

// Learns the given vector.
// y is the classification of x, and should be either 1 or -1.
func (s *Svm) LearnFloat(x []float64, y int) {
	// Add bias
	bias := 0
	xBiased := x
	if s.biased {
		xBiased = append([]float64{1}, x...)
		bias = 1
	}

	// Input check
	if len(xBiased) != len(s.w) {
		panic(fmt.Sprintf("inconsistent dimention: %d, expected %d",
				len(x), len(s.w) - bias))
	}
	
	if y != 1 && y != -1 {
		panic(fmt.Sprintf("bad y: %d, expected 1 or -1", y))
	}
	
	// Check hinge loss
	s.t++
	w := multiplyScalar(s.theta, 1.0 / s.lambda / s.t)
	s.w = add(s.w, w)
	if float64(y) * dot(xBiased, w) < 1 {
		// Add x to theta
		s.theta = add(s.theta, multiplyScalar(xBiased, float64(y)))
	}
}

// Classifies the given vector (1 or -1).
func (s *Svm) ClassifyFloat(x []float64) int {
	// Add bias
	bias := 0
	xBiased := x
	if s.biased {
		xBiased = append([]float64{1}, x...)
		bias = 1
	}
	
	// Input check
	if len(xBiased) != len(s.w) {
		panic(fmt.Sprintf("inconsistent dimention: %d, expected %d",
				len(x), len(s.w) - bias))
	}

	dotProd := dot(xBiased, s.w)
	
	switch {
	case dotProd == 0:
		return 0
	case dotProd > 0:
		return 1
	case dotProd < 0:
		return -1
	default:
		panic("how did I get here?!")
	}
}

// Learns the given vector.
// 'y' is the classification of 'x', and should be either 1 or -1.
func (s *Svm) LearnInt(x []int, y int) {
	s.LearnFloat(intsToFloats(x), y)
}

// Classifies the given vector (1 or -1).
func (s *Svm) ClassifyInt(x []int) int {
	return s.ClassifyFloat(intsToFloats(x))
}

// Returns a copy of the separator vector.
func (s *Svm) W() []float64 {
	return multiplyScalar(s.w, 1.0 / s.t)
}
