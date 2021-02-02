package fastq

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestTrim(t *testing.T) {
	input := "\r\n\r\namitamit\n\ramit\n\n\r\n"
	got := string(trimNewLines([]byte(input)))
	want := "amitamit\n\ramit"
	if got != want {
		t.Errorf("trimNewLines(%q)=%v, want %v", input, got, want)
	}
}

func TestNext_simple(t *testing.T) {
	input := "@hello\nAAATTTGG\n+\n!!!@@@!!"
	want := &Fastq{[]byte("hello"), []byte("AAATTTGG"), []byte("!!!@@@!!")}
	got, err := Next(bufio.NewReader(strings.NewReader(input)))
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
		if got, err := Next(bufio.NewReader(strings.NewReader(input))); err == nil {
			t.Errorf("Next(%q)=%v, want fail", input, got)
		}
	}
}

func TestForEach(t *testing.T) {
	input := "@a\nAA\n+\n!!\n@c\nCCC\n+\nKKK"
	want := []*Fastq{
		{[]byte("a"), []byte("AA"), []byte("!!")},
		{[]byte("c"), []byte("CCC"), []byte("KKK")},
	}
	var got []*Fastq
	err := ForEach(strings.NewReader(input), func(fq *Fastq) error {
		got = append(got, fq)
		return nil
	})
	if err != nil {
		t.Fatalf("ForEach(%q) failed: %v", input, err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("ForEach(%q)=%v, want %v", input, got, want)
	}
}
