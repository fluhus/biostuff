package qualmodel

// A unit test for model.

import (
	"testing"
	"math"
	"math/rand"
	"time"
)

func TestQual(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// Create a model
	counts := [][]int {
		[]int {0, 0, 4, 0, 0},
		[]int {10, 0, 0, 0, 10},
		[]int {0, 2, 4, 0, 0},
		[]int {0, 0, 4, 2, 0},
	}
	
	model := New(counts)
	
	// Create empiric counts
	counts2 := make([][]int, len(counts))
	for i := range counts2 {
		counts2[i] = make([]int, len(counts[i]))
	}
	
	for i := 0; i < 100000; i++ {
		for j := range counts2 {
			counts2[j][model.Qual(j)]++
		}
	}
	
	// Test counts
	for i := range counts {  // for each position
		// Sum of position
		sum1 := 0  // expected
		sum2 := 0  // empiric
		for j := range counts[i] {
			sum1 += counts[i][j]
			sum2 += counts2[i][j]
		}
		
		for j := range counts[i] {  // for each score
			ratio1 := float64(counts[i][j]) / float64(sum1)
			ratio2 := float64(counts2[i][j]) / float64(sum2)
			
			diff := math.Abs(ratio1 - ratio2)
			if diff > 0.01 {
				t.Errorf("ratios are too different in [%d][%d]:" +
						" expected %f actual %f",
						i, j, ratio1, ratio2)
			}
		}
	}
}

func TestMarshal(t *testing.T) {
	// Create a model
	counts := [][]int {
		[]int {0, 0, 4, 0, 0},
		[]int {10, 0, 0, 0, 10},
		[]int {0, 2, 4, 0, 0},
		[]int {0, 0, 4, 2, 0},
	}
	
	model := NewWithComment(counts, "Hello\nWorld")
	bytes, _ := model.MarshalText()
	str := string(bytes)
	
	expected := "# Hello\n# World\n4\n5 0 0 4 0 0\n5 10 0 0 0 10\n" +
			"5 0 2 4 0 0\n5 0 0 4 2 0\n"
			
	if str != expected {
		t.Errorf("expected:\n%s\nactual:\n%s", expected, str)
	}
}


