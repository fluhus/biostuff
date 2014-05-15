/* ****************************************************************************
 * Calculates distance correlations.
 * ****************************************************************************/

package main

import "fmt"
import "bioformats"
import "strdist"
import "tools"
import "stat"

const numOfDists = 100000
const ssl = 50
const fastaPath = "C:\\dev\\Go Workspace\\jo\\data\\fasta\\Yeast.fa"

func main() {
	tools.Randomize()

	// Read fasta
	fmt.Println("reading fasta...")
	fs := bioformats.NewFastaSubsequencer(bioformats.FastaFromFile(fastaPath))
	if fs == nil {panic("Failed to open fasta: " + fastaPath)}
	fs.SetSSLength(ssl)

	fmt.Printf("n=%d\n\n", fs.GetSSCount())
	
	// Create distance arrays
	edit := make([]float64, numOfDists)
	blast := make([]float64, numOfDists)
	bigram := make([]float64, numOfDists)
	hamming := make([]float64, numOfDists)
	
	// Calculate distances
	fmt.Printf("calculating %d distances...\n\n", numOfDists)
	for i := 0; i < numOfDists; i++ {
		// Get random sequences
		s1,_ := fs.GetRandSS()
		s2,_ := fs.GetRandSS()
		
		// Update distances
		edit[i] = float64(strdist.EditDistanceBytes(s1, s2))
		blast[i] = float64(strdist.BlastDistanceBytes(s1, s2,
				strdist.BlastDefaultScores()))
		bigram[i] = float64(strdist.BigramDistanceBytes(s1, s2))
		hamming[i] = float64(strdist.HammingDistanceBytes(s1, s2))
	}
	
	// Calculate correlations
	fmt.Println("correlations:")
	fmt.Println("L=blast E=edit G=bigram H=hamming")
	
	fmt.Printf("LE %.2f\n", stat.Correlation(edit, blast))
	fmt.Printf("LG %.2f\n", stat.Correlation(bigram, blast))
	fmt.Printf("LH %.2f\n", stat.Correlation(hamming, blast))
}







