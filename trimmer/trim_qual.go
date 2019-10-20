package main

// Handles trimming of low quality ends of the read.

import (
	"github.com/fluhus/golgi/formats/fastq"
)

// Trims the given fastq according to the given threshold.
// Algorithm is like Trim Galore's - subtract threshold from all quals, then
// sum the quals from the beginning, and trim where the sum is minimal.
func trimQual(fq *fastq.Fastq, offset int, thresholdQual int) {
	// Check input
	if fq == nil {
		panic("unexpected nil fastq")
	}

	// Find trimming coordinates
	sum := 0
	minSum := 0
	minPos := 0
	maxSum := 0
	maxPos := 0

	for i, qual := range fq.Quals {
		sum += int(qual) - int(offset) - thresholdQual
		if sum < minSum {
			minSum = sum
			minPos = i + 1
		}
		if sum >= maxSum {
			maxSum = sum
			maxPos = i + 1
		}
	}

	// Bad quality makes max go past min
	if maxPos <= minPos {
		fq.Sequence = nil
		fq.Quals = nil
	} else {
		fq.Sequence = fq.Sequence[minPos:maxPos]
		fq.Quals = fq.Quals[minPos:maxPos]
	}
}
