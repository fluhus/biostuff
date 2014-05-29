// Search scratchpad.
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

type leadersType [][]myindex.GenPos

func (l leadersType) count() int {
	result := 0
	for _,arr := range l {
		result += len(arr)
	}
	return result
}

func (l leadersType) toSlice() []myindex.GenPos {
	var result []myindex.GenPos
	for _,arr := range l {
		result = append(result, arr...)
	}
	return result
}

func scoreLeaders(matches map[myindex.GenPos]int,
		howLess int) leadersType {
	// Find max
	max := 0
	for _,score := range matches {
		if score > max {
			max = score
		}
	}
	
	result := make([][]myindex.GenPos, howLess + 1)
	for match, score := range matches {
		if max - score <= howLess {
			result[max - score] = append(result[max - score], match)
		}
	}
	
	return result
}

func main() {
	// p := fmt.Println

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
	subsure := 0
	subsureBad := 0
	
	for i := 0; i < numOfReads; i++ {
		// Generate random read
		const readLength = 53
		seq, chrom, pos := fa.Subsequence(readLength,
			rand.Intn(fa.NumberOfSubsequences(readLength)))
		if rand.Intn(2) == 0 {
			seq = seqtools.ReverseComplement(seq)
		}
		
		// Mutate
		seq = seqtools.MutateDel(seq, 3)
		
		// Search
		matches := idx.Search(seq, 1, true)
		
		// Pick the best scoring positions
		leaders := scoreLeaders(matches, 2)
		
		// If not found
		if leaders.count() == 0 {
			continue
		}
		
		// Make a guess
		var guess myindex.GenPos
		var guesses []myindex.GenPos
		guessD := len(seq)
		sureGuess := false
		if len(leaders[0]) == 1 && len(leaders[1]) == 0 {
			sure++
			sureGuess = true
			guess = leaders[0][0]
		} else {
			unsure++
		
			// Pick the position with the least number of SNPs
			for _,leader := range leaders.toSlice() {
				upto := min(leader.Pos() + len(seq),
						len(fa[leader.Chr()].Sequence))
				
				guessSeq := fa[leader.Chr()].Sequence[leader.Pos() : upto]
				if leader.Strand() == myindex.Minus {
					guessSeq = seqtools.ReverseComplement(guessSeq)
				}
				
				ham := strdist.HammingDistance(seq, guessSeq) //+
						// strdist.EditDistance(seq, guessSeq)
			
				if ham < guessD {
					guesses = []myindex.GenPos{leader}
					guessD = ham
				} else if ham == guessD {
					guesses = append(guesses, leader)
				}
			}
			
			// Pick one at random
			guess = guesses[rand.Intn(len(guesses))]
			if len(guesses) == 1 {subsure++}
		}
			
		// Test if position is correct
		if guess.Chr() == chrom &&
				abs(guess.Pos() - pos) <= 5 {
			yay++
			
			if !sureGuess && len(guesses) == 1 {
				fmt.Printf("- d=%d guesses=%d leaders=%d\tseq=%s\n",
						guessD, len(guesses), len(leaders), seq)
			}
		} else if sureGuess {
			// I was sure but still wrong
			sureBad++
		} else {
			// Unsure and wrong
			unsureBad++
			if len(guesses) == 1 {
				fmt.Printf("X d=%d guesses=%d leaders=%d\tseq=%s\n",
						guessD, len(guesses), len(leaders), seq)
				subsureBad++
			}
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
	fmt.Fprintf(os.Stderr, "subsure %.1f%% (%.3f%% success)\n",
			100 * float64(subsure) / float64(numOfReads),
			100 * float64(subsure - subsureBad) / float64(subsure))
	// pe("subsure", subsure, "subsureBad", subsureBad)
}
