package main

import (
	"fmt"
	"bioformats/fastq"
	"bufio"
	// "time"
	// "strdist"
	// "math/rand"
	"os"
	// "stat"
)

func pe(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func main() {
	pe("start")
	
	f, err := os.Open("data\\sample.fastq")
	if err != nil {
		pe("error opening fastq:", err.Error())
		return
	}
	
	b := bufio.NewReader(f)
	
	length := -1
	count  := 0
	var quals [][]int
	
	// for fq, err := fastq.ReadNext(b); err == nil; fq, err = fastq.ReadNext(b) {
	for i := 0; i < 100000; i++ {
		fq, _ := fastq.ReadNext(b)
		
		if length == -1 {
			length = len(fq.Quals)
			quals = make([][]int, length)
			pe("read length:", length)
			
			for j := range quals {
				quals[j] = make([]int, 60)
			}
		}
		
		for j := range fq.Quals {
			qual := int( fq.Quals[j] ) - 33
			quals[j][qual]++
		}
		
		count++
	}
	
	pe(count, "reads processed")
	
	// Print statistics
	// for i := range quals {
		// fmt.Printf("%f\t%f\t%f\t%f\n",
				// stat.Mean(quals[i]),
				// stat.Std(quals[i]),
				// stat.Min(quals[i]),
				// stat.Max(quals[i]))
	// }
	
	pe("end")
}










