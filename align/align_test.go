package align

import (
	"reflect"
	"testing"
)

func TestDecideOnStep(t *testing.T) {
	tests := []struct {
		mch, del, ins float64
		want          block
	}{
		{1, 0, 0, block{step: Match, score: 1}},
		{1, 2, 0, block{step: Deletion, score: 2}},
		{1, 2, 3, block{step: Insertion, score: 3}},
	}
	for _, test := range tests {
		if got := decideOnStep(test.mch, test.del, test.ins); got != test.want {
			t.Errorf("decideOnStep(%v,%v,%v)=%v, want %v",
				test.mch, test.del, test.ins, got, test.want)
		}
	}
}

func TestTraceAlignmentSteps(t *testing.T) {
	tests := []struct {
		blocks    []block
		bn        int
		want      []Step
		wantScore float64
	}{
		{[]block{{}, {}, {}, {score: 5, step: Match}},
			2, []Step{Match}, 5},
		{[]block{{}, {}, {}, {}, {step: Match}, {}, {}, {}, {score: 10, step: Match}},
			3, []Step{Match, Match}, 10},
		{[]block{{}, {step: Insertion}, {}, {}, {},
			{step: Match}, {}, {}, {score: 8, step: Deletion}},
			3, []Step{Insertion, Match, Deletion}, 8},
		{[]block{{}, {}, {},
			{step: Deletion}, {}, {},
			{step: Deletion}, {step: Insertion}, {score: 7, step: Insertion}},
			3, []Step{Deletion, Deletion, Insertion, Insertion}, 7},
	}
	for _, test := range tests {
		got, gotScore := traceAlignmentSteps(test.blocks, test.bn)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("traceAlignmentSteps(...)=%v, want %v", got, test.want)
		}
		if gotScore != test.wantScore {
			t.Errorf("traceAlignmentSteps(...)= score=%v, want %v", gotScore, test.wantScore)
		}
	}
}

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

func TestSymmetrical(t *testing.T) {
	input := SubstitutionMatrix{
		{'a', 'a'}: 5,
		{'a', 'b'}: 3,
		{'d', 'c'}: 2,
		{'e', 'f'}: 6,
		{'f', 'e'}: 6,
	}
	want := SubstitutionMatrix{
		{'a', 'a'}: 5,
		{'a', 'b'}: 3,
		{'b', 'a'}: 3,
		{'c', 'd'}: 2,
		{'d', 'c'}: 2,
		{'e', 'f'}: 6,
		{'f', 'e'}: 6,
	}
	if got := input.Symmetrical(); !reflect.DeepEqual(got, want) {
		t.Fatalf("%v.Symmetrical()=%v, want %v", input, got, want)
	}
}

func TestSymmetrical_bad(t *testing.T) {
	defer func() { recover() }()
	input := SubstitutionMatrix{
		{'a', 'a'}: 5,
		{'a', 'b'}: 3,
		{'b', 'a'}: 2,
	}
	got := input.Symmetrical()
	t.Fatalf("%v.Symmetrical()=%v, want panic", input, got)
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

func TestGoString(t *testing.T) {
	input := SubstitutionMatrix{
		{'a', 'b'}: 1,
		{'a', 'c'}: 16,
		{'a', Gap}: 49,
		{'b', 'b'}: 4,
		{'b', 'c'}: 25,
		{'b', Gap}: 64,
		{Gap, 'b'}: 9,
		{Gap, 'c'}: 36,
		{Gap, Gap}: 81,
	}
	want := "SubstitutionMatrix{\n" +
		"{'a','b'}:1,\n{'a','c'}:16,\n{'a',Gap}:49,\n" +
		"{'b','b'}:4,\n{'b','c'}:25,\n{'b',Gap}:64,\n" +
		"{Gap,'b'}:9,\n{Gap,'c'}:36,\n{Gap,Gap}:81,\n" +
		"}\n"
	if got := input.GoString(); got != want {
		t.Fatalf("%v.GoString=%q, want %q", input, got, want)
	}
}
