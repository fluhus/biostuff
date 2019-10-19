// A server that serves fasta sequences.
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fluhus/golgi/bioformats/fasta"
)

func main() {
	// Handle command-line arguments.
	if len(os.Args) == 1 { // No arguments
		fmt.Println("A server for querying fasta files.")
		fmt.Println("\nUsage:\nfastaserver [options] myfile.fasta")
		fmt.Println("\nOptions:")
		flag.PrintDefaults()
		os.Exit(1)
	}
	parseArguments()
	if args.err != nil {
		fmt.Println("Error parsing arguments:", args.err)
		os.Exit(1)
	}

	// Read fasta file.
	fmt.Println("Reading fasta...")
	now := time.Now()
	var err error
	fa, err = readFastaFile(args.file)
	if err != nil {
		fmt.Println("Error reading fasta:", err)
		os.Exit(2)
	}
	reportf("Took %v.\n", time.Now().Sub(now))

	fmt.Print("Ready! Listening on port ", *args.port, ". Hit ctrl+C to"+
		" exit.\n")

	// Listen on port.
	http.HandleFunc("/sequence", sequenceHandler)
	http.HandleFunc("/meta", metaHandler)
	err = http.ListenAndServe(":"+*args.port, nil)
	if err != nil {
		fmt.Println("Error listening:", err)
	}
}

var args struct {
	port    *string
	verbose *bool
	file    string
	err     error
}

// Parses command-line arguments and places everything in args.
// args.err will be non-nil if a parsing error occurred.
func parseArguments() {
	args.port = flag.String("port", "1912", "Port number to listen on.")
	args.verbose = flag.Bool("v", false, "Print out lots of stuff.")
	flag.Parse()

	if len(flag.Args()) == 0 {
		args.err = fmt.Errorf("No fasta input given.")
		return
	}
	if len(flag.Args()) == 0 {
		args.err = fmt.Errorf("Too many arguments.")
		return
	}

	args.file = flag.Arg(0)
}

// Returns a fasta object from the given file.
func readFastaFile(file string) ([]*fasta.Entry, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	return fasta.ReadFasta(f)
}

// All fasta data will be here.
var fa []*fasta.Entry

// Print if verbose.
func report(a ...interface{}) {
	if *args.verbose {
		fmt.Println(a...)
	}
}

// Printf if verbose.
func reportf(s string, a ...interface{}) {
	if *args.verbose {
		fmt.Printf(s, a...)
	}
}
