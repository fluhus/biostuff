package stats

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStats(t *testing.T) {
	assert := assert.New(t)

	v1 := []float64{1, 3, 5}
	v2 := []float64{10, 7, 4}
	v3 := []float64{1, 1, 1, 1}
	v4 := []float64{1, 1, 1, 1, 0, 0}

	assert.Equal(Sum(v1), 9.0)
	assert.Equal(Mean(v1), 3.0)
	assert.Equal(Var(v1), 8.0/3.0)
	assert.Equal(Corr(v1, v2), -1.0)
	assert.Equal(Ent(v3), 2.0)
	assert.Equal(Ent(v4), 2.0)
}
