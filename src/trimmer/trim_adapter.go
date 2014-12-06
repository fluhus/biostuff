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
	
	trimPos := len(sequence)
	
	// For each overlap
	outerLoop: for si := start; si < len(sequence); si++ {
		matchLength := len(sequence) - si
		remainingMismatches := matchLength / 5  // mismath tolerance
		
		// Compare to adapter starting from this index
		for ai := 0; ai < matchLength; ai++ {
			if sequence[ai+si] != adapter[ai] {
				remainingMismatches--
				if remainingMismatches < 0 {
					continue outerLoop
				}
			}
		}
		
		// If reached here, then adapter was found
		trimPos = si
		break outerLoop
	}
	
	// Trim!
	fq.Sequence = fq.Sequence[:trimPos]
	fq.Quals = fq.Quals[:trimPos]
}



