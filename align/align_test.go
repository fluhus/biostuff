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
