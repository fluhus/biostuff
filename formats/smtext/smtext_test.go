package smtext

import (
	"reflect"
	"strings"
	"testing"

	"github.com/fluhus/biostuff/align"
)

func TestReadNCBI(t *testing.T) {
	tests := []struct {
		input string
		want  align.SubstitutionMatrix
	}{
		{
			" A B \nB 1 2\n A 3 4\n",
			align.SubstitutionMatrix{
				{'B', 'A'}: 1,
				{'B', 'B'}: 2,
				{'A', 'A'}: 3,
				{'A', 'B'}: 4,
			},
		},
		{
			"# comment\n A B \nB 1 2\n A 3 4\n#more commentsss",
			align.SubstitutionMatrix{
				{'B', 'A'}: 1,
				{'B', 'B'}: 2,
				{'A', 'A'}: 3,
				{'A', 'B'}: 4,
			},
		},
	}
	for _, test := range tests {
		got, err := ReadNCBI(strings.NewReader(test.input))
		if err != nil {
			t.Fatalf("ReadNCBI(%q) failed: %v", test.input, err)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("ReadNCBI(%q)=%v, want %v",
				test.input, got, test.want)
		}
	}
}

func TestReadNCBI_bad(t *testing.T) {
	tests := []string{
		" A B \nB 1 2 3\n A 3 4\n",
		" A B \nB 1 2f\n A 3 4\n",
	}
	for _, test := range tests {
		got, err := ReadNCBI(strings.NewReader(test))
		if err == nil {
			t.Fatalf("ReadNCBI(%q)=%v, want error", test, got)
		}
	}
}

func TestGoString(t *testing.T) {
	input := align.SubstitutionMatrix{
		{'a', 'b'}:             1,
		{'a', 'c'}:             16,
		{'a', align.Gap}:       49,
		{'b', 'b'}:             4,
		{'b', 'c'}:             25,
		{'b', align.Gap}:       64,
		{align.Gap, 'b'}:       9,
		{align.Gap, 'c'}:       36,
		{align.Gap, align.Gap}: 81,
	}
	want := "SubstitutionMatrix{\n" +
		"{'a','b'}:1,\n{'a','c'}:16,\n{'a',Gap}:49,\n" +
		"{'b','b'}:4,\n{'b','c'}:25,\n{'b',Gap}:64,\n" +
		"{Gap,'b'}:9,\n{Gap,'c'}:36,\n{Gap,Gap}:81,\n" +
		"}\n"
	if got := GoString(input); string(got) != want {
		t.Fatalf("%v.GoString=%q, want %q", input, got, want)
	}
}
