package regions

import (
	"reflect"
	"testing"
)

func TestIndex(t *testing.T) {
	starts := []int{0, 5, 10}
	ends := []int{3, 15, 12}
	want := &Index{[]interval{
		{0, []int{0}},
		{3, nil},
		{5, []int{1}},
		{10, []int{1, 2}},
		{12, []int{1}},
		{15, nil},
	}}
	got := NewIndex(starts, ends)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("NewIndex(%v, %v)=%v, want %v", starts, ends, got, want)
	}

	tests := []struct {
		pos  int
		want []int
	}{
		{-1, nil},
		{0, []int{0}},
		{1, []int{0}},
		{2, []int{0}},
		{3, nil},
		{4, nil},
		{5, []int{1}},
		{6, []int{1}},
		{7, []int{1}},
		{8, []int{1}},
		{9, []int{1}},
		{10, []int{1, 2}},
		{11, []int{1, 2}},
		{12, []int{1}},
		{13, []int{1}},
		{14, []int{1}},
		{15, nil},
		{16, nil},
	}
	for _, test := range tests {
		if at := got.At(test.pos); !reflect.DeepEqual(at, test.want) {
			t.Errorf("At(%v)=%v, want %v", test.pos, at, test.want)
		}
	}
}
