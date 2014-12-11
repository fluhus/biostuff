package main

import (
	"bioformats/fastq"
	"fmt"
	"os"
	"io"
)

func main() {
	// Check for argument parsing error
	parseArguments()
	if argumentError != nil {
		fmt.Fprintln(os.Stderr, "Bad arguments:", argumentError.Error())
		fmt.Fprintln(os.Stderr, "Run 'trimmer --help' for usage instructions.")
		os.Exit(1)
	}
	
	// Print help message if needed
	if printHelp {
		fmt.Fprintln(os.Stderr, usage)
		return
	}
	
	// Print work plan
	printWorkPlan()
	
	// Start working!
	processReads()
	
	fmt.Fprintln(os.Stderr, "Trimmer: operation successful!")
}

// Prints details about current run.
func printWorkPlan() {
	fmt.Fprintln(os.Stderr, "Biostuff Trimmer - Workplan")
	fmt.Fprintln(os.Stderr, "~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	
	if inputFile == os.Stdin {
		fmt.Fprintln(os.Stderr, "Input:  stdin")
	} else {
		fmt.Fprintln(os.Stderr, "Input: ", inputFile.Name())
	}
	
	if outputFile == os.Stdout {
		fmt.Fprintln(os.Stderr, "Output: stdout")
	} else {
		fmt.Fprintln(os.Stderr, "Output:", outputFile.Name())
	}
	
	fmt.Fprintln(os.Stderr, "Actions:")
	if qualThreshold != 0 {
		fmt.Fprintf(os.Stderr, "\tTrim low quality ends;" +
				" threshold=%d, offset=%d\n",
				qualThreshold, phredOffset)
	}
	
	if len(adapterStart) > 0 {
		fmt.Fprintln(os.Stderr, "\tTrim adapter from start:",
				string(adapterStart))
	}
	
	if len(adapterEnd) > 0 {
		fmt.Fprintln(os.Stderr, "\tTrim adapter from end:",
				string(adapterEnd))
	}

	fmt.Fprintln(os.Stderr)
}

// Does the read processing, exits on error.
func processReads() {
	// Collect statistics
	readCount := 0
	qualCount := 0
	adapterStartCount := make([]int, len(adapterStart) + 1)
	adapterEndCount   := make([]int, len(adapterEnd) + 1)

	// Read fastq
	var err error
	var fq *fastq.Fastq
	
	for fq, err = fastq.ReadNext(inputReader); err == nil;
			fq, err = fastq.ReadNext(inputReader) {
		readCount++
	
		if qualThreshold != 0 {
			lenBefore := len(fq.Sequence)
			trimQual(fq, phredOffset, qualThreshold)
			lenAfter := len(fq.Sequence)

			qualCount += lenBefore - lenAfter
		}
		
		if len(adapterStart) > 0 {
			lenBefore := len(fq.Sequence)
			trimAdapterStart(fq, adapterStart, 10)  // 10 is arbitrary for now
			lenAfter := len(fq.Sequence)

			adapterStartCount[lenBefore - lenAfter]++
		}

		if len(adapterEnd) > 0 {
			lenBefore := len(fq.Sequence)
			trimAdapterEnd(fq, adapterEnd, 10)    // 10 is arbitrary for now
			lenAfter := len(fq.Sequence)

			adapterEndCount[lenBefore - lenAfter]++
		}
		
		if len(fq.Sequence) > 0 {
			outputWriter.WriteString(fq.String())
			outputWriter.WriteByte('\n')
		}
	}
	
	if err != io.EOF {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, "Output file contents are invalid.")
		os.Exit(1)
	}
	
	flushAndCloseFiles()
	printStatistics(readCount, qualCount, adapterStartCount, adapterEndCount)
}

// Nicely prints the run statistics.
func printStatistics(readCount int, qualCount int, adapterStartCount []int, adapterEndCount []int) {
	// All reads count
	fmt.Fprintln(os.Stderr, "Number of reads processed:", readCount)

	// Quality trimming count
	if qualThreshold != 0 {
		fmt.Fprintln(os.Stderr, "Number of low quality nucleotide trimmed:", qualCount)
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

// Flushes and closes i/o files.
// Run this before exiting the program, if i/o was done.
func flushAndCloseFiles() {
	if inputFile != os.Stdin {
		inputFile.Close()
	}
	
	outputWriter.Flush()
	if outputFile != os.Stdout {
		outputFile.Close()
	}
}

const usage =
`Biostuff Trimmer
~~~~~~~~~~~~~~~~

Trims low quality ends and adapter contamination from reads.

Written by Amit Lavon.

Usage:
trimmer [options] -in <input file> -out <output file>

Options:
	-h
	-help
		Print this help message and ignore all other arguments.
	-i
	-in
		Input fastq file. Give 'stdin' for standard input.
	-o
	-out
		Output fastq file. Give 'stdout' for standard output.
	-q
	-qual-threshold
		Quality trimmming threshold. Give 0 to avoid quality trimming.
		Default: 20.
	-p
	-phred-offset
		Phred quality score offset. Default: 33.
	-as
	-adapter-start
		Adapter to trim at the beginning (5') of the read. Default: none.
	-ae
	-adapter-end
		Adapter to trim at the end (3') of the read. Default: none.
`
