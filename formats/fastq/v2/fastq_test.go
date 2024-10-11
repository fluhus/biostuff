package fastq

import (
	"reflect"
	"strings"
	"testing"
)

func TestReader_simple(t *testing.T) {
	input := "@hello\nAAATTTGG\n+\n!!!@@@!!"
	want := []*Fastq{{[]byte("hello"), []byte("AAATTTGG"), []byte("!!!@@@!!")}}
	var got []*Fastq
	for fq, err := range Reader(strings.NewReader(input)) {
		if err != nil {
			t.Fatalf("Reader(%q) failed: %v", input, err)
		}
		got = append(got, fq)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Reader(%q)=%v, want %v", input, got, want)
	}
}

func TestReader_bad(t *testing.T) {
	inputs := []string{
		"@hello\nAAA\n+",
		"@hello\nAAA\n+!!",
		"hello\nAAA\n+\n!!!",
		"@hello\nAAA\n-\n!!!",
	}
	for _, input := range inputs {
		for got, err := range Reader(strings.NewReader(input)) {
			if err == nil {
				t.Errorf("Reader(%q)=%v, want fail", input, got)
			}
			break
		}
	}
}

func TestReader_many(t *testing.T) {
	input := "@a\nAA\n+\n!!\n@c\nCCC\n+\nKKK"
	want := []*Fastq{
		{[]byte("a"), []byte("AA"), []byte("!!")},
		{[]byte("c"), []byte("CCC"), []byte("KKK")},
	}
	var got []*Fastq
	for fq, err := range Reader(strings.NewReader(input)) {
		if err != nil {
			t.Fatalf("Reader(%q) failed: %v", input, err)
		}
		got = append(got, fq)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Reader(%q)=%v, want %v", input, got, want)
	}
}

func TestMarshalTextText(t *testing.T) {
	input := &Fastq{[]byte("Hello"), []byte("AGAGAG"), []byte("!@##@!")}
	want := "@Hello\nAGAGAG\n+\n!@##@!\n"
	if got, err := input.MarshalText(); err != nil || string(got) != want {
		t.Fatalf("%v.Text()=%v,%v want %v", input, got, err, want)
	}
}
