// Generates a quality model from fastq.
package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"qualmodel"
	"bioformats/fastq"
)

func pe(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func main() {
	// Parse arguments
	if len(os.Args) != 2 && len(os.Args) != 3 {
		pe("Generates a quality model from fastq.\n\n" +
				"Usage:\nqual_learn <fastq> [number of reads]\n\n" +
				"Number of reads may be omitted for reading the entire file." +
				"\n\nThe created model will be printed to stdout, ready to " +
				"be unmarshaled.")
		return
	}
	
	f, err := os.Open(os.Args[1])
	if err != nil {
		pe("error opening fastq:", err.Error())
		return
	}
	
	// Parse number of reads
	numOfReads := -1
	if len(os.Args) == 3 {
		numOfReads, err = strconv.Atoi(os.Args[2])
		
		if err != nil {
			pe("error parsing number of reads:", err.Error())
			return
		}
	}
	
	b := bufio.NewReader(f)
	
	// Helpers
	length := -1
	var counts [][]int
	readCount := 0
	
	// Start iterating
	for fq, err := fastq.ReadNext(b); err == nil; fq, err = fastq.ReadNext(b) {
		// Check if should stop
		if numOfReads > -1 && readCount == numOfReads {
			break
		} else {
			readCount++
		}
	
		// Set length and initialize counts
		if length == -1 {
			length = len(fq.Quals)
			counts = make([][]int, length)
			pe("read length:", length)
			
			for j := range counts {
				counts[j] = make([]int, 60)
			}
		}
		
		for j := range fq.Quals {
			qual := int( fq.Quals[j] ) - 33
			counts[j][qual]++
		}
	}
	
	pe(readCount, "reads processed")
	
	// Create model
	model := qualmodel.New(counts)
	modelText,_ := model.MarshalText()
	fmt.Println(string(modelText))
}










