// Tests a mapping performance on a SAM file.
package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
	"bioformats/sam"
)

func pe(a ...interface{}) {
	fmt.Fprint(os.Stderr, a...)
	fmt.Fprint(os.Stderr, "\n")
}

func pfe(s string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, s, a...)
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func main() {
	// Open buffer on stdin
	in := bufio.NewReader(os.Stdin)
	
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
		if chr == line.Rname && abs(pos - (line.Pos - 1)) <= 3 {
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
	pfe("\nreads\t\t%.1f%%\t%d\n", 100.0, reads)
	pfe("correct\t\t%.1f%%\t%d\n", float64(goodMaps) / float64(reads) * 100, goodMaps)
	pfe("incorrect\t%.1f%%\t%d\n", float64(badMaps) / float64(reads) * 100, badMaps)
	pfe("unmapped\t%.1f%%\t%d\n", float64(unMaps) / float64(reads) * 100, unMaps)
}





