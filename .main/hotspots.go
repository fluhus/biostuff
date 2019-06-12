// Analyzes hot-spots in sequencing results.
package main

import (
	"os"
	"fmt"
	"bufio"
	"tools"
	"myindex"
	// "seqtools"
	"bioformats/fasta"
	"bioformats/fastq"
)

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

func main() {
	p := fmt.Println

	// Load fasta
	pe("loading fasta...")
	fa, err := fasta.FastaFromFile("data/fasta/Yeast.fa")
	// fa, err := fasta.FastaFromFile("fasta/Yeast.fa")
	if err != nil { panic(err.Error()) }
	
	// Create index
	pe("building index...")
	tools.Tic()
	idx, err := myindex.New(fa, 12, 6)
	pe("took", tools.Toc())
	pe(idx)
	if err != nil { panic(err.Error()) }
	
	// Prepare mapping histogram
	const histInterval = 1000
	hist := make([][]int, len(fa))
	for i := range hist {
		hist[i] = make([]int, len(fa[i].Sequence) / histInterval + 1)
	}
	
	// Read fastq
	fastqIn, err := os.Open("data/matan/wt2/s_1_sequence.fastq")
	if err != nil { panic(err.Error()) }
	fastqBuf := bufio.NewReaderSize(fastqIn, tools.Mega)
	
	const numOfReads = 100000
	pe("looking up", numOfReads, "reads...")
	tools.Tic()
	
	i := 0
	yay := 0
	for fq, err := fastq.ReadNext(fastqBuf); err == nil && i < numOfReads;
			fq, err = fastq.ReadNext(fastqBuf) {
		// Search
		matches := idx.Search(fq.Sequence, true)
		
		// Pick the best scoring positions
		leaders := scoreLeaders(matches, 1)
		
		// If at most X found, then we're cool
		if len(leaders) == 1 {
			yay++
			
			// Add to histogram
			hist[leaders[0].Chr()][leaders[0].Pos() / histInterval]++
		}
		
		i++
	}
	
	pe("took", tools.Toc())
	fmt.Fprintf(os.Stderr,
			"mapped %.1f%%\n", 100 * float64(yay) / float64(numOfReads))
	
	// Print histogram
	for i := range hist {
		fmt.Printf("%d ", len(hist[i]))
		for j := range hist[i] {
			fmt.Printf("%d ", hist[i][j])
		}
		p("")
	}
}
