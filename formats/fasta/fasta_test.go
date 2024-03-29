package fasta

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestNext_simple(t *testing.T) {
	input := ">foo\nAaTtGnNngcCaN"
	want := &Fasta{[]byte("foo"), []byte("AaTtGnNngcCaN")}
	got, err := NewReader(strings.NewReader(input)).Read()
	if err != nil {
		t.Fatalf("Next(%q) failed: %v", input, err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Next(%q)=%v, want %v", input, got, want)
	}
}

func TestNext_noName(t *testing.T) {
	input := "AaTtGngcCaN"
	want := &Fasta{nil, []byte("AaTtGngcCaN")}
	got, err := NewReader(strings.NewReader(input)).Read()
	if err != nil {
		t.Fatalf("Next(%q) failed: %v", input, err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Next(%q)=%v, want %v", input, got, want)
	}
}

func TestNext_multiline(t *testing.T) {
	input := ">foo\nAaTtGngcCaN\nGGgg\n>foo"
	want := &Fasta{[]byte("foo"), []byte("AaTtGngcCaNGGgg")}
	got, err := NewReader(strings.NewReader(input)).Read()
	if err != nil {
		t.Fatalf("Next(%q) failed: %v", input, err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Next(%q)=%v, want %v", input, got, want)
	}
}

func TestNext_multiple(t *testing.T) {
	input := ">foo\nAaTtGngcCaN\nGGgg\n>bar\naaaGcgnnNcTAtgGa\nAA\n\nGagaGNtCc"
	r := NewReader(strings.NewReader(input))
	want := []*Fasta{
		{[]byte("foo"), []byte("AaTtGngcCaNGGgg")},
		{[]byte("bar"), []byte("aaaGcgnnNcTAtgGaAAGagaGNtCc")},
	}
	var got []*Fasta
	var err error
	var fa *Fasta
	for fa, err = r.Read(); err == nil; fa, err = r.Read() {
		got = append(got, fa)
	}
	if err != io.EOF {
		t.Fatalf("ForEach(%q) failed: %v", input, err)
	}
	if len(got) != 2 {
		t.Fatalf("len(ForEach(%q))=%v, want 2", input, len(got))
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("ForEach(%q)=%v, want %v", input, got, want)
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
