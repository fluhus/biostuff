package vectors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestL1(t *testing.T) {
	assert := assert.New(t)

	v1 := []float64{0, 0, 0}
	v2 := []float64{1, 0, 0}
	v3 := []float64{0, 1, 1}
	v4 := []float64{1, 1, 1}

	assert.Equal(L1(v1, v1), 0.0)
	assert.Equal(L1(v1, v2), 1.0)
	assert.Equal(L1(v1, v3), 2.0)
	assert.Equal(L1(v1, v4), 3.0)
	assert.Equal(L1(v2, v4), 2.0)
	assert.Equal(L1(v3, v4), 1.0)
}

func TestL2(t *testing.T) {
	assert := assert.New(t)

	v1 := []float64{0, 0}
	v2 := []float64{0, 1}
	v3 := []float64{4, 4}

	assert.Equal(L2(v1, v1), 0.0)
	assert.Equal(L2(v1, v2), 1.0)
	assert.Equal(L2(v3, v2), 5.0)
}
