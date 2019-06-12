// A server that serves fasta sequences.
package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/fluhus/biostuff/bioformats/fasta"
	"github.com/fluhus/biostuff/gobz"
	"github.com/fluhus/biostuff/myflag"
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
	now := time.Now()
	var err error
	fa, err = readFastaFile(args.file)
	if err != nil {
		fmt.Println("Error reading fasta:", err)
		os.Exit(2)
	}
	reportf("Took %v.\n", time.Now().Sub(now))

	if *args.serialize {
		fmt.Println("Serializing fasta...")
		now = time.Now()
		newFile, err := serializeFasta()
		if err != nil {
			fmt.Println("Error serializing fasta:", err)
			os.Exit(2)
		} else {
			reportf("Took %v.\n", time.Now().Sub(now))
			fmt.Printf("Serialization successful! New file is: %s\n", newFile)
			return
		}
	}

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
	port      *string
	verbose   *bool
	serialize *bool
	file      string
	err       error
}

// Parses command-line arguments and places everything in args.
// args.err will be non-nil if a parsing error occurred.
func parseArguments() {
	args.port = myflag.String("port", "p", "number",
		"Port number to listen on. Default: 1912.", "1912")
	args.verbose = myflag.Bool("verbose", "v", "Print lots of stuff.", false)
	args.serialize = myflag.Bool("serialize", "s",
		"Serialize a fasta file for fast loading. Generates a file with "+
			"the same name with a '.serialized' suffix.", false)

	args.err = myflag.Parse()
	if args.err != nil {
		return
	}

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
func readFastaFile(file string) ([]*fasta.Entry, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	// Is input file serialized?
	if strings.HasSuffix(args.file, ".serialized") {
		// Yes! Deserialize.
		var fas []*fasta.SerializableEntry
		err := gobz.Load(args.file, &fas)
		if err != nil {
			return nil, err
		}

		// Convert to regular fasta.
		fa := make([]*fasta.Entry, len(fas))
		for i := range fas {
			fa[i] = fasta.FromSerializable(fas[i])
		}

		// Release unused memory.
		debug.FreeOSMemory()

		return fa, nil
	} else {
		// No! Read textual fasta.
		return fasta.ReadFasta(f)
	}
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

// Serializes the input fasta file.
func serializeFasta() (string, error) {
	// Convert to serializable.
	fas := make([]*fasta.SerializableEntry, len(fa))
	for i := range fas {
		fas[i] = fasta.ToSerializable(fa[i])
	}

	newFile := args.file + ".serialized"
	err := gobz.Save(newFile, fas)

	return newFile, err
}
