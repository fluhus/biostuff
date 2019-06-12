package fasta

import (
	"bufio"
	"strings"
	"testing"
)

func TestReadEntry_simple(t *testing.T) {
	input := ">foo\nAaTtGnNngcCaN"
	fa, err := ReadEntry(bufio.NewReader(strings.NewReader(input)))
	if err != nil {
		t.Fatalf("ReadEntry(%q) failed: %v", input, err)
	}
	checkFasta(t, fa, "foo", "AaTtGnNngcCaN")
	//assert.Equal(2, len(fa.nEnds))
}

func TestReadEntry_noName(t *testing.T) {
	input := "AaTtGngcCaN"
	fa, err := ReadEntry(bufio.NewReader(strings.NewReader(input)))
	if err != nil {
		t.Fatalf("ReadEntry(%q) failed: %v", input, err)
	}
	checkFasta(t, fa, "", input)
}

func TestReadEntry_multiline(t *testing.T) {
	input := ">foo\nAaTtGngcCaN\nGGgg\n>foo"

	fa, err := ReadEntry(bufio.NewReader(strings.NewReader(input)))
	if err != nil {
		t.Fatalf("ReadEntry(%q) failed: %v", input, err)
	}
	checkFasta(t, fa, "foo", "AaTtGngcCaNGGgg")
}

func TestReadEntry_badChars(t *testing.T) {
	input := ">foo\naaaaaaKaaaaa"
	_, err := ReadEntry(bufio.NewReader(strings.NewReader(input)))
	if err == nil {
		t.Fatalf("ReadEntry(%q) succeeded, want fail", input)
	}
}

func TestReadFasta_simple(t *testing.T) {
	input := ">foo\nAaTtGngcCaN\nGGgg\n>bar\naaaGcgnnNcTAtgGa\nAA\n\nGagaGNtCc"

	fas, err := ReadFasta(strings.NewReader(input))
	if err != nil {
		t.Fatalf("ReadFasta(%q) failed: %v", input, err)
	}
	if len(fas) != 2 {
		t.Fatalf("len(ReadFasta(%q))=%v, want 2", input, len(fas))
	}
	checkFasta(t, fas[0], "foo", "AaTtGngcCaNGGgg")
	checkFasta(t, fas[1], "bar", "aaaGcgnnNcTAtgGaAAGagaGNtCc")
}

func checkFasta(t *testing.T, fa *Entry, name string, sequence string) {
	if name != fa.Name() {
		t.Fatalf("Name()=%q, want %q", fa.Name(), name)
	}

	usequence := strings.ToUpper(sequence)
	fasequence := ""
	for i := 0; i < fa.Length(); i++ {
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
