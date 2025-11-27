package rarefy

import (
	"slices"
	"testing"
)

func TestSteps(t *testing.T) {
	tests := []struct {
		sum, step int
		want      []int
	}{
		{50, 10, []int{10, 20, 30, 40, 50}},
		{55, 10, []int{10, 20, 30, 40, 50, 55}},
		{15, 10, []int{10, 15}},
		{10, 10, []int{10}},
		{7, 10, []int{7}},
	}
	for _, test := range tests {
		got := slices.Collect(steps(test.sum, test.step))
		if !slices.Equal(got, test.want) {
			t.Errorf("steps(%v,%v)=%v, want %v",
				test.sum, test.step, got, test.want)
		}
	}
}
