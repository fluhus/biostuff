package newick

import (
	"reflect"
	"strings"
	"testing"
)

func TestPreOrder(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{"A;", []string{"A"}},
		{"(B)A;", []string{"A", "B"}},
		{"((A,B)C,(D,E)F)R;", []string{"R", "C", "A", "B", "F", "D", "E"}},
	}
	for _, test := range tests {
		tree, err := newReader(strings.NewReader(test.input)).read()
		if err != nil {
			t.Fatalf("Read(%q) failed: %v", test.input, err)
		}
		var got []string
		for n := range tree.PreOrder() {
			got = append(got, n.Name)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%q.PreOrder()=%v, want %v", test.input, got, test.want)
		}
	}
}

func TestPostOrder(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{"A;", []string{"A"}},
		{"(B)A;", []string{"B", "A"}},
		{"((A,B)C,(D,E)F)R;", []string{"A", "B", "C", "D", "E", "F", "R"}},
	}
	for _, test := range tests {
		tree, err := newReader(strings.NewReader(test.input)).read()
		if err != nil {
			t.Fatalf("Read(%q) failed: %v", test.input, err)
		}
		var got []string
		for n := range tree.PostOrder() {
			got = append(got, n.Name)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%q.PostOrder()=%v, want %v", test.input, got, test.want)
		}
	}
}
