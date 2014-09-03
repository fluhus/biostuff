package main

import (
	"fmt"
	"bioformats/fasta"
	"time"
	"strdist"
	"math/rand"
	"os"
)

func pe(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func main() {
	pe("start")
	rand.Seed( time.Now().UnixNano() )

	// Read fasta
	fa, err := fasta.FastaFromFile("../aligners/fasta/Yeast.fa")
	if err != nil {
		pe("error: " + err.Error())
		return
	}

	// Take only chromosome 1
	//fa = fa[0:1]

	pe("fasta length:", len(fa))

	//distFunc := strdist.BigramDistance
	const ssl = 70
	const measurements = 10000

	h := make([]int, ssl + 1)
	e := make([]int, ssl + 1)

	n := fa.NumberOfSubsequences(ssl)
	for i := 0; i < measurements; i++ {
		seq1,_,_ := fa.Subsequence(ssl, rand.Intn(n))
		seq2,_,_ := fa.Subsequence(ssl, rand.Intn(n))

		dists[distFunc(seq1, seq2)]++
	}
	
	for i := range dists {
		fmt.Println(dists[i])
	}

	pe("end")
}










