// A server that serves fasta sequences.
//
// Request format: chr,start,length; where start is 0-based.
//
// Response format: 0sequence or 1error-message
package main

import (
	"net"
	"myflag"
	"fmt"
	"bioformats/fasta"
	"os"
	"bytes"
	"strconv"
)

func main() {
	// Handle command-line arguments.
	parseArguments()
	if !myflag.HasAny() {
		fmt.Println("A server for querying fasta files.")
		fmt.Println("\nUsage:\nfastaserver [options] myfile.fasta")
		fmt.Println("\nOptions:")
		fmt.Print(myflag.HelpString())
		os.Exit(1)
	}
	if args.err != nil {
		fmt.Println("Error parsing arguments:", args.err)
		os.Exit(1)
	}
	
	// Read fasta file.
	fmt.Println("Reading fasta...")
	var err error
	fa, err = readFastaFile(args.file)
	if err != nil {
		fmt.Println("Error reading fasta:", err)
		os.Exit(2)
	}
	
	// Listen on port.
	ln, err := net.Listen("tcp", ":" + args.port)
	if err != nil {
		fmt.Println("Error opening port:", err)
		os.Exit(2)
	}
	
	fmt.Print("Ready! Listening on port ", args.port, ". Hit ctrl+C to" +
			" exit.\n")
	
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		
		go handleConnection(conn)
	}
}

var args struct {
	port string
	verbose bool
	file string
	err error
}

// Parses command-line arguments and places everything in args.
// args.err will be non-nil if a parsing error occurred.
func parseArguments() {
	port := myflag.String("port", "p", "number",
			"Port number to listen on. Default: 1912.", "1912")
	verbose := myflag.Bool("verbose", "v", "Print lots of stuff.", false)
	
	args.err = myflag.Parse()
	if args.err != nil { return }
	
	args.port = *port
	args.verbose = *verbose
	
	if len(myflag.Args()) == 0 {
		args.err = fmt.Errorf("No fasta input given.")
		return
	}
	if len(myflag.Args()) == 0 {
		args.err = fmt.Errorf("Too many arguments.")
		return
	}
	
	args.file = myflag.Args()[0]
}

// Returns a fasta object from the given file.
func readFastaFile(file string) ([]*fasta.FastaEntry, error) {
	f, err := os.Open(file)
	if err != nil { return nil, err }
	
	return fasta.ReadFasta(f)
}

// All fasta data will be here.
var fa []*fasta.FastaEntry

// Handles a single request from a fasta client.
func handleConnection(c net.Conn) {
	report("Got connection.")
	defer c.Close()

	// Read until uncountering ';'.
	msg := bytes.NewBuffer(nil)
	b := []byte{0}
	for {
		_, err := c.Read(b)
		if err != nil {
			fmt.Fprintf(c, "1Error reading message: %s", err)
			return
		}
		
		// Break upon ';'.
		if b[0] == ';' { break }
		msg.Write(b)
		
		// Too long means error.
		if msg.Len() > 1000 {
			fmt.Fprintf(c, "1Message is too long. Max is 1000.")
			return
		}
	}
	report("Message is:", msg.String())
	
	// Parse message fields.
	words := bytes.Split(msg.Bytes(), []byte(","))
	if len(words) != 3 {
		fmt.Fprintf(c, "1Bad message format. Expected 3 fields and got %d.",
				len(words))
		return
	}
	
	// Parse message fields.
	chr := words[0]
	start, err := strconv.Atoi(string(words[1]))
	if err != nil || start < 0 {
		fmt.Fprintf(c, "1Bad start position: %s.", words[1])
		return
	}
	length, err := strconv.Atoi(string(words[2]))
	if err != nil || length < 1 {
		fmt.Fprintf(c, "1Bad length: %s.", words[2])
		return
	}
	
	// Find fasta entry.
	var entry *fasta.FastaEntry
	for _, e := range fa {
		if e.Name() == string(chr) {
			entry = e
		}
	}
	
	if entry == nil {
		fmt.Fprintf(c, "1No such chromosome: '%s'.", chr)
		return
	}
	
	if start + length > entry.Length() {
		fmt.Fprintf(c, "1Position exceeds chromosome length (max %d).",
				entry.Length())
		return
	}
	
	reportf("chr=%s start=%d len=%d\n", chr, start, length)
	seq := entry.Subsequence(start, length)
	fmt.Fprintf(c, "0%s", seq)
}

// Print if verbose.
func report(a ...interface{}) {
	if args.verbose {
		fmt.Println(a...)
	}
}

// Printf if verbose.
func reportf(s string, a ...interface{}) {
	if args.verbose {
		fmt.Printf(s, a...)
	}
}
