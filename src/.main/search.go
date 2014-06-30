// Search sandbox.
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

func pef(s string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, s, a...)
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
	const readLength = 47
	seq, chr, pos = fa.Subsequence(readLength,
		rand.Intn(fa.NumberOfSubsequences(readLength)))
	
	// Minus strand
	if rand.Intn(2) == 0 {
		seq = seqtools.ReverseComplement(seq)
	}
	
	// Mutate
	// seq = seqtools.MutateIns(seq, 3)
	
	return
}

func distanceTo(seq []byte, pos myindex.GenPos, fa fasta.Fasta) int {
	chr  := pos.Chr()
	from := pos.Pos()
	to   := min(from + len(seq), len(fa[chr].Sequence))
	
	return strdist.HammingDistance(seq, fa[chr].Sequence[from:to])
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

func floatsToString(slice []float64) string {
	result := ""
	for i := range slice {
		if i > 0 {
			result += " "
		}
		result += fmt.Sprintf("%f", slice[i])
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

func distanceLeaders(seq []byte, matches []myindex.GenPos,
		fa fasta.Fasta, howMore int) leadersType {
	// Find min
	min := len(seq)
	
	for i := range matches {
		dist := distanceTo(seq, matches[i], fa)
		if dist < min {
			min = dist
		}
	}
	
	// Gather leaders
	result := make([][]myindex.GenPos, howMore + 1)
	for i := range matches {
		if diff := distanceTo(seq, matches[i], fa) - min; diff <= howMore {
			result[diff] = append(result[diff], matches[i])
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
	idx, err := myindex.New(fa, 12, 1)
	pe("took", tools.Toc())
	pe(idx)
	if err != nil { panic(err.Error()) }
	
	// Variables
	sureGood := 0
	sureBad := 0
	unsureGood := 0
	unsureBad := 0
	unsureUnsure := 0
	
	// Seach simulated reads
	const numOfReads = 10000
	pe("learning with", numOfReads, "SIMULATED reads...")
	tools.Randomize()
	tools.Tic()
	
	for i := 0; i < numOfReads; i++ {
		// Generate random read
		seq, chr, pos := randomRead(fa)
		if rand.Intn(1) == 0 { seq = seqtools.ReverseComplement(seq) }
		
		// Search
		matches := idx.Search(seq, 1, true)
		if len(matches) == 0 {
			continue
		}
		
		// Pick the best scoring positions
		leaders := scoreLeaders(matches, 1)
		
		// Sure
		if leaders.count() == 1 {
			leader := leaders[0][0]
			
			if leader.Chr() == chr && abs(leader.Pos() - pos) <= 5 {
				sureGood++
			} else {
				sureBad++
			}
		
		// Unsure
		} else {
			leaders = distanceLeaders(seq, leaders.toSlice(), fa, 2)
			if leaders.count() == 1 {
				leader := leaders[0][0]
				
				if leader.Chr() == chr && abs(leader.Pos() - pos) <= 5 {
					unsureGood++
				} else {
					unsureBad++
				}
			} else {
				unsureUnsure++
			}
		}
	}
	
	pe("took", tools.Toc())
	
	pef("sure   + %4d (%4.1f%%)\n", sureGood, float64(sureGood) / float64(numOfReads) * 100)
	pef("sure   - %4d (%4.1f%%)\n", sureBad, float64(sureBad) / float64(numOfReads) * 100)
	pef("unsure + %4d (%4.1f%%)\n", unsureGood, float64(unsureGood) / float64(numOfReads) * 100)
	pef("unsure - %4d (%4.1f%%)\n", unsureBad, float64(unsureBad) / float64(numOfReads) * 100)
	pef("UNSURE   %4d (%4.1f%%)\n", unsureUnsure, float64(unsureUnsure) / float64(numOfReads) * 100)
}
