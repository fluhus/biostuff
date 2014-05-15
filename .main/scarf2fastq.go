// Converts SCARF formatted files to fastq.
package main

import (
	"os"
	"fmt"
	"bytes"
	"bufio"
	"tools"
)

func main() {
	// Print help
	fmt.Fprintln(os.Stderr, "Converts SCARF files to fastq files.")
	fmt.Fprintln(os.Stderr, "Reads from standard input and writes to standard" +
			" output.\n\nReading standard input...")

	// Prepare buffers
	bufin := bufio.NewReaderSize(os.Stdin, tools.Mega)
	bufout := bufio.NewWriterSize(os.Stdout, tools.Mega)
	
	// Iterate over lines
	for line, err := bufin.ReadBytes('\n'); err == nil;
			line, err = bufin.ReadBytes('\n') {
		// Prepare fastq fields
		line = bytes.Trim(line, "\n\r")
		split := bytes.Split(line, []byte(":"))
		
		qual := split[len(split) - 1]
		seq := split[len(split) - 2]
		id := split[:len(split) - 2]
		
		// Print
		bufout.WriteString("@")
		bufout.Write(bytes.Join(id, []byte(":")))
		bufout.WriteString("\n")
		bufout.Write(seq)
		bufout.WriteString("\n+\n")
		bufout.Write(qual)
		bufout.WriteString("\n")
	}
	
	// Finalize
	bufout.Flush()
	fmt.Fprintln(os.Stderr, "\nDone.")
}
