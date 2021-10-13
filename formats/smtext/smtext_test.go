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
