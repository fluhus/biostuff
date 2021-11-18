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
		tree, err := NewReader(strings.NewReader(test.input)).Read()
		if err != nil {
			t.Fatalf("Read(%q) failed: %v", test.input, err)
		}
		var got []string
		tree.PreOrder(func(n *Node) bool {
			got = append(got, n.Name)
			return true
		})
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
		tree, err := NewReader(strings.NewReader(test.input)).Read()
		if err != nil {
			t.Fatalf("Read(%q) failed: %v", test.input, err)
		}
		var got []string
		tree.PostOrder(func(n *Node) bool {
			got = append(got, n.Name)
			return true
		})
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%q.PostOrder()=%v, want %v", test.input, got, test.want)
		}
	}
}
