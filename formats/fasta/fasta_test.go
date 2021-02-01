package fasta

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestNext_simple(t *testing.T) {
	input := ">foo\nAaTtGnNngcCaN"
	want := &Fasta{[]byte("foo"), []byte("AaTtGnNngcCaN")}
	got, err := Next(bufio.NewReader(strings.NewReader(input)))
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
	got, err := Next(bufio.NewReader(strings.NewReader(input)))
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
	got, err := Next(bufio.NewReader(strings.NewReader(input)))
	if err != nil {
		t.Fatalf("Next(%q) failed: %v", input, err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Next(%q)=%v, want %v", input, got, want)
	}
}

func TestForEach_simple(t *testing.T) {
	input := ">foo\nAaTtGngcCaN\nGGgg\n>bar\naaaGcgnnNcTAtgGa\nAA\n\nGagaGNtCc"
	want := []*Fasta{
		{[]byte("foo"), []byte("AaTtGngcCaNGGgg")},
		{[]byte("bar"), []byte("aaaGcgnnNcTAtgGaAAGagaGNtCc")},
	}
	var got []*Fasta
	err := ForEach(strings.NewReader(input), func(fa *Fasta) error {
		got = append(got, fa)
		return nil
	})
	if err != nil {
		t.Fatalf("ForEach(%q) failed: %v", input, err)
	}
	if len(got) != 2 {
		t.Fatalf("len(ForEach(%q))=%v, want 2", input, len(got))
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("ForEach(%q)=%v, want %v", input, got, want)
	}
}
