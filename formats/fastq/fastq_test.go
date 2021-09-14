package fastq

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestNext_simple(t *testing.T) {
	input := "@hello\nAAATTTGG\n+\n!!!@@@!!"
	want := &Fastq{[]byte("hello"), []byte("AAATTTGG"), []byte("!!!@@@!!")}
	got, err := NewReader(strings.NewReader(input)).Read()
	if err != nil {
		t.Fatalf("Next(%q) failed: %v", input, err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Next(%q)=%v, want %v", input, got, want)
	}
}

func TestNext_bad(t *testing.T) {
	inputs := []string{
		"@hello\nAAA\n+",
		"@hello\nAAA\n+!!",
		"hello\nAAA\n+\n!!!",
		"@hello\nAAA\n-\n!!!",
	}
	for _, input := range inputs {
		if got, err := NewReader(strings.NewReader(input)).Read(); err == nil {
			t.Errorf("Next(%q)=%v, want fail", input, got)
		}
	}
}

func TestNext_many(t *testing.T) {
	input := "@a\nAA\n+\n!!\n@c\nCCC\n+\nKKK"
	want := []*Fastq{
		{[]byte("a"), []byte("AA"), []byte("!!")},
		{[]byte("c"), []byte("CCC"), []byte("KKK")},
	}
	var got []*Fastq
	r := NewReader(strings.NewReader(input))
	var fq *Fastq
	var err error
	for fq, err = r.Read(); err == nil; fq, err = r.Read() {
		got = append(got, fq)
	}
	if err != nil && err != io.EOF {
		t.Fatalf("ForEach(%q) failed: %v", input, err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("ForEach(%q)=%v, want %v", input, got, want)
	}
}

func TestText(t *testing.T) {
	input := &Fastq{[]byte("Hello"), []byte("AGAGAG"), []byte("!@##@!")}
	want := "@Hello\nAGAGAG\n+\n!@##@!\n"
	if got := string(input.Text()); got != want {
		t.Fatalf("%v.Text()=%v want %v", input, got, want)
	}
}
