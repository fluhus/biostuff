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
	got, err := NewReader(strings.NewReader(input)).Next()
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
	got, err := NewReader(strings.NewReader(input)).Next()
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
	got, err := NewReader(strings.NewReader(input)).Next()
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
	for fa, err = r.Next(); err == nil; fa, err = r.Next() {
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

func TestText(t *testing.T) {
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
		if got := string(test.input.Text()); got != test.want {
			t.Errorf("%v.Text()=%v, want %v", test.input, got, test.want)
		}
	}
}
