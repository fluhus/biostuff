package main

// Handles trimming of adapters.

import (
	"bioformats/fastq"
	"bytes"
)

// Trims the given fastq's end (3') according to the given adapter.
// It takes the longest overlap that includes the adapter's start and
// sequence's end, and has at most n/5 mismatches, where n is the length of
// the overlap. Assumes no indels in the adapter.
func trimAdapterEnd(fq *fastq.Fastq, adapter []byte) {
	// Check input
	if fq == nil {
		panic("unexpected nil fastq")
	}
	
	// Turn sequences to uppercase for case-insensitivity
	adapter   = bytes.ToUpper(adapter)
	sequence := bytes.ToUpper(fq.Sequence)
	
	// Match search start position
	start := 0
	if len(sequence) > len(adapter) {
		start = len(sequence) - len(adapter)
	}
	
	// For each overlap
	outerLoop: for si := start; si < len(sequence); si++ {
		matchLength := len(sequence) - si
		numberOfMismatches := 0
		
		// Compare to adapter starting from this index
		innerLoop: for ai := 0; ai < matchLength; ai++ {
			continueHere
		}
	}
}



