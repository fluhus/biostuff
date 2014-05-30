// Search scratchpad.
package main

import (
	"os"
	"fmt"
	// "bufio"
	"tools"
	"myindex"
	// "strdist"
	"learning"
	"seqtools"
	"math/rand"
	"bioformats/fasta"
	// "bioformats/fastq"
)

// *** MATH HELPERS ***********************************************************

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

// *** OTHER HELPERS **********************************************************

func pe(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func trimLeft(slice []int) []int {
	for i,v := range slice {
		if v != 0 {
			return slice[i:]
		}
	}
	return nil
}

func randomRead(fa fasta.Fasta) (seq []byte, chr int, pos int) {
	const readLength = 50
	seq, chr, pos = fa.Subsequence(readLength,
		rand.Intn(fa.NumberOfSubsequences(readLength)))
	
	// Minus strand
	if rand.Intn(2) == 0 {
		seq = seqtools.ReverseComplement(seq)
	}
	
	// Mutate
	// seq = seqtools.MutateSNP(seq, 3)
	
	return
}

func intsToString(slice []int) string {
	result := ""
	for i := range slice {
		if i > 0 {
			result += " "
		}
		result += fmt.Sprintf("%d", slice[i])
	}
	return result
}

// *** LEADERS TYPE ***********************************************************

type leadersType [][]myindex.GenPos

func (l leadersType) count() int {
	result := 0
	for _,arr := range l {
		result += len(arr)
	}
	return result
}

func (l leadersType) lens() []int {
	result := make([]int, len(l))
	for i,arr := range l {
		result[i] = len(arr)
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

// *** MAIN *******************************************************************

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
	
	// Variables
	perc := learning.NewSvmUnbiased(3, 0.1)
	
	// Learn from simulated reads
	const numOfReads = 100000
	pe("learning with", numOfReads, "SIMULATED reads...")
	tools.Randomize()
	tools.Tic()
	
	for i := 0; i < numOfReads; i++ {
		// Generate random read
		seq, chr, pos := randomRead(fa)
		
		// Search
		matches := idx.Search(seq, 1, true)
		
		// Pick the best scoring positions
		leaders := scoreLeaders(matches, 2)
		
		// If not found
		if leaders.count() == 0 {
			continue
		}
		
		// If one leading leader, learn how to classify
		if len(leaders[0]) == 1 {
			leader := leaders[0][0]
			
			// If correct
			y := 0
			if leader.Chr() == chr && abs(leader.Pos() - pos) <= 5 {
				y = 1
			
			// If incorrect
			} else {
				y = -1
			}
			
			perc.LearnInt(leaders.lens(), y)
		}
	}
	
	// Test predictions
	pe("testing on", numOfReads, "SIMULATED reads...")
	
	classPosGood := 0
	classPosBad  := 0
	classNegGood := 0
	classNegBad  := 0
	
	for i := 0; i < numOfReads; i++ {
		// Generate random read
		seq, chr, pos := randomRead(fa)
		
		// Search
		matches := idx.Search(seq, 1, true)
		
		// Pick the best scoring positions
		leaders := scoreLeaders(matches, 2)
		
		// If not found
		if leaders.count() == 0 {
			continue
		}
		
		// If one leading leader, classify and predict
		if len(leaders[0]) == 1 {
			leader := leaders[0][0]
			
			// If correct
			if leader.Chr() == chr && abs(leader.Pos() - pos) <= 5 {
				// fmt.Printf("- %v\t%v\n", leader, leaders.lens())
				fmt.Printf("1 %s\n", intsToString(leaders.lens()))
				if c := perc.ClassifyInt(leaders.lens()); c == 1 {
				// if len(leaders[1]) == 0 {
					classPosGood++
				} else if c == -1 {
				// } else {
					classNegGood++
				}
			
			// If incorrect
			} else {
				// fmt.Printf("X %v\t%v\n", leader, leaders.lens())
				fmt.Printf("-1 %s\n", intsToString(leaders.lens()))
				if c := perc.ClassifyInt(leaders.lens()); c == 1 {
				// if len(leaders[1]) == 0 {
					classPosBad++
				} else if c == -1 {
				// } else {
					classNegBad++
				}
			}
		}
	}
	
	pe("took", tools.Toc())
	
	pe("classPosGood", classPosGood)
	pe("classPosBad", classPosBad)
	pe("classNegGood", classNegGood)
	pe("classNegBad", classNegBad)
	
	pe("w:", perc.W())
}
