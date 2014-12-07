package main

import (
	"bioformats/fastq"
	"fmt"
	"os"
	"io"
)

func main() {
	// Check for parsing error
	if argumentError != nil {
		fmt.Fprintln(os.Stderr, "bad arguments:", argumentError.Error())
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
	
	
	fmt.Fprintln(os.Stderr, "Operation successful!")
}

// Prints details about current run.
func printWorkPlan() {
	fmt.Fprintln(os.Stderr, "Biostuff Trimmer - Workplan")
	fmt.Fprintln(os.Stderr, "~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	
	if inputFile == os.Stdin {
		fmt.Fprintln(os.Stderr, "Input: stdin")
	} else {
		fmt.Fprintln(os.Stderr, "Input:", inputFile.Name())
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
	
	fmt.Fprintln(os.Stderr, "\n(Run 'trimmer --help' for usage instructions.)")
}

// Does the read processing, exits on error.
func processReads() {
	var err error
	var fq *fastq.Fastq
	
	for fq, err = fastq.ReadNext(inputReader); err == nil;
			fq, err = fastq.ReadNext(inputReader) {
	
		if qualThreshold != 0 {
			trimQual(fq, phredOffset, qualThreshold)
		}
		
		trimAdapterStart(fq, adapterStart, 5)  // 5 is arbitrary for now
		trimAdapterEnd(fq, adapterStart, 5)    // 5 is arbitrary for now
		
		if len(fq.Sequence) > 0 {
			outputWriter.WriteString(fq.String())
			outputWriter.WriteByte('\n')
		}
	}
	
	if err != io.EOF {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, "output file contents are invalid.")
		os.Exit(1)
	}
	
	flushAndCloseFiles()
	
	// TODO: Add statistics.
}

// Flushes and closes i/o files.
// Run this before exiting the program, if files were opened.
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
trimmer [options] --in <input file> --out <output file>

Options:
	-h
	--help
		Print this help message and ignore all other arguments.
	-i
	--in
		Input fastq file. Give 'stdin' for standard input.
	-o
	--out
		Output fastq file. Give 'stdout' for standard output.
	-q
	--qual-threshold
		Quality trimmming threshold. Give 0 to avoid quality trimming.
		Default: 20.
	-p
	--phred-offset
		Phred quality score offset. Default: 33.
	-as
	--adapter-start
		Adapter to trim at the beginning (5') of the read. Default: none.
	-ae
	--adapter-end
		Adapter to trim at the end (3') of the read. Default: none.`
