package bed

import (
	"reflect"
	"testing"
)

func TestIndex_simple(t *testing.T) {
	b := NewIndexBuilder()
	b.Add("chr1", 0, 10, "1")
	idx := b.Build()
	t.Log(idx.str())

	tests := []struct {
		chr  string
		pos  int
		want []string
	}{
		{"chr1", 0, []string{"1"}},
		{"chr1", 5, []string{"1"}},
		{"chr1", 9, []string{"1"}},
		{"chr1", 10, []string{}},
	}

	for _, test := range tests {
		if n := idx.Names(test.chr, test.pos); !reflect.DeepEqual(n, names(test.want)) {
			t.Errorf("Names(%v,%v)=%v, want %v",
				test.chr, test.pos, n, names(test.want))
		}
	}
}

func TestIndex_complex(t *testing.T) {
	b := NewIndexBuilder()
	b.Add("chr1", 0, 10, "1")
	b.Add("chr1", 5, 15, "2")
	b.Add("chr1", 10, 20, "4")
	b.Add("chr1", 21, 30, "8")
	b.Add("chr1", 25, 35, "8")
	b.Add("chr2", 100, 200, "a")
	b.Add("chr2", 150, 250, "b")
	b.Add("chr2", 120, 170, "a")
	b.Add("chr2", 180, 270, "60")
	idx := b.Build()

	tests := []struct {
		chr  string
		pos  int
		want []string
	}{
		{"chr1", 0, []string{"1"}},
		{"chr1", 4, []string{"1"}},
		{"chr1", 5, []string{"1", "2"}},
		{"chr1", 10, []string{"2", "4"}},
		{"chr1", 20, []string{}},
		{"chr1", 21, []string{"8"}},
		{"chr1", 25, []string{"8"}},

		{"chr2", 0, []string{}},
		{"chr2", 99, []string{}},
		{"chr2", 100, []string{"a"}},
		{"chr2", 150, []string{"a", "b"}},
		{"chr2", 170, []string{"a", "b"}},
		{"chr2", 180, []string{"a", "b", "60"}},
		{"chr2", 200, []string{"b", "60"}},

		{"chr3", 0, []string{}},
	}

	for _, test := range tests {
		if n := idx.Names(test.chr, test.pos); !reflect.DeepEqual(n, names(test.want)) {
			t.Errorf("Names(%v,%v)=%v, want %v",
				test.chr, test.pos, n, names(test.want))
		}
	}
}

func names(n []string) map[string]struct{} {
	result := map[string]struct{}{}
	for _, name := range n {
		result[name] = struct{}{}
	}
	return result
}
