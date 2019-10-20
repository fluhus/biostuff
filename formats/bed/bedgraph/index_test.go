package bedgraph

import (
	"reflect"
	"testing"
)

func TestIndex_simple(t *testing.T) {
	b := NewIndexBuilder()
	b.Add("chr1", 0, 10, 1)
	idx := b.Build()
	t.Log(idx.str())

	tests := []struct {
		chr  string
		pos  int
		want float64
	}{
		{"chr1", 0, 1.0},
		{"chr1", 9, 1.0},
		{"chr1", 10, 0.0},
	}

	for _, test := range tests {
		if got := idx.Value(test.chr, test.pos); got != test.want {
			t.Errorf("Value(%v,%v)=%v, want %v",
				test.chr, test.pos, got, test.want)
		}
	}
}

func TestIndex_complex(t *testing.T) {
	builder := NewIndexBuilder()
	builder.Add("chr1", 0, 10, 1)
	builder.Add("chr1", 5, 15, 2)
	builder.Add("chr1", 10, 20, 4)
	builder.Add("chr1", 21, 30, 8)
	builder.Add("chr1", 25, 35, 8)
	builder.Add("chr2", 180, 270, 60)
	idx := builder.Build()

	tests := []struct {
		chr  string
		pos  int
		want float64
	}{
		{"chr1", 0, 1.0},
		{"chr1", 4, 1.0},
		{"chr1", 5, 3.0},
		{"chr1", 10, 6.0},
		{"chr1", 20, 0.0},
		{"chr1", 21, 8.0},
		{"chr1", 25, 16.0},
		{"chr1", 30, 8.0},
		{"chr1", 35, 0.0},

		{"chr2", 179, 0.0},
		{"chr2", 180, 60.0},
		{"chr2", 200, 60.0},
		{"chr2", 269, 60.0},
		{"chr2", 270, 0.0},

		{"chr3", 0, 0.0},
	}

	for _, test := range tests {
		if got := idx.Value(test.chr, test.pos); got != test.want {
			t.Errorf("Value(%v,%v)=%v, want %v",
				test.chr, test.pos, got, test.want)
		}
	}
}

func TestIndex_valueRange(t *testing.T) {
	builder := NewIndexBuilder()
	builder.Add("chr1", 5, 15, 1)
	builder.Add("chr1", 8, 20, 2)
	idx := builder.Build()

	tests := []struct {
		chr  string
		from int
		to   int
		want []float64
	}{
		{"chr1", 4, 7, []float64{0, 1, 1}},
		{"chr1", 0, 23, []float64{0, 0, 0, 0, 0, 1, 1, 1, 3, 3, 3, 3,
			3, 3, 3, 2, 2, 2, 2, 2, 0, 0, 0}},
	}

	for _, test := range tests {
		if got := idx.ValueRange(
			test.chr, test.from, test.to); !reflect.DeepEqual(got, test.want) {
			t.Errorf("Value(%v,%v,%v)=%v, want %v",
				test.chr, test.from, test.to, got, test.want)
		}
	}
}
