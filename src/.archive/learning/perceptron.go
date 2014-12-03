package learning

// An implementation of the perceptron learning algorithm.

import (
	"fmt"
)

// A single perceptron. Records learned data and provides classification
// functionality.
type Perceptron struct {
	w []float64   // Classification hyperplane
	biased bool   // True if biased
	updateCoefficients [3]float64  // Coefficients of update for false
	                               // positive(0) and false negative(2)
}

// Returns a new unbiased perceptron with the given dimention.
func NewPerceptronUnbiased(dimention int,
		falsePositiveWeight, falseNegativeWeight float64) *Perceptron {
	// BUG( ) TODO check dimention
	return &Perceptron{ make([]float64, dimention), false,
			[3]float64{falsePositiveWeight, 0, falseNegativeWeight} }
}

// Returns a new biased perceptron with the given dimention.
func NewPerceptronBiased(dimention int,
		falsePositiveWeight, falseNegativeWeight float64) *Perceptron {
	// BUG( ) TODO check dimention
	return &Perceptron{ make([]float64, dimention + 1), true,
			[3]float64{falsePositiveWeight, 0, falseNegativeWeight} }
}

// Learns the given vector, if it maps incorrectly.
// 'y' is the classification of 'x', and should be either 1 or -1.
func (p *Perceptron) LearnFloat(x []float64, y int) {
	// Add bias
	bias := 0
	xBiased := x
	if p.biased {
		xBiased = append([]float64{1}, x...)
		bias = 1
	}

	// Input check
	if len(xBiased) != len(p.w) {
		panic(fmt.Sprintf("inconsistent dimention: %d, expected %d",
				len(x), len(p.w) - bias))
	}
	
	if y != 1 && y != -1 {
		panic(fmt.Sprintf("bad y: %d, expected 1 or -1", y))
	}
	
	// Check if classifies ok
	if y * p.ClassifyFloat(x) <= 0 {
		// Add x to w
		for i := range xBiased {
			p.w[i] += float64(y) * xBiased[i] *
					float64(p.updateCoefficients[y+1])
		}
	}
}

// Classifies the given vector (1 or -1).
func (p *Perceptron) ClassifyFloat(x []float64) int {
	// Add bias
	bias := 0
	xBiased := x
	if p.biased {
		xBiased = append([]float64{1}, x...)
		bias = 1
	}
	
	// Input check
	if len(xBiased) != len(p.w) {
		panic(fmt.Sprintf("inconsistent dimention: %d, expected %d",
				len(x), len(p.w) - bias))
	}

	dotProd := dot(xBiased, p.w)
	
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

// Learns the given vector, if it maps incorrectly.
// 'y' is the classification of 'x', and should be either 1 or -1.
func (p *Perceptron) LearnInt(x []int, y int) {
	p.LearnFloat(intsToFloats(x), y)
}

// Classifies the given vector (1 or -1).
func (p *Perceptron) ClassifyInt(x []int) int {
	return p.ClassifyFloat(intsToFloats(x))
}

// Returns a copy of the separator vector.
func (p *Perceptron) W() []float64 {
	w := make([]float64, len(p.w))
	copy(w, p.w)
	return w
}

// Sets the classification of this perceptro.
func (p *Perceptron) SetW(w []float64) {
	// BUG( ) should I keep this function?
	
	// Input check
	if len(w) != len(p.w) {
		panic(fmt.Sprintf("inconsistent dimention: %d, expected %d",
				len(w), len(p.w)))
	}
	
	p.w = w
}




