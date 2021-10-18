package align

import (
	"reflect"
	"testing"
)

func TestGlobal(t *testing.T) {
	m := SubstitutionMatrix{
		{'a', 'a'}: 1,
		{'b', 'b'}: 1,
		{'a', 'b'}: 0,
		{'a', Gap}: 0,
		{'b', Gap}: 0,
		{Gap, Gap}: -1,
	}.Symmetrical()
	tests := []struct {
		a, b      string
		want      []Step
		wantScore float64
	}{
		{"aba", "aba", []Step{Match, Match, Match}, 3},
		{"aba", "aa", []Step{Match, Deletion, Match}, 1},
		{"aa", "aba", []Step{Match, Insertion, Match}, 1},
	}
	for _, test := range tests {
		got, gotScore := Global([]byte(test.a), []byte(test.b), m)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Global(%q,%q)=%v, want %v", test.a, test.b, got, test.want)
		}
		if gotScore != test.wantScore {
			t.Errorf("Global(%q,%q) score=%v, want %v",
				test.a, test.b, gotScore, test.wantScore)
		}
	}
}

func TestLevenshtein(t *testing.T) {
	tests := []struct {
		a, b string
		want float64
	}{
		{"kitten", "sitten", -1},
		{"sittin", "sitting", -1},
		{"brexit", "exit", -2},
		{"super", "uber", -2},
	}
	for _, test := range tests {
		if steps, got := Global([]byte(test.a), []byte(test.b),
			Levenshtein); got != test.want {
			t.Errorf("Global(%q,%q)=%v score=%v, want %v",
				test.a, test.b, steps, got, test.want)
		}
	}
}
