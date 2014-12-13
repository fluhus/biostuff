package main

// Parses command line arguments.

import (
	"bufio"
	"flag"
	"errors"
	"io/ioutil"
	"os"
)

// Parsed arguments.
var (
	inputFile     *os.File      // input file
	outputFile    *os.File      // output file
	inputReader   *bufio.Reader // read from here
	outputWriter  *bufio.Writer // write to here
	adapterStart  []byte        // adapter to trim from start
	adapterEnd    []byte        // adapter to trim from end
	phredOffset   int           // phred quality offset
	qualThreshold int           // quality trimming threshold
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
	
	flags.BoolVar(&printHelp, "help", false, "")
	flags.BoolVar(&printHelp, "h", false, "")
	
	argumentError = flags.Parse(os.Args[1:])
	if argumentError != nil { return }
	
	// Check if help (no need to do further parsing)
	if printHelp { return }
	
	// Check if any action was selected
	if qualThreshold == 0 && len(adapterStart) > 0 && len(adapterEnd) > 0 {
		argumentError = errors.New("No trimming action selected.")
		return
	}
	
	// Convert adapter strings to byte slices
	adapterStart = []byte(adapterStartString)
	adapterEnd = []byte(adapterEndString)
	
	// Initialize statistic slices
	adapterStartCount = make([]int, len(adapterStart) + 1)
	adapterEndCount = make([]int, len(adapterEnd) + 1)
	
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
		if argumentError != nil { return }
	}
	
	if *output == "stdout" {
		outputFile = os.Stdout
	} else {
		outputFile, argumentError = os.Create(*output)
		if argumentError != nil { return }
	}
	
	// Create buffered i/o
	inputReader = bufio.NewReader(inputFile)
	outputWriter = bufio.NewWriter(outputFile)
}
