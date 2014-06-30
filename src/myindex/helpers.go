package myindex

// Helper functions and stuff.

import (
	"os"
	"fmt"
	"math"
)

// Enables assertions.
const assert = false

// Reports if assertions are enabled.
func init() {
	if assert {
		fmt.Fprintln(os.Stderr, "*** package myindex: assertions enabled ***")
	}
}

// Maps nucleotides to integers.
var nt2int = map[byte]int {
	'a' : 0,
	'c' : 1,
	'g' : 2,
	't' : 3,
	'A' : 0,
	'C' : 1,
	'G' : 2,
	'T' : 3,
}

// Returns the index of the given k-mer, out of all k-long
// sequences. Returns -1 for characters not in "acgtACGT".
func kmerIndex(kmer []byte) int {
	result := 0
	
	for _,nuc := range kmer {
		nucInt, ok := nt2int[nuc]
		
		if !ok {
			// Unknown nucleotide
			return -1
		}
		
		result = result*4 + nucInt
	}
	
	return result
}

// Returns the number of possible k-long k-mers.
func numOfKmers(k int) int {
	return int( math.Pow(4, float64(k)) )
}

