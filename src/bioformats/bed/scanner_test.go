package bed

import (
	"testing"
	"strings"
)

func TestScannerNoHeader(t *testing.T) {
	bedString := "chr1\t10\t20\nchr4\t50\t66\n"
	scanner := NewScanner( strings.NewReader(bedString) )
	
	exp1 := &Bed{"chr1", 10, 20}
	exp2 := &Bed{"chr4", 50, 66}
	
	if !scanner.Scan() {
		t.Fatal("Scanning failed. Error:", scanner.Err())
	}
	
	if !compare(scanner.Current(), exp1) {
		t.Fatal("Bad bed scanned:", scanner.Current(), "expected:", exp1)
	}
	
	if !scanner.Scan() {
		t.Fatal("Scanning failed. Error:", scanner.Err())
	}
	
	if !compare(scanner.Current(), exp2) {
		t.Fatal("Bad bed scanned:", scanner.Current(), "expected:", exp2)
	}
}

func TestScannerWithHeader(t *testing.T) {
	bedString := "hjkdsahlkjf\tdsajda\tasdjdakh\nchr1\t10\t20\nchr4\t50\t66\n"
	scanner := NewScanner( strings.NewReader(bedString) )
	
	exp1 := &Bed{"chr1", 10, 20}
	exp2 := &Bed{"chr4", 50, 66}
	
	if !scanner.Scan() {
		t.Fatal("Scanning failed. Error:", scanner.Err())
	}
	
	if !compare(scanner.Current(), exp1) {
		t.Fatal("Bad bed scanned:", scanner.Current(), "expected:", exp1)
	}
	
	if !scanner.Scan() {
		t.Fatal("Scanning failed. Error:", scanner.Err())
	}
	
	if !compare(scanner.Current(), exp2) {
		t.Fatal("Bad bed scanned:", scanner.Current(), "expected:", exp2)
	}
}

// Compares 2 bed entries.
func compare(b1, b2 *Bed) bool {
	return b1.Chr == b1.Chr && b1.Start == b2.Start && b1.End == b2.End
}

