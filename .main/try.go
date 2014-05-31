package main

import (
	// "os"
	"fmt"
	// "seqtools"
	// "bufio"
	// "bioformats/fastq"
)

func main() {
	var b []byte
	fmt.Printf("len=%d cap=%d\n", len(b), cap(b))
	for i := 0; i < 8; i++ {
		b = append(b, "a"...)
		fmt.Printf("len=%d cap=%d\n", len(b), cap(b))
	}
}
