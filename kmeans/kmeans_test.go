package kmeans

import (
	"testing"
	"math/rand"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestKmeans(t *testing.T) {
	assert := assert.New(t)
	rand.Seed(time.Now().UnixNano())
	
	m := [][]float64 {
		[]float64 { 0.1, 0.0 },
		[]float64 { 0.9, 1.0 },
		[]float64 { -0.1, 0.0 },
		[]float64 { 0.0, -0.1 },
		[]float64 { 1.1, 1.0 },
		[]float64 { 1.0, 1.1 },
		[]float64 { 1.0, 0.9 },
		[]float64 { 0.0, 0.1 },
	}
	
	tags, means := Kmeans(m, 2)
	
	if tags[0] == 0 {
		assert.Equal(tags, []int {0, 1, 0, 0, 1, 1, 1, 0})
		assert.Equal(means, [][]float64 { []float64 {0,0}, []float64 {1,1} })
	} else {
		assert.Equal(tags, []int {1, 0, 1, 1, 0, 0, 0, 1})
		assert.Equal(means, [][]float64 { []float64 {1,1}, []float64 {0,0} })
	}
}

