package sequtil

import (
	"reflect"
	"testing"
)

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
	for _, test := range tests {
		got := make([]byte, len(test.want))
		DNATo2Bit(got, test.input)
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
		{[]byte{0b11001100, 0b11000000}, []byte("TATAT")},
		{[]byte{0b01010100}, []byte("CCC")},
	}
	for _, test := range tests {
		got := make([]byte, len(test.want))
		DNAFrom2Bit(got, test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("DNAFrom2Bit(%v)=%v, want %v", test.input, got, test.want)
		}
	}
}

func BenchmarkDNATo2Bit(b *testing.B) {
	dna := []byte("acgtacgtacgtacgtacgtacgtacgtacgt")
	twobit := make([]byte, len(dna)/4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DNATo2Bit(twobit, dna)
	}
}

func BenchmarkDNAFrom2Bit(b *testing.B) {
	dna := []byte("acgtacgtacgtacgtacgtacgtacgtacgt")
	twobit := make([]byte, len(dna)/4)
	DNATo2Bit(twobit, dna)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DNAFrom2Bit(dna, twobit)
	}
}
