package main

// Handles statistics printing.

import (
	"fmt"
	"os"
)

var (
	readCount           int  // How many reads in input file
	nucleotideCount     int  // How many nucleotides in input file
	qualCount           int  // How many nucleotides were dropped for low quality
	adapterStartCount []int  // Histogram of trimmed adapter lengths
	adapterEndCount   []int  // Histogram of trimmed adapter lengths
)

// Nicely prints the run statistics.
func printStatistics() {
	// All reads count
	fmt.Fprintln(os.Stderr, "Number of reads processed:", readCount)
	fmt.Fprintln(os.Stderr, "Number of nucleotides in reads:", nucleotideCount)

	// Quality trimming count
	if qualThreshold != 0 {
		fmt.Fprintln(os.Stderr, "Number of low quality nucleotides trimmed:",
				qualCount)
	}

	// Shorten adapter count slices (up to last non-zero)
	if len(adapterStartCount) > 0 {
		trimAt := 0
		for i,v := range adapterStartCount {
			if v > 0 {
				trimAt = i
			}
		}
		adapterStartCount = adapterStartCount[:trimAt+1]
	}

	if len(adapterEndCount) > 0 {
		trimAt := 0
		for i,v := range adapterEndCount {
			if v > 0 {
				trimAt = i
			}
		}
		adapterEndCount = adapterEndCount[:trimAt+1]
	}

	// Adapter trimming counts
	if len(adapterStart) > 0 {
		fmt.Fprintln(os.Stderr, "\n5' adapters trimmed:\nlength\tcount")
		for i := 1; i < len(adapterStartCount); i++ {
			fmt.Fprintf(os.Stderr, "%d\t%d\n", i, adapterStartCount[i])
		}
	}

	if len(adapterEnd) > 0 {
		fmt.Fprintln(os.Stderr, "\n3' adapters trimmed:\nlength\tcount")
		for i := 1; i < len(adapterEndCount); i++ {
			fmt.Fprintf(os.Stderr, "%d\t%d\n", i, adapterEndCount[i])
		}
	}
}
