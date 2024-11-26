package fasta

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestReader_simple(t *testing.T) {
	input := ">foo\nAaTtGnNngcCaN"
	want := []*Fasta{{[]byte("foo"), []byte("AaTtGnNngcCaN")}}
	var got []*Fasta
	for fa, err := range Reader(strings.NewReader(input)) {
		if err != nil {
			t.Fatalf("Reader(%q) failed: %v", input, err)
		}
		got = append(got, fa)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Reader(%q)=%v, want %v", input, got, want)
	}
}

func TestReader_noName(t *testing.T) {
	input := "AaTtGngcCaN"
	want := []*Fasta{{nil, []byte("AaTtGngcCaN")}}
	var got []*Fasta
	for fa, err := range Reader(strings.NewReader(input)) {
		if err != nil {
			t.Fatalf("Reader(%q) failed: %v", input, err)
		}
		got = append(got, fa)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Reader(%q)=%v, want %v", input, got, want)
	}
}

func TestReader_multiline(t *testing.T) {
	input := ">foo\nAaTtGngcCaN\nGGgg\n>foo"
	want := []*Fasta{
		{[]byte("foo"), []byte("AaTtGngcCaNGGgg")},
		{[]byte("foo"), nil},
	}
	var got []*Fasta
	for fa, err := range Reader(strings.NewReader(input)) {
		if err != nil {
			t.Fatalf("Reader(%q) failed: %v", input, err)
		}
		got = append(got, fa)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Reader(%q)=%v, want %v", input, got, want)
	}
}

func TestReader_multiple(t *testing.T) {
	input := ">foo\nAaTtGngcCaN\nGGgg\n>bar\naaaGcgnnNcTAtgGa\nAA\n\nGagaGNtCc"
	want := []*Fasta{
		{[]byte("foo"), []byte("AaTtGngcCaNGGgg")},
		{[]byte("bar"), []byte("aaaGcgnnNcTAtgGaAAGagaGNtCc")},
	}
	var got []*Fasta
	for fa, err := range Reader(strings.NewReader(input)) {
		if err != nil {
			t.Fatalf("Reader(%q) failed: %v", input, err)
		}
		got = append(got, fa)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Reader(%q)=%v, want %v", input, got, want)
	}
}

func TestMarshalText(t *testing.T) {
	tests := []struct {
		input *Fasta
		want  string
	}{
		{&Fasta{[]byte("Hello"), []byte("ATGGCC")}, ">Hello\nATGGCC\n"},
		{&Fasta{[]byte("Bye"), nil}, ">Bye\n"},
		{&Fasta{[]byte("Howdy"), bytes.Repeat([]byte("AATTGGCC"), 25)},
			fmt.Sprintf(">Howdy\n%s\n%s\n%s\n",
				strings.Repeat("AATTGGCC", 10),
				strings.Repeat("AATTGGCC", 10),
				strings.Repeat("AATTGGCC", 5),
			)},
	}
	for _, test := range tests {
		if got, err := test.input.MarshalText(); err != nil ||
			string(got) != test.want {
			t.Errorf("%v.Text()=%v,%v, want %v", test.input, got, err, test.want)
		}
	}
}

func BenchmarkMarshalText(b *testing.B) {
	fa := &Fasta{Name: []byte("bla bla bla bla bla bla"),
		Sequence: bytes.Repeat([]byte("a"), 1000)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fa.MarshalText()
	}
}
