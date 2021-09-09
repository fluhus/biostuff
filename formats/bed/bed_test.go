package bed

import (
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	input := []string{"chr1", "10", "20", "Hello", "150", "+", "11", "13",
		"50,100,150", "2", "40,60", "100,200"}
	want := &BED{"chr1", 10, 20, "Hello", 150, "+", 11, 13, [3]byte{50, 100, 150},
		2, []int{40, 60}, []int{100, 200}}

	// Full line.
	got, n, err := parseLine(input)
	if err != nil {
		t.Fatalf("parseLine(%v) failed: %v", input, err)
	}
	if n != 12 {
		t.Fatalf("parseLine(%v) n=%v want %v", input, n, 12)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("parseLine(%v)=%v want %v", input, got, want)
	}

	// Partial line.
	input = input[:6]
	want = &BED{Chrom: "chr1", ChromStart: 10, ChromEnd: 20, Name: "Hello",
		Score: 150, Strand: "+"}
	got, n, err = parseLine(input)
	if err != nil {
		t.Fatalf("parseLine(%v) failed: %v", input, err)
	}
	if n != 6 {
		t.Fatalf("parseLine(%v) n=%v want %v", input, n, 12)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("parseLine(%v)=%v want %v", input, got, want)
	}
}

func TestParseLine_bad(t *testing.T) {
	input := []string{"chr1", "10", "20", "Hello", "150", "+", "11", "13",
		"50,100,150", "2", "40,60", "100,200"}
	cp := make([]string, len(input))
	copy(cp, input)

	// Check good input.
	if _, _, err := parseLine(cp); err != nil {
		t.Fatalf("parseLine(%v) failed: %v", cp, err)
	}

	// Make bad modifications.
	if got, _, err := parseLine(cp[:2]); err == nil {
		t.Fatalf("parseLine(%v)=%v want error", cp[:2], got)
	}
	cp[5] = "t" // Bad strand
	if got, _, err := parseLine(cp); err == nil {
		t.Fatalf("parseLine(%v)=%v want error", cp, got)
	}
	copy(cp, input)
	cp[8] = "100" // Bad colors
	if got, _, err := parseLine(cp); err == nil {
		t.Fatalf("parseLine(%v)=%v want error", cp, got)
	}
	copy(cp, input)
	cp[8] += "0" // Bad colors (overflow)
	if got, _, err := parseLine(cp); err == nil {
		t.Fatalf("parseLine(%v)=%v want error", cp, got)
	}
	copy(cp, input)
	cp[10] += ",200" // Bad block starts
	if got, _, err := parseLine(cp); err == nil {
		t.Fatalf("parseLine(%v)=%v want error", cp, got)
	}
}
