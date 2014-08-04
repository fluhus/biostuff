// Simulates reads from a given fasta file.
package main

import (
	"os"
	"fmt"
	"tools"
	"strconv"
	"seqtools"
	"math/rand"
	"bioformats/fasta"
	"bioformats/fastq"
)

func pe(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

const usage = "Usage:\nreadsim <fasta> <number of reads per mutation> " +
		"<read length> <mutation amplitude>"

func main() {
	// Parse arguments (fasta, number of reads, length, mutation amplitude)
	if len(os.Args) != 5 { pe(usage); return }
	
	fastaFile := os.Args[1]
	
	numReads, err := strconv.ParseInt(os.Args[2], 0, 64)
	if err != nil { pe(err.Error(), "\n" + usage); return }
	
	readLength, err := strconv.ParseInt(os.Args[3], 0, 64)
	if err != nil { pe(err.Error(), "\n" + usage); return }

	mutAmp, err := strconv.ParseInt(os.Args[4], 0, 64)
	if err != nil { pe(err.Error(), "\n" + usage); return }

	tools.Randomize()

	// Open fasta
	fa, err := fasta.FastaFromFile(fastaFile)
	if err != nil { pe(err.Error(), "\n", usage); return }
	
	// Generate unmutated
	numOfSubsequences := fa.NumberOfSubsequences(int(readLength))
	
	for i := int64(0); i < numReads; i++ {
		seq, chr, pos :=
				fa.Subsequence(int(readLength), rand.Intn(numOfSubsequences))
		
		strand := rand.Intn(2)
		if strand == 1 {
			seq = seqtools.ReverseComplement(seq)
		}
		
		quals := fastq.MakeQuals(seq)
		id := fmt.Sprintf("%s.%d.%d", fa[chr].Title, pos, strand)
		
		fmt.Println(&fastq.Fastq{[]byte(id), seq, quals})
	}
	
	// Generate SNPed
	for i := int64(0); i < numReads; i++ {
		seq, chr, pos :=
				fa.Subsequence(int(readLength), rand.Intn(numOfSubsequences))
		
		strand := rand.Intn(2)
		if strand == 1 {
			seq = seqtools.ReverseComplement(seq)
		}
		
		seq = seqtools.MutateSNP(seq, int(mutAmp))
		
		quals := fastq.MakeQuals(seq)
		id := fmt.Sprintf("%s.%d.%d", fa[chr].Title, pos, strand)
		
		fmt.Println(&fastq.Fastq{[]byte(id), seq, quals})
	}
	
	// Generate insertions
	numOfSubsequences = fa.NumberOfSubsequences(int(readLength - mutAmp))
	
	for i := int64(0); i < numReads; i++ {
		seq, chr, pos :=
				fa.Subsequence(int(readLength - mutAmp),
				rand.Intn(numOfSubsequences))
		
		strand := rand.Intn(2)
		if strand == 1 {
			seq = seqtools.ReverseComplement(seq)
		}
		
		seq = seqtools.MutateIns(seq, int(mutAmp))
		
		quals := fastq.MakeQuals(seq)
		id := fmt.Sprintf("%s.%d.%d", fa[chr].Title, pos, strand)
		
		fmt.Println(&fastq.Fastq{[]byte(id), seq, quals})
	}
	
	// Generate deletions
	numOfSubsequences = fa.NumberOfSubsequences(int(readLength + mutAmp))
	
	for i := int64(0); i < numReads; i++ {
		seq, chr, pos :=
				fa.Subsequence(int(readLength + mutAmp),
				rand.Intn(numOfSubsequences))
		
		strand := rand.Intn(2)
		if strand == 1 {
			seq = seqtools.ReverseComplement(seq)
		}
		
		seq = seqtools.MutateDel(seq, int(mutAmp))
		
		quals := fastq.MakeQuals(seq)
		id := fmt.Sprintf("%s.%d.%d", fa[chr].Title, pos, strand)
		
		fmt.Println(&fastq.Fastq{[]byte(id), seq, quals})
	}
}
