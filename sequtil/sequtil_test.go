package seqtools

// Unit test for seqtools.

import (
	"fmt"
	"reflect"
	"testing"
)

// Compares the output of ReverseComplement with the expected output.
func helper_ReverseComplement(t *testing.T, input string, expected string) {
	if ReverseComplementString(input) != expected {
		t.Error(fmt.Sprintf("rc(%s) gave %s, expected %s",
			input, ReverseComplementString(input), expected))
	}
}

func Test_ReverseComplement(t *testing.T) {
	helper_ReverseComplement(t, "A", "T")
	helper_ReverseComplement(t, "AAA", "TTT")
	helper_ReverseComplement(t, "aaa", "ttt")
	helper_ReverseComplement(t, "AACTTGGG", "CCCAAGTT")
	helper_ReverseComplement(t, "TGTGTG", "CACACA")
	helper_ReverseComplement(t, "", "")
}

func TestDNATo2Bit(t *testing.T) {
	tests := []struct {
		input []byte
		want  []byte
	}{
		{[]byte("acgtTGCA"), []byte{0b11100100, 0b00011011}},
		{[]byte("tatat"), []byte{0b00110011, 0b00000011}},
		{[]byte("ccc"), []byte{0b00010101}},
	}
	for _, test := range tests {
		got := make([]byte, len(test.want))
		DNATo2Bit(got, test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("DNATo2Bit(%q)=%v, want %v", test.input, got, test.want)
		}
	}
}
