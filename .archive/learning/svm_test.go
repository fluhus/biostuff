package learning

// Unit test for SVM.

import (
	"testing"
)

func Test_svm(t *testing.T) {
	s := NewSvmBiased(2, 1)
	
	// Learn
	s.LearnInt( []int{2,2}, 1 )
	s.LearnInt( []int{1,3}, 1 )
	s.LearnInt( []int{3,1}, 1 )
	s.LearnInt( []int{-2,-2}, -1 )
	s.LearnInt( []int{-1,-3}, -1 )
	s.LearnInt( []int{-3,-1}, -1 )
	
	// Test classifications
	if y := s.ClassifyInt( []int{3,3} ); y != 1 {
		t.Errorf("bad classification: %v, expected 1", y)
	}
	
	if y := s.ClassifyInt( []int{-3,-3} ); y != -1 {
		t.Errorf("bad classification: %v, expected -1", y)
	}
	
	if y := s.ClassifyInt( []int{2,3} ); y != 1 {
		t.Errorf("bad classification: %v, expected 1", y)
	}
	
	// t.Errorf("w=%v", s.W())
}
