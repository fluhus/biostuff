// Simulates reads from a given fasta file.
package main

import (
	"fmt"
	"tools"
	"seqtools"
	"math/rand"
	"bioformats/fasta"
	"bioformats/fastq"
)

func main() {
	tools.Randomize()

	// Open fasta
	fa, err := fasta.FastaFromFile("data/fasta/Yeast.fa")
	if err != nil { panic(err.Error()) }
	
	const ssl = 53
	const numOfReads = 10000
	numOfSubsequences := fa.NumberOfSubsequences(ssl)
	
	for i := 0; i < numOfReads; i++ {
		seq, chr, pos :=
				fa.Subsequence(ssl, rand.Intn(numOfSubsequences))
		
		strand := 0
		if rand.Intn(2) == 0 {
			strand = 1
			seq = seqtools.ReverseComplement(seq)
		}
		
		seq = seqtools.MutateDel(seq, 3)
		
		quals := fastq.MakeQuals(seq)
		id := fmt.Sprintf("%s.%d.%d", fa[chr].Title, pos, strand)
		
		fmt.Println(&fastq.Fastq{[]byte(id), []byte(seq), []byte(quals)})
	}
}