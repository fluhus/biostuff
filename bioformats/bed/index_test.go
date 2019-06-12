package bed

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndex_simple(t *testing.T) {
	assert := assert.New(t)
	builder := NewIndexBuilder()
	builder.Add("chr1", 0, 10, "1")
	idx := builder.Build()
	t.Log(idx.str())
	assert.Equal(names("1"), idx.Names("chr1", 0))
	assert.Equal(names("1"), idx.Names("chr1", 5))
	assert.Equal(names("1"), idx.Names("chr1", 9))
	assert.Equal(names(), idx.Names("chr1", 10))
}

func TestIndex_complex(t *testing.T) {
	assert := assert.New(t)

	builder := NewIndexBuilder()
	builder.Add("chr1", 0, 10, "1")
	builder.Add("chr1", 5, 15, "2")
	builder.Add("chr1", 10, 20, "4")
	builder.Add("chr1", 21, 30, "8")
	builder.Add("chr1", 25, 35, "8")
	builder.Add("chr2", 100, 200, "a")
	builder.Add("chr2", 150, 250, "b")
	builder.Add("chr2", 120, 170, "a")
	builder.Add("chr2", 180, 270, "60")

	idx := builder.Build()
	assert.Equal(names("1"), idx.Names("chr1", 0))
	assert.Equal(names("1"), idx.Names("chr1", 4))
	assert.Equal(names("1", "2"), idx.Names("chr1", 5))
	assert.Equal(names("2", "4"), idx.Names("chr1", 10))
	assert.Equal(names(), idx.Names("chr1", 20))
	assert.Equal(names("8"), idx.Names("chr1", 21))
	assert.Equal(names("8"), idx.Names("chr1", 25))

	assert.Equal(names(), idx.Names("chr2", 0))
	assert.Equal(names(), idx.Names("chr2", 99))
	assert.Equal(names("a"), idx.Names("chr2", 100))
	assert.Equal(names("a", "b"), idx.Names("chr2", 150))
	assert.Equal(names("a", "b"), idx.Names("chr2", 170))
	assert.Equal(names("a", "b", "60"), idx.Names("chr2", 180))
	assert.Equal(names("b", "60"), idx.Names("chr2", 200))

	assert.Equal(names(), idx.Names("chr3", 0))
}

func names(n ...string) map[string]struct{} {
	result := map[string]struct{}{}
	for _, name := range n {
		result[name] = struct{}{}
	}
	return result
}
