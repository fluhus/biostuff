package main

import (
	"time"
	"math/rand"
	"os"
	"fmt"
	"bioformats/fasta"
	"strdist"
)

func pe(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func pef(s string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, s, a...)
}

func main() {
	pe("start")
	
	fa, err := fasta.FastaFromFile("..\\reference\\yeast.fa")
	if err != nil {
		pe(err.Error)
		return
	}
	
	fmt.Println("loaded", len(fa), "chromosomes")
	
	rand.Seed(time.Now().UnixNano())
	seq,_,_ := fa.Subsequence(50, rand.Intn(fa.NumberOfSubsequences(50)))
	
	distances := make([]int, 60)
	for i := 0; i < 10000; i++ {
		seq2,_,_ := fa.Subsequence(50, rand.Intn(fa.NumberOfSubsequences(50)))
		d := strdist.HammingDistance(seq, seq2)
		distances[d]++
	}
	
	for i,d := range distances {
		pef("%d:\t%d\n", i, d)
	}
	
	pe("end")
}










