package seqtools

// Unit test for seqtools.

import (
	"testing"
	"fmt"
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
