package sequtil

import (
	"reflect"
	"testing"

	"golang.org/x/exp/slices"
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
