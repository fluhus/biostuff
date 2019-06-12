package bedgraph

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIndex_simple(t *testing.T) {
	assert := assert.New(t)
	builder := NewIndexBuilder()
	builder.Add("chr1", 0, 10, 1)
	idx := builder.Build()
	t.Log(idx.str())
	assert.Equal(1.0, idx.Value("chr1", 0))
	assert.Equal(1.0, idx.Value("chr1", 9))
	assert.Equal(0.0, idx.Value("chr1", 10))
}

func TestIndex_complex(t *testing.T) {
	assert := assert.New(t)
	
	builder := NewIndexBuilder()
	builder.Add("chr1", 0, 10, 1)
	builder.Add("chr1", 5, 15, 2)
	builder.Add("chr1", 10, 20, 4)
	builder.Add("chr1", 21, 30, 8)
	builder.Add("chr1", 25, 35, 8)
	builder.Add("chr2", 180, 270, 60)
	
	idx := builder.Build()
	assert.Equal(1.0, idx.Value("chr1", 0))
	assert.Equal(1.0, idx.Value("chr1", 4))
	assert.Equal(3.0, idx.Value("chr1", 5))
	assert.Equal(6.0, idx.Value("chr1", 10))
	assert.Equal(0.0, idx.Value("chr1", 20))
	assert.Equal(8.0, idx.Value("chr1", 21))
	assert.Equal(16.0, idx.Value("chr1", 25))
	assert.Equal(8.0, idx.Value("chr1", 30))
	assert.Equal(0.0, idx.Value("chr1", 35))
	
	assert.Equal(0.0, idx.Value("chr2", 179))
	assert.Equal(60.0, idx.Value("chr2", 180))
	assert.Equal(60.0, idx.Value("chr2", 200))
	assert.Equal(60.0, idx.Value("chr2", 269))
	assert.Equal(0.0, idx.Value("chr2", 270))
	
	assert.Equal(0.0, idx.Value("chr3", 0))
}

func TestIndex_range(t *testing.T) {
	assert := assert.New(t)
	
	builder := NewIndexBuilder()
	builder.Add("chr1", 5, 15, 1)
	builder.Add("chr1", 8, 20, 2)
	idx := builder.Build()

	assert.Equal([]float64{0, 1, 1}, idx.ValueRange("chr1", 4, 7))
	assert.Equal([]float64{0,0,0,0,0,1,1,1,3,3,3,3,3,3,3,2,2,2,2,2,0,0,0},
			idx.ValueRange("chr1", 0, 23))
}



