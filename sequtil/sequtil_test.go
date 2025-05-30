package sequtil

import (
	"reflect"
	"slices"
	"testing"

	"github.com/fluhus/gostuff/snm"
)

func TestReverseComplement(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"A", "T"},
		{"AAA", "TTT"},
		{"aaa", "ttt"},
		{"AACTTGGG", "CCCAAGTT"},
		{"TGTGTG", "CACACA"},
		{"", ""},
	}
	var got []byte
	for _, test := range tests {
		got = ReverseComplement(got[:0], []byte(test.input))
		if string(got) != test.want {
			t.Errorf("ReverseComplement(%q)=%q, want %q",
				test.input, got, test.want)
		}
	}
}

func TestReverseComplementString(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"A", "T"},
		{"AAA", "TTT"},
		{"aaa", "ttt"},
		{"AACTTGGG", "CCCAAGTT"},
		{"TGTGTG", "CACACA"},
		{"", ""},
	}
	for _, test := range tests {
		got := ReverseComplementString(test.input)
		if got != test.want {
			t.Errorf("ReverseComplementString(%q)=%v, want %v",
				test.input, got, test.want)
		}
	}
}

func TestDNATo2Bit(t *testing.T) {
	tests := []struct {
		input []byte
		want  []byte
	}{
		{[]byte("acgtTGCA"), []byte{0b00011011, 0b11100100}},
		{[]byte("tatat"), []byte{0b11001100, 0b11000000}},
		{[]byte("ccc"), []byte{0b01010100}},
	}
	var got []byte
	for _, test := range tests {
		got = DNATo2Bit(got[:0], test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("DNATo2Bit(%q)=%v, want %v", test.input, got, test.want)
		}
	}
}

func TestDNAFrom2Bit(t *testing.T) {
	tests := []struct {
		input []byte
		want  []byte
	}{
		{[]byte{0b00011011, 0b11100100}, []byte("ACGTTGCA")},
		{[]byte{0b11001100, 0b11000000}, []byte("TATATAAA")},
		{[]byte{0b01010100}, []byte("CCCA")},
	}
	var got []byte
	for _, test := range tests {
		got = DNAFrom2Bit(got[:0], test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("DNAFrom2Bit(%v)=%q, want %q", test.input, got, test.want)
		}
	}
}

func BenchmarkDNATo2Bit(b *testing.B) {
	dna := []byte("acgtacgtacgtacgtacgtacgtacgtacgt")
	var twobit []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twobit = DNATo2Bit(twobit[:0], dna)
	}
}

func BenchmarkDNAFrom2Bit(b *testing.B) {
	dna := []byte("acgtacgtacgtacgtacgtacgtacgtacgt")
	twobit := DNATo2Bit(nil, dna)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dna = DNAFrom2Bit(dna[:0], twobit)
	}
}

func FuzzDNA2Bit(f *testing.F) {
	f.Add([]byte{})
	f.Fuzz(func(t *testing.T, a []byte) {
		aa := slices.Clone(a)
		b := DNAFrom2Bit(nil, a)
		c := DNATo2Bit(nil, b)
		if !slices.Equal(aa, c) {
			t.Errorf("DNATo2Bit(DNAFrom2Bit(...)) before=%q after=%q",
				aa, c)
		}
	})
}

func TestCanonical(t *testing.T) {
	input := "GATTACCA"
	want := []string{"GA", "AT", "AA", "TA", "AC", "CC", "CA"}
	var got []string
	for seq := range CanonicalSubsequences([]byte(input), 2) {
		got = append(got, string(seq))
	}
	if !slices.Equal(got, want) {
		t.Fatalf("CanonicalSubsequences(%q)=%v, want %v", input, got, want)
	}
}

func TestSubsequencesWith(t *testing.T) {
	tests := []struct {
		input string
		chars string
		want  []string
	}{
		{"", "a", nil},
		{"sdfsdfaafd", "a", []string{"aa"}},
		{"sdfsdfaafd", "p", nil},
		{"actattagcatcga", "atcg", []string{"actattagcatcga"}},
		{"actattagcatcga", "atcgATCG", []string{"actattagcatcga"}},
		{"actattagcatcga", "atc", []string{"actatta", "catc", "a"}},
	}
	for _, test := range tests {
		gotBytes := slices.Collect(SubsequencesWith([]byte(test.input), test.chars))
		got := snm.SliceToSlice(gotBytes, func(b []byte) string { return string(b) })
		if !slices.Equal(got, test.want) {
			t.Errorf("SubsequencesWith(%q,%q)=%q, want %q",
				test.input, test.chars, got, test.want)
		}
	}
}
