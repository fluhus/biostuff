package main

import (
	"fmt"
	"bioformats/fastq"
	"bufio"
	// "time"
	// "strdist"
	// "math/rand"
	"os"
	"stat"
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
	var quals1 []float64
	var quals2 []float64
	
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
			
			if j == 40 { quals1 = append(quals1, float64(qual)) }
			if j == 60 { quals2 = append(quals2, float64(qual)) }
		}
		
		count++
	}
	
	pe(count, "reads processed")
	
	// pe(quals[100])
	pe("corr:", stat.Correlation(quals1, quals2))
	
	pe("end")
}










