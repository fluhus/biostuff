package bed

import (
	"reflect"
	"slices"
	"strings"
	"testing"
)

func TestParseLine(t *testing.T) {
	input := []string{"chr1", "10", "20", "Hello", "150", "+", "11", "13",
		"50,100,150", "2", "40,60", "100,200"}
	want := &BED{12, "chr1", 10, 20, "Hello", 150, "+", 11, 13, [3]byte{50, 100, 150},
		2, []int{40, 60}, []int{100, 200}}

	// Full line.
	got, err := parseLine(input)
	if err != nil {
		t.Fatalf("parseLine(%v) failed: %v", input, err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("parseLine(%v)=%v want %v", input, got, want)
	}

	// Partial line.
	input = input[:6]
	want = &BED{N: 6, Chrom: "chr1", ChromStart: 10, ChromEnd: 20, Name: "Hello",
		Score: 150, Strand: "+"}
	got, err = parseLine(input)
	if err != nil {
		t.Fatalf("parseLine(%v) failed: %v", input, err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("parseLine(%v)=%v want %v", input, got, want)
	}
}

func TestParseLine_bad(t *testing.T) {
	input := []string{"chr1", "10", "20", "Hello", "150", "+", "11", "13",
		"50,100,150", "2", "40,60", "100,200"}
	cp := slices.Clone(input)

	// Check good input.
	if _, err := parseLine(cp); err != nil {
		t.Fatalf("parseLine(%v) failed: %v", cp, err)
	}

	// Make bad modifications.
	if got, err := parseLine(cp[:2]); err == nil {
		t.Fatalf("parseLine(%v)=%v want error", cp[:2], got)
	}
	cp[5] = "t" // Bad strand
	if got, err := parseLine(cp); err == nil {
		t.Fatalf("parseLine(%v)=%v want error", cp, got)
	}
	cp = slices.Clone(input)
	cp[8] = "100" // Bad colors
	if got, err := parseLine(cp); err == nil {
		t.Fatalf("parseLine(%v)=%v want error", cp, got)
	}
	cp = slices.Clone(input)
	cp[8] += "0" // Bad colors (overflow)
	if got, err := parseLine(cp); err == nil {
		t.Fatalf("parseLine(%v)=%v want error", cp, got)
	}
	cp = slices.Clone(input)
	cp[10] += ",200" // Bad block starts
	if got, err := parseLine(cp); err == nil {
		t.Fatalf("parseLine(%v)=%v want error", cp, got)
	}
}

func TestReader(t *testing.T) {
	input := "chr1\t10\t20\tHello\t150\t+\t11\t13\t50,100,150\t2\t40,60\t100,200\n"
	want := []*BED{{12, "chr1", 10, 20, "Hello", 150, "+", 11, 13, [3]byte{50, 100, 150},
		2, []int{40, 60}, []int{100, 200}}}
	var got []*BED

	for bed, err := range Reader(strings.NewReader(input)) {
		if err != nil {
			t.Fatalf("Next() failed: %v", err)
		}
		got = append(got, bed)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Next()=%v want %v", got, want)
	}
}

func TestMarshalText(t *testing.T) {
	want := "chr1\t10\t20\tHello\t150\t+\t11\t13\t50,100,150\t2\t40,60\t100,200\n"
	input := &BED{12, "chr1", 10, 20, "Hello", 150, "+", 11, 13, [3]byte{50, 100, 150},
		2, []int{40, 60}, []int{100, 200}}
	got, err := input.MarshalText()
	if err != nil {
		t.Fatalf("%v.MarshalText() failed: %v", input, err)
	}
	if string(got) != want {
		t.Fatalf("%v.MarshalText()=%q, want %q", input, got, want)
	}
	input.N = 6
	want = want[:22] + "\n"
	got, err = input.MarshalText()
	if err != nil {
		t.Fatalf("%v.MarshalText() failed: %v", input, err)
	}
	if string(got) != want {
		t.Fatalf("%v.MarshalText()=%q, want %q", input, got, want)
	}
}
