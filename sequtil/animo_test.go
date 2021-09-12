package sequtil

import (
	"reflect"
	"testing"
)

func TestTranslate(t *testing.T) {
	input := "AGAcatTGGgat"
	want := "RHWD"
	got := make([]byte, 4)
	Translate(got, []byte(input))
	if string(got) != want {
		t.Fatalf("Translate(%q)=%q, want %q", input, got, want)
	}
}

func TestTranslate_badBase(t *testing.T) {
	defer func() { recover() }()
	input := "AGAcatTGGgad"
	Translate(make([]byte, 4), []byte(input))
	t.Fatalf("Translate(%q) succeeded, want panic", input)
}

func TestTranslate_badLength(t *testing.T) {
	defer func() { recover() }()
	input := "AGAcatTGGgatt"
	Translate(make([]byte, 5), []byte(input))
	t.Fatalf("Translate(%q) succeeded, want panic", input)
}

func TestTranslateReadingFrames(t *testing.T) {
	input := "AGAcatTGGgat"
	want := [3][]byte{[]byte("RHWD"), []byte("DIG"), []byte("TLG")}
	if got := TranslateReadingFrames([]byte(input)); !reflect.DeepEqual(got, want) {
		t.Fatalf("TranslateReadingFrames(%q)=%v, want %v", input, got, want)
	}
}

func TestAminoName(t *testing.T) {
	tests := []struct {
		input byte
		want1 string
		want2 string
	}{
		{'H', "His", "Histidine"},
		{'Q', "Gln", "Glutamine"},
		{'h', "His", "Histidine"},
		{'q', "Gln", "Glutamine"},
		{'*', "*", "Stop codon"},
	}

	for _, test := range tests {
		got1, got2 := AminoName(test.input)
		if got1 != test.want1 || got2 != test.want2 {
			t.Fatalf("AminoName('%c')=(%q,%q), want (%q,%q)",
				test.input, got1, got2, test.want1, test.want2)
		}
	}
}

func TestAminoName_bad(t *testing.T) {
	defer func() { recover() }()
	AminoName('U')
	t.Fatalf("AminoName('U') succeeded, want panic")
}
