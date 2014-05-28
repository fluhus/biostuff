// Analyzes hot-spots in sequencing results.
package main

import (
	"os"
	"fmt"
	// "bufio"
	"tools"
	"myindex"
	"strdist"
	"seqtools"
	"math/rand"
	"bioformats/fasta"
	// "bioformats/fastq"
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
	// p := fmt.Println

	// Load fasta
	pe("loading fasta...")
	fa, err := fasta.FastaFromFile("data/fasta/Yeast.fa")
	// fa, err := fasta.FastaFromFile("fasta/Yeast.fa")
	if err != nil { panic(err.Error()) }
	
	// Create index
	pe("building index...")
	tools.Tic()
	idx, err := myindex.New(fa, 12, 4)
	pe("took", tools.Toc())
	pe(idx)
	if err != nil { panic(err.Error()) }
	
	// Look up simulated reads
	const numOfReads = 10000
	pe("looking up", numOfReads, "SIMULATED reads...")
	tools.Randomize()
	tools.Tic()
	
	yay := 0
	sure := 0
	sureBad := 0
	unsure := 0
	unsureBad := 0
	
	distances := make([]int, 51)
	// distancesYay := make([]int, 31)
	minuses := 0
	
	for i := 0; i < numOfReads; i++ {
		// Generate random read
		const readLength = 53
		seq, chrom, pos := fa.Subsequence(readLength,
			rand.Intn(fa.NumberOfSubsequences(readLength)))
		if true {//rand.Intn(2) == 0 {
			minuses++
			seq = seqtools.ReverseComplement(seq)
		}
		
		// Mutate
		seq = seqtools.MutateDel(seq, 3)
		
		// Search
		matches := idx.Search(seq, 1, true)
		
		// Pick the best scoring positions
		leaders := scoreLeaders(matches, 1)
		
		// Make a guess
		var guess myindex.GenPos
		if len(leaders) > 1 {
			unsure++
		
			// Pick the position with the least number of SNPs
			guessD := len(seq)
			for _,leader := range leaders {
				upto := min(leader.Pos() + len(seq),
						len(fa[leader.Chr()].Sequence))
				
				guessSeq := fa[leader.Chr()].Sequence[leader.Pos() : upto]
				// if leader.Strand() == myindex.Minus {
					// guessSeq = seqtools.ReverseComplement(guessSeq)
				// }
				ham := strdist.HammingDistanceBytes(seq, guessSeq)
			
				if ham < guessD {
					guess = leader
					guessD = ham
				}
			}
			if guessD < len(distances) {
				distances[guessD]++
			}
			
		} else if len(leaders) == 1 {
			sure++
			guess = leaders[0]
		}
		
		// Test if position is correct
		if guess.Chr() == chrom &&
				abs(guess.Pos() - pos) <= 5 {
			yay++
		} else if len(leaders) == 1 {
			// I was sure but still wrong
			sureBad++
		} else if len(leaders) > 1 {
			// Unsure and wrong
			unsureBad++
		}
	}
	
	pe("took", tools.Toc())
	fmt.Fprintf(os.Stderr,
			"succeeded %.1f%%\n", 100 * float64(yay) / float64(numOfReads))
	fmt.Fprintf(os.Stderr, "sure %.1f%% (%.3f%% success)\n",
			100 * float64(sure) / float64(numOfReads),
			100 * float64(sure - sureBad) / float64(sure))
	fmt.Fprintf(os.Stderr, "unsure %.1f%% (%.3f%% success)\n",
			100 * float64(unsure) / float64(numOfReads),
			100 * float64(unsure - unsureBad) / float64(unsure))
	
	if false {
		pe("minuses:", minuses)
		pe(distances)
	}
}
