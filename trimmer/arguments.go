package main

// Parses command line arguments.

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime/pprof"
)

// Parsed arguments.
var (
	inputFile     *os.File      // input file
	outputFile    *os.File      // output file
	inputReader   *bufio.Reader // read from here
	outputWriter  *bufio.Writer // write to here
	profileFile   *os.File      // profiling information is written to here
	adapterStart  []byte        // adapter to trim from start
	adapterEnd    []byte        // adapter to trim from end
	phredOffset   int           // phred quality offset
	qualThreshold int           // quality trimming threshold
	minReadLength int           // Shorter reads are omitted
	printHelp     bool          // should I print help message?
	argumentError error         // not nil if an error occured
)

// Parses command line arguments.
func parseArguments() {
	// Get argument values
	flags := flag.NewFlagSet("trimmer", flag.ContinueOnError)
	flags.SetOutput(ioutil.Discard)

	input := flags.String("in", "", "")
	flags.StringVar(input, "i", "", "")

	output := flags.String("out", "", "")
	flags.StringVar(output, "o", "", "")

	profile := flags.String("profile", "", "")

	var adapterStartString string
	flags.StringVar(&adapterStartString, "adapter-start", "", "")
	flags.StringVar(&adapterStartString, "as", "", "")

	var adapterEndString string
	flags.StringVar(&adapterEndString, "adapter-end", "", "")
	flags.StringVar(&adapterEndString, "ae", "", "")

	flags.IntVar(&phredOffset, "phred-offset", 33, "")
	flags.IntVar(&phredOffset, "p", 33, "")

	flags.IntVar(&qualThreshold, "qual-threshold", 20, "")
	flags.IntVar(&qualThreshold, "q", 20, "")

	flags.IntVar(&minReadLength, "min-length", 20, "")
	flags.IntVar(&minReadLength, "l", 20, "")
	if minReadLength < 1 {
		argumentError = fmt.Errorf("Bad min length threshold: %d",
			minReadLength)
		return
	}

	flags.BoolVar(&printHelp, "help", false, "")
	flags.BoolVar(&printHelp, "h", false, "")

	argumentError = flags.Parse(os.Args[1:])
	if argumentError != nil {
		return
	}

	// Check if help (no need to do further parsing)
	if printHelp {
		return
	}

	// Check if any action was selected
	if qualThreshold == 0 && len(adapterStart) > 0 && len(adapterEnd) > 0 {
		argumentError = errors.New("No trimming action selected.")
		return
	}

	// Convert adapter strings to byte slices
	adapterStart = []byte(adapterStartString)
	adapterEnd = []byte(adapterEndString)

	// Initialize statistic slices
	adapterStartCount = make([]int, len(adapterStart)+1)
	adapterEndCount = make([]int, len(adapterEnd)+1)

	// Open i/o files
	if *input == "" {
		argumentError = errors.New("No input file given.")
		return
	}

	if *output == "" {
		argumentError = errors.New("No output file given.")
		return
	}

	if *input == "stdin" {
		inputFile = os.Stdin
	} else {
		inputFile, argumentError = os.Open(*input)
		if argumentError != nil {
			return
		}
	}

	if *output == "stdout" {
		outputFile = os.Stdout
	} else {
		outputFile, argumentError = os.Create(*output)
		if argumentError != nil {
			return
		}
	}

	if *profile != "" {
		profileFile, argumentError = os.Create(*profile)
		if argumentError != nil {
			return
		}
		argumentError = pprof.StartCPUProfile(profileFile)
		if argumentError != nil {
			return
		}
	}

	// Create buffered i/o
	inputReader = bufio.NewReader(inputFile)
	outputWriter = bufio.NewWriter(outputFile)
}

// Printed if arguments are bad.
const usage = `Trims low quality ends and adapter contamination from reads.

Written by Amit Lavon (amitlavon1@gmail.com).

Usage:
trimmer [options] -in <input file> -out <output file>

Options:
	-h
	-help
		Print this help message and ignore all other arguments.

	-i <path>
	-in <path>
		Input fastq file. Give 'stdin' for standard input.

	-o <path>
	-out <path>
		Output fastq file. Give 'stdout' for standard output.

	-q <integer>
	-qual-threshold <integer>
		Quality trimmming threshold. Give 0 to avoid quality trimming.
		Default: 20.

	-p <integer>
	-phred-offset <integer>
		Phred quality score offset. Default: 33.

	-as <string>
	-adapter-start <string>
		Adapter to trim at the beginning (5') of the read. Default: none.

	-ae <string>
	-adapter-end <string>
		Adapter to trim at the end (3') of the read. Default: none.

	-l <integer>
	-min-length <integer>
		Reads that become shorter than the given value are ommitted.
		Default: 20.

	-profile <path>
		Print profiling information to the given file. Default: none.
		(For development only.)
`
