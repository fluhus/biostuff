package main

// Parses command line arguments.

import (
	"flag"
	"fmt"
	// "os"
	"io/ioutil"
)

func init() {
	flags := flag.NewFlagSet("trimmer", flag.ContinueOnError)
	flags.SetOutput(ioutil.Discard)
	
	input := flags.String("in", "", "Input fastq file. Default: stdin.")
	output := flags.String("out", "", "Output fastq file. Default: stdout.")
	adapterStart := flags.String("adapter-start", "",
			"Adapter to trim at the beginning (5') of the read. Default: none.")
	adapterEnd := flags.String("adapter-end", "",
			"Adapter to trim at the end (3') of the read. Default: none.")
	phredOffset := flags.Int("phred-offset", 33,
			"Phred quality score offset. Default: 33.")
	qualThreshold := flags.Uint("qual-threshold", 20,
			"Quality trimmming threshold. Give 0 to avoid quality trimming." +
			" Default: 20")
	
	// flags.Parse(os.Args[1:])
	err := flags.Parse([]string{ "--phred-offset", "12", "-adapter-start", "oolool", "yoink" })
	fmt.Println("err:", err)
	fmt.Println(*input, *output, *adapterStart, *adapterEnd, *phredOffset,
			*qualThreshold, flags.Args())
	flags.VisitAll(visitor)
}

// Visits flags and prints their usage in a user-friendly format.
func printUsage(f *flag.Flag) {
	fmt.Printf("\t--%s\n\t\t%s\n", f.Name, f.Usage)
}

