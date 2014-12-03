package main

// Handles trimming of low quality ends of the read.

import (
	"bioformats/fastq"
)

// Trims the given fastq according to the given threshold.
// Algorithm is like Trim Galore's - subtract threshold from all quals, then
// sum the quals from the beginning, and trim where the sum is minimal.
func trimQual(fq *fastq.Fastq, offset fastq.PhredOffset, thresholdQual int) {
	// Check input
	if fq == nil {
		panic("unexpected nil fastq")
	}
	
	if len(fq.Quals) == 0 {
		return
	}

	// Trim from start
	sum := 0
	pos := 0
	minSum := 0
	
	for i,qual := range fq.Quals {
		sum += int(qual) - int(offset) - thresholdQual
		if sum < minSum {
			minSum = sum
			pos = i+1
		}
	}
	
	fq.Sequence = fq.Sequence[pos:]
	fq.Quals = fq.Quals[pos:]
	
	if len(fq.Quals) == 0 {
		return
	}
	
	// Trim from finish
	sum = 0
	pos = len(fq.Quals)
	minSum = 0
	
	finishThisFunction
}
