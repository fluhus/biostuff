// Aligner prototype.
package main

import (
	"os"
	"fmt"
	"bufio"
	"tools"
	"myindex"
	// "seqtools"
	// "math/rand"
	"bioformats/sam"
	"bioformats/fasta"
	"bioformats/fastq"
)

func max(a, b int) int {
	if a < b { return b }
	return a
}

func min(a, b int) int {
	if a > b { return b }
	return a
}

func abs(a int) int {
	if a < 0 { return -a }
	return a
}

func pe(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func scoreDist(matches map[myindex.GenPos]int) []int {
	// Find max
	max := 0
	for _,score := range matches {
		if score > max {
			max = score
		}
	}
	
	// Create distribution
	result := make([]int, max+1)
	for _,score := range matches {
		result[score]++
	}
	
	return result
}

func scoreLeaders(matches map[myindex.GenPos]int,
		howLess int) []myindex.GenPos {
	// Find max
	max := 0
	for _,score := range matches {
		if score > max {
			max = score
		}
	}
	
	var result []myindex.GenPos
	for match, score := range matches {
		if score >= max - howLess {
			result = append(result, match)
		}
	}
	
	return result
}

func hamming(s1, s2 []byte) int {
	// Make sure s1 is the longer
	if len(s2) > len(s1) {
		s1, s2 = s2, s1
	}
	
	result := len(s1) - len(s2)
	
	// Count differences (iterate over shorter sequence)
	for i := range s2 {
		if s1[i] != s2[i] {
			result++
		}
	}
	
	return result
}

func main() {
	// Load fasta
	pe("loading fasta...")
	fa, err := fasta.FastaFromFile("data/fasta/Yeast.fa")
	if err != nil { panic(err.Error()) }
	
	// Create index
	pe("building index...")
	tools.Tic()
	idx, err := myindex.New(fa, 12, 4)
	pe("took", tools.Toc())
	pe(idx)
	if err != nil { panic(err.Error()) }
	
	// Open fastq
	in := bufio.NewReaderSize(os.Stdin, tools.Mega)
	
	// Look up reads
	pe("looking up reads (standard input)...")
	tools.Tic()
	var fq *fastq.Fastq
	
	for fq, err = fastq.ReadNext(in); err == nil;
			fq, err = fastq.ReadNext(in) {
		// Search
		matches := idx.Search(fq.Sequence, 1, true)
		
		// Pick the best scoring positions
		leaders := scoreLeaders(matches, 1)
		
		// If mapped
		var samLine sam.Sam
		if len(leaders) == 1 {
			leader := leaders[0]
			samLine.Rname = string(fa[leader.Chr()].Title)
			samLine.Pos = leader.Pos() + 1 // sam are 1-based
			samLine.Mapq = 60
		
		// If not mapped
		} else {
			samLine.Rname = "*"
			samLine.Pos = 0
			samLine.Mapq = 0
		}
		
		samLine.Qname = string(fq.Id)
		samLine.Cigar = "*"
		samLine.Rnext = "*"
		samLine.Pnext = "*"
		samLine.Seq =  string(fq.Sequence)
		samLine.Qual = string(fq.Quals)
		
		fmt.Println(samLine)
	}
	
	pe("took", tools.Toc())
	pe("error:", err.Error())
}
