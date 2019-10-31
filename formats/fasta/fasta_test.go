package fasta

import (
	"bufio"
	"strings"
	"testing"
)

func TestRead_simple(t *testing.T) {
	input := ">foo\nAaTtGnNngcCaN"
	fa, err := Read(bufio.NewReader(strings.NewReader(input)))
	if err != nil {
		t.Fatalf("ReadEntry(%q) failed: %v", input, err)
	}
	checkFasta(t, fa, "foo", "AaTtGnNngcCaN")
	//assert.Equal(2, len(fa.nEnds))
}

func TestRead_noName(t *testing.T) {
	input := "AaTtGngcCaN"
	fa, err := Read(bufio.NewReader(strings.NewReader(input)))
	if err != nil {
		t.Fatalf("Read(%q) failed: %v", input, err)
	}
	checkFasta(t, fa, "", input)
}

func TestReady_multiline(t *testing.T) {
	input := ">foo\nAaTtGngcCaN\nGGgg\n>foo"

	fa, err := Read(bufio.NewReader(strings.NewReader(input)))
	if err != nil {
		t.Fatalf("Read(%q) failed: %v", input, err)
	}
	checkFasta(t, fa, "foo", "AaTtGngcCaNGGgg")
}

func TestRead_badChars(t *testing.T) {
	input := ">foo\naaaaaaKaaaaa"
	_, err := Read(bufio.NewReader(strings.NewReader(input)))
	if err == nil {
		t.Fatalf("Read(%q) succeeded, want fail", input)
	}
}

func TestReadAll_simple(t *testing.T) {
	input := ">foo\nAaTtGngcCaN\nGGgg\n>bar\naaaGcgnnNcTAtgGa\nAA\n\nGagaGNtCc"

	fas, err := ReadAll(strings.NewReader(input))
	if err != nil {
		t.Fatalf("ReadAll(%q) failed: %v", input, err)
	}
	if len(fas) != 2 {
		t.Fatalf("len(ReadAll(%q))=%v, want 2", input, len(fas))
	}
	checkFasta(t, fas[0], "foo", "AaTtGngcCaNGGgg")
	checkFasta(t, fas[1], "bar", "aaaGcgnnNcTAtgGaAAGagaGNtCc")
}

func checkFasta(t *testing.T, fa *Fasta, name string, sequence string) {
	if name != fa.Name() {
		t.Fatalf("Name()=%q, want %q", fa.Name(), name)
	}

	usequence := strings.ToUpper(sequence)
	fasequence := ""
	for i := 0; i < fa.Len(); i++ {
		fasequence += string([]byte{fa.At(i)})
	}
	if usequence != fasequence {
		t.Fatalf("At(...)=%q, want %q", fasequence, usequence)
	}
	for i := range usequence {
		for l := 0; l < len(usequence)-i; l++ {
			usub := usequence[i : i+l]
			fasub := string(fa.Subsequence(i, l))
			if usub != fasub {
				t.Fatalf("Subsequence(%v,%v)=%q, want %q",
					i, l, fasub, usub)
			}
		}
	}
}
