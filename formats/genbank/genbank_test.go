package genbank

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/fluhus/gostuff/iterx"
)

func TestSplitLine(t *testing.T) {
	tests := []struct {
		input, want1, want2 string
	}{
		{"", "", ""},
		{"a", "a", ""},
		{"aaa", "aaa", ""},
		{"aaa    ", "aaa", ""},
		{"aaa    bbbb", "aaa", "bbbb"},
		{"aaa    bbbb c  df", "aaa", "bbbb c  df"},
		{"  ORGANISM", "  ORGANISM", ""},
		{"  ORGANISM    ", "  ORGANISM", ""},
		{"  ORGANISM    ha hi ho", "  ORGANISM", "ha hi ho"},
	}
	for _, test := range tests {
		got := splitLine(test.input, nil)
		got1, got2 := got[1], got[2]
		if got1 != test.want1 || got2 != test.want2 {
			t.Errorf("splitLine(%q)=%q,%q want %q,%q",
				test.input, got1, got2, test.want1, test.want2)
		}
	}
}

func TestReader(t *testing.T) {
	tests := []struct {
		input string
		want  *GenBank
	}{
		{input1, want1},
		{input2, want2},
		{input3, want3},
		{input4, want4},
	}

	for i, test := range tests {
		input, want := test.input, test.want
		gots, err := iterx.CollectErr(Reader(strings.NewReader(input)))
		if err != nil {
			t.Fatalf("test #%d: failed to parse: %v", i+1, err)
		}
		if len(gots) != 1 {
			t.Fatalf("expected one result, got %v", len(gots))
		}
		got := gots[0]

		if got.Locus != want.Locus {
			t.Errorf("Reader(...).Locus=%q, want %q", got.Locus, want.Locus)
		}
		if got.Definition != want.Definition {
			t.Errorf("Reader(...).Definition=%q, want %q", got.Definition, want.Definition)
		}
		if err := compareSlices(got.Accessions, want.Accessions); err != nil {
			t.Errorf("Reader(...).Accessions: %v", err)
		}
		if got.Version != want.Version {
			t.Errorf("Reader(...).Version=%q, want %q", got.Version, want.Version)
		}
		if got.Keywords != want.Keywords {
			t.Errorf("Reader(...).Keywords=%q, want %q", got.Keywords, want.Keywords)
		}
		if got.Source != want.Source {
			t.Errorf("Reader(...).Source=%q, want %q", got.Source, want.Source)
		}
		if got.Organism != want.Organism {
			t.Errorf("Reader(...).Organism=%q, want %q", got.Organism, want.Organism)
		}
		if got.OrganismTax != want.OrganismTax {
			t.Errorf("Reader(...).OrganismTax=%q, want %q", got.OrganismTax, want.OrganismTax)
		}
		if err := compareSlices(got.References, want.References); err != nil {
			t.Errorf("Reader(...).References: %v", err)
		}
		if err := compareSlices(got.Features, want.Features); err != nil {
			t.Errorf("Reader(...).Features: %v", err)
		}
		if got.Origin != want.Origin {
			t.Errorf("Reader(...).Origin=%q, want %q", got.Origin, want.Origin)
		}
	}
}

func compareSlices[T any](a, b []T) error {
	if len(a) != len(b) {
		return fmt.Errorf("mismatching lengths: got %d, want %d", len(a), len(b))
	}
	for i := range a {
		if !reflect.DeepEqual(a[i], b[i]) {
			return fmt.Errorf("#%d mismatch: got %v, want %v", i+1, a[i], b[i])
		}
	}
	return nil
}
