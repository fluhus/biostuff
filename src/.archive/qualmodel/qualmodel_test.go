package qualmodel

// A unit test for model.

import (
	"testing"
	"math"
	"math/rand"
	"time"
)

func TestQualGeneration(t *testing.T) {
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
	
	// Marshal
	bytes, _ := model.MarshalText()
	
	// Unmarshal
	model2 := &Model{}
	err := model2.UnmarshalText(bytes)
	
	if err != nil {
		t.Fatal(err.Error())
	}
	
	// Compare models
	if model.comment != model2.comment {
		t.Error("different comments:", model.comment, ",", model2.comment)
	}
	
	if len(model.counts) != len(model2.counts) {
	t.Fatal("different counts length:", len(model.counts), ",",
			len(model2.counts))
	}
	
	for i := range model.counts {
		if len(model.counts[i]) != len(model2.counts[i]) {
			t.Errorf("different counts[%d] length: %d , %d",
				i, len(model.counts[i]), len(model2.counts))
		} else {
			for j := range model.counts[i] {
				if model.counts[i][j] != model2.counts[i][j] {
					t.Errorf("different values in counts[%d][%d]: %d , %d",
							i, j, model.counts[i][j], model2.counts[i][j])
				}
			}
		}
	}
}



