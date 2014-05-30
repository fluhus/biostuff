package learning

// A unit test for the perceptron.

import (
	"testing"
)

func Test_perceptron(t *testing.T) {
	p := NewPerceptronBiased(2)
	
	// Learn
	p.LearnInt( []int{2,2}, 1 )
	p.LearnInt( []int{1,3}, 1 )
	p.LearnInt( []int{3,1}, 1 )
	p.LearnInt( []int{-2,-2}, -1 )
	p.LearnInt( []int{-1,-3}, -1 )
	p.LearnInt( []int{-3,-1}, -1 )
	
	// Test classifications
	if y := p.ClassifyInt( []int{3,3} ); y != 1 {
		t.Errorf("bad classification: %v, expected 1", y)
	}
	
	if y := p.ClassifyInt( []int{-3,-3} ); y != -1 {
		t.Errorf("bad classification: %v, expected -1", y)
	}
	
	if y := p.ClassifyInt( []int{2,3} ); y != 1 {
		t.Errorf("bad classification: %v, expected 1", y)
	}
}
