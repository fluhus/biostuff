/* ****************************************************************************
 * Distance statistics
 * ****************************************************************************/

package main

import "fmt"
import "bioformats"
import "strdist"
import "tools"
import "stat"

const numOfDists = 1000000
const ssl = 50
const fastaPath = "C:\\dev\\Go Workspace\\jo\\data\\fasta\\Yeast.fa"

func dist(s1, s2 []byte) int {
	return strdist.BigramDistanceBytes(s1, s2)
	//return strdist.BlastDistanceBytes(s1, s2, strdist.BlastDefaultScores())
}

func main() {
	tools.Randomize()

	// Read fasta
	fmt.Println("reading fasta...")
	fs := bioformats.NewFastaSubsequencer(bioformats.FastaFromFile(fastaPath))
	if fs == nil {panic("Failed to open fasta: " + fastaPath)}
	fs.SetSSLength(ssl)

	fmt.Printf("n=%d\n\n", fs.GetSSCount())
	
	// Create distance arrays
	distances := make([]float64, numOfDists)
	distribution := make([]int, ssl+1)
	
	// Calculate distances
	fmt.Printf("calculating %d distances...\n\n", numOfDists)
	for i := 0; i < numOfDists; i++ {
		// Get random sequences
		s1,_ := fs.GetRandSS()
		s2,_ := fs.GetRandSS()
		
		// Update distances
		d := dist(s1, s2)
		distances[i] = float64(d)
		//if d+10 < 0 || d+10 >= len(distribution) {fmt.Println("d=", d+10)}
		distribution[d]++
	}
	
	// Print statistics
	fmt.Printf("mean=%.2f\nstd=%.2f\n",
			stat.Mean(distances), stat.Std(distances))
	
	fd := make([]float64, len(distribution))
	for i := range distribution {
		fd[i] = float64(distribution[i])
	}
	
	sum := stat.Sum(fd)
	for i := range fd {
		fmt.Printf("%d\t%.6f\n", i, fd[i]/sum)
	}
}







