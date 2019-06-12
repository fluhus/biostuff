package learning

// Unit test for helper functions.

import (
	"testing"
)

func Test_dot(t *testing.T) {
	a := []float64{1,2,3}
	b := []float64{4,5,6}
	d := dot(a, b)
	if d != 4 + 10 + 18 {
		t.Errorf("bad result: %d, expected %d", d, 4 + 10 + 18)
	}
	
	a = []float64{-1,5}
	b = []float64{4,-10}
	d = dot(a, b)
	if d != -54 {
		t.Errorf("bad result: %d, expected %d", d, -54)
	}
}

func Test_add(t *testing.T) {
	a := []float64{1,2,3}
	b := []float64{4,5,6}
	s := add(a, b)
	
	if s[0] != 5 || s[1] != 7 || s[2] != 9 {
		t.Errorf("bad result: %v, expected %v", s, []float64{5,7,9})
	}
}

func Test_scalar(t *testing.T) {
	a := []float64{1,2,3}
	s := multiplyScalar(a, 3)
	
	if s[0] != 3 || s[1] != 6 || s[2] != 9 {
		t.Errorf("bad result: %v, expected %v", s, []float64{3,6,9})
	}
}
