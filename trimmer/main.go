package main

import (
	"fmt"
	"io"
	"os"
	"runtime/pprof"

	"github.com/fluhus/golgi/formats/fastq"
)

func main() {
	// Check for argument parsing error
	parseArguments()
	if argumentError != nil {
		fmt.Fprintln(os.Stderr, "Bad arguments:", argumentError.Error())
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	// Print help message if needed
	if printHelp {
		fmt.Fprintln(os.Stderr, usage)
		return
	}

	// Get to work!
	printWorkPlan()
	processReads()
	flushAndCloseFiles()
	printStatistics()

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

	if profileFile != nil {
		fmt.Fprintln(os.Stderr, "Profiling info:", profileFile.Name())
	}

	fmt.Fprintln(os.Stderr, "Actions:")
	if qualThreshold != 0 {
		fmt.Fprintf(os.Stderr, "\tTrim low quality ends;"+
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

	fmt.Fprintln(os.Stderr, "\tOmit reads shorter than:",
		minReadLength)

	fmt.Fprintln(os.Stderr)
}

// Does the read processing, exits on error.
func processReads() {
	// Read fastq
	var err error
	var fq *fastq.Fastq

	for fq, err = fastq.Read(inputReader); err == nil; fq, err =
		fastq.Read(inputReader) {
		readCount++
		nucleotideCount += len(fq.Sequence)

		if qualThreshold != 0 {
			lenBefore := len(fq.Sequence)
			trimQual(fq, phredOffset, qualThreshold)
			lenAfter := len(fq.Sequence)

			qualCount += lenBefore - lenAfter
		}

		if len(adapterStart) > 0 {
			lenBefore := len(fq.Sequence)
			trimAdapterStart(fq, adapterStart, 10) // 10 is arbitrary for now
			lenAfter := len(fq.Sequence)

			adapterStartCount[lenBefore-lenAfter]++
		}

		if len(adapterEnd) > 0 {
			lenBefore := len(fq.Sequence)
			trimAdapterEnd(fq, adapterEnd, 10) // 10 is arbitrary for now
			lenAfter := len(fq.Sequence)

			adapterEndCount[lenBefore-lenAfter]++
		}

		// Print if long enough
		// TODO: add as command line option
		if len(fq.Sequence) >= minReadLength {
			outputWriter.WriteString(fq.String())
			outputWriter.WriteByte('\n')
		} else {
			shortCount++
		}
	}

	if err != io.EOF {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, "Output file contents are invalid.")
		os.Exit(1)
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

	if profileFile != nil {
		pprof.StopCPUProfile()
		profileFile.Close()
	}
}
