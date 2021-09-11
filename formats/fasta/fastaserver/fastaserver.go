// Command fastaserver is a server that serves sequences from a fasta file.
//
// Requests to /meta return the available sequences and their lengths.
//
// Requests to /sequence return sequences. Parameters are:
//  chr: chromosome name (required)
//  start: sequence start position, 0-based (default: 0)
//  length: number of bases from start to return (default: all)
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fluhus/golgi/formats/fasta"
)

var (
	port    = flag.String("p", "", "Port number to listen on.")
	verbose = flag.Bool("v", false, "Log incoming requests.")
	file    = flag.String("f", "", "Input fasta file.")
)

func main() {
	if err := parseArguments(); err != nil {
		fmt.Println("Error parsing arguments:", err)
		os.Exit(1)
	}

	// Read fasta file.
	log.Println("Reading fasta...")
	now := time.Now()
	var err error
	fa, err = readFastaFile(*file)
	if err != nil {
		fmt.Println("Error reading fasta:", err)
		os.Exit(2)
	}
	log.Println("Took", time.Since(now))

	log.Println("Ready! Listening on port", *port)
	log.Println("Hit ctrl+C to exit")

	// Listen on port.
	http.HandleFunc("/sequence", sequenceHandler)
	http.HandleFunc("/meta", metaHandler)
	err = http.ListenAndServe(":"+*port, nil)
	if err != nil {
		fmt.Println("Error listening:", err)
	}
}

// Parses command-line arguments.
func parseArguments() error {
	flag.Parse()
	if *file == "" {
		return fmt.Errorf("no fasta input given")
	}
	if *port == "" {
		return fmt.Errorf("no port given")
	}
	return nil
}

// Returns a fasta object from the given file.
func readFastaFile(file string) ([]*fasta.Fasta, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := fasta.NewReader(f)

	var fas []*fasta.Fasta
	var fa *fasta.Fasta
	for fa, err = r.Next(); err == nil; fa, err = r.Next() {
		fas = append(fas, fa)
	}
	if err != io.EOF {
		return nil, err
	}
	return fas, nil
}

// All fasta data will be here.
var fa []*fasta.Fasta

// Log if verbose.
func report(a ...interface{}) {
	if *verbose {
		log.Println(a...)
	}
}

// Logf if verbose.
func reportf(s string, a ...interface{}) {
	if *verbose {
		log.Printf(s, a...)
	}
}
