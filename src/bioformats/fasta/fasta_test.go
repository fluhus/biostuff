package fasta

import (
	"testing"
)

func Test_Basic(t *testing.T) {
	f := newFastaEntry()
	f.name = "amit"

	sequence := "AAAATTTT"
	for i := range sequence {
		t.Logf("Appending '%c'", sequence[i])
		f.append(sequence[i])
	}

	if f.Name() != "amit" {
		t.Errorf("Bad name: '%s', expected 'amit'.")
	}

	if f.Length() != len(sequence) {
		t.Fatalf("Bad fasta length: %d, expected %d.", f.length, len(sequence))
	}

	if f.sequence[0] != 0 {
		t.Errorf("Bad sequence value at 0 (AAAA): %d, expected 0.", f.sequence[0])
	}

	if f.sequence[1] != 255 {
		t.Errorf("Bad sequence value at 1 (TTTT): %d, expected 255.", f.sequence[0])
	}
}

func Test_Basic2(t *testing.T) {
	f := newFastaEntry()
	f.name = "lavon"

	sequence := "AACCNGGTTCANN"
	for i := range sequence {
		t.Logf("Appending '%c'", sequence[i])
		f.append(sequence[i])
	}

	if f.Name() != "lavon" {
		t.Errorf("Bad name: '%s', expected 'lavon'.")
	}

	if f.Length() != len(sequence) {
		t.Fatalf("Bad fasta length: %d, expected %d.", f.length, len(sequence))
	}

	for i := range sequence {
		if f.At(i) != sequence[i] {
			t.Errorf("Bad nucleotide at index %d: %c, expected %c.", i, f.At(i), sequence[i])
		}
	}
}

func Test_Subsequence(t *testing.T) {
	f := newFastaEntry()
	f.name = "yoink"

	sequence := "AACCNGGTTCANN"
	for i := range sequence {
		t.Logf("Appending '%c'", sequence[i])
		f.append(sequence[i])
	}

	if f.Length() != len(sequence) {
		t.Fatalf("Bad fasta length: %d, expected %d.", f.length, len(sequence))
	}

	for i := range sequence {
		for j := i; j <= len(sequence); j++ {
			subs := string(f.Subsequence(i, j-i))
			if subs != sequence[i:j] {
				t.Errorf("Bad subsequence: '%s', expected '%s'.",
						subs, sequence[i:j])
			}
		}
	}
}




