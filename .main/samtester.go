// Tests a mapping performance on a SAM file.
package main

import (
	"os"
	"fmt"
	"bufio"
	"tools"
	"strconv"
	"strings"
	"bioformats/sam"
)

func pe(a ...interface{}) {
	fmt.Fprint(os.Stderr, a...)
	fmt.Fprint(os.Stderr, "\n")
}

func main() {
	// Open buffer on stdin
	in := bufio.NewReaderSize(os.Stdin, tools.Mega)
	
	// Read from stdin
	pe("reading sam from stdin...")
	var err error
	var line *sam.Sam
	reads := 0
	goodMaps := 0
	badMaps  := 0
	unMaps   := 0
	for line, err = sam.ReadNext(in); err == nil;
			line, err = sam.ReadNext(in) {
		reads++
		
		// Split qname to fields
		split := strings.Split(line.Qname, ".")
		chr := strings.Join(split[0 : len(split) - 2], "")
		pos, _ := strconv.Atoi(split[len(split) - 2])
		
		// Check if mapped correctly
		if chr == line.Rname && pos == line.Pos - 1 {
			goodMaps++
		} else {
			if line.Mapq == 0 {
				// Unmapped
				unMaps++
			} else {
				// Badly mapped
				badMaps++
			}
		}
	}
	
	pe("\nerr=", err.Error())
	pe("\nreads\t\t", reads)
	pe("correct\t\t", goodMaps)
	pe("incorrect\t", badMaps)
	pe("unmapped\t", unMaps)
}





