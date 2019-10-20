package bedgraph

import (
	"strings"
	"testing"
)

func TestScannerNoHeader(t *testing.T) {
	bedString := "chr1\t10\t20\t3.14\nchr4\t50\t66\t2.7\n"
	scanner := NewScanner(strings.NewReader(bedString))

	exp1 := &BedGraph{"chr1", 10, 20, 3.14}
	exp2 := &BedGraph{"chr4", 50, 66, 2.7}

	if !scanner.Scan() {
		t.Fatal("Scanning failed. Error:", scanner.Err())
	}

	if !compare(scanner.Bed(), exp1) {
		t.Fatal("Bad bed scanned:", scanner.Bed(), "expected:", exp1)
	}

	if !scanner.Scan() {
		t.Fatal("Scanning failed. Error:", scanner.Err())
	}

	if !compare(scanner.Bed(), exp2) {
		t.Fatal("Bad bed scanned:", scanner.Bed(), "expected:", exp2)
	}
}

func TestScannerWithHeader(t *testing.T) {
	bedString := "hjkdsahlkjf\tdsajda\tasdjdakh\nchr1\t10\t20\t3.14\n" +
		"chr4\t50\t66\t2.7\n"
	scanner := NewScanner(strings.NewReader(bedString))

	exp1 := &BedGraph{"chr1", 10, 20, 3.14}
	exp2 := &BedGraph{"chr4", 50, 66, 2.7}

	if !scanner.Scan() {
		t.Fatal("Scanning failed. Error:", scanner.Err())
	}

	if !compare(scanner.Bed(), exp1) {
		t.Fatal("Bad bed scanned:", scanner.Bed(), "expected:", exp1)
	}

	if !scanner.Scan() {
		t.Fatal("Scanning failed. Error:", scanner.Err())
	}

	if !compare(scanner.Bed(), exp2) {
		t.Fatal("Bad bed scanned:", scanner.Bed(), "expected:", exp2)
	}
}

// Compares 2 bed entries.
func compare(b1, b2 *BedGraph) bool {
	return b1.Chr == b1.Chr && b1.Start == b2.Start && b1.End == b2.End &&
		b1.Value == b2.Value
}
