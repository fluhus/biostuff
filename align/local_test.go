package align

import (
	"reflect"
	"testing"
)

func TestLocal(t *testing.T) {
	m := SubstitutionMatrix{
		{'a', 'a'}: 1,
		{'b', 'b'}: 1,
		{'c', 'c'}: 1,
		{'a', 'b'}: -1,
		{'a', 'c'}: -1,
		{'b', 'c'}: -1,
		{'a', Gap}: -1,
		{'b', Gap}: -1,
		{'c', Gap}: -1,
		{Gap, Gap}: 0,
	}.Symmetrical()
	tests := []struct {
		a, b           string
		want           []Step
		wantAI, wantBI int
		wantScore      float64
	}{
		{"aaaa", "aaaa", []Step{Match, Match, Match, Match}, 0, 0, 4},
		{"aaaa", "bbbb", nil, -1, -1, 0},
		{"bbaaaacc", "ccaaaabb", []Step{Match, Match, Match, Match}, 2, 2, 4},
	}
	for _, test := range tests {
		got, ai, bi, gotScore := Local([]byte(test.a), []byte(test.b), m)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Local(%q,%q)=%v, want %v", test.a, test.b, got, test.want)
		}
		if ai != test.wantAI {
			t.Errorf("Local(%q,%q) ai=%v, want %v",
				test.a, test.b, ai, test.wantAI)
		}
		if bi != test.wantBI {
			t.Errorf("Local(%q,%q) bi=%v, want %v",
				test.a, test.b, bi, test.wantBI)
		}
		if gotScore != test.wantScore {
			t.Errorf("Local(%q,%q) score=%v, want %v",
				test.a, test.b, gotScore, test.wantScore)
		}
	}
}
