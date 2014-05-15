/* ****************************************************************************
 * A test code for net index.
 * ****************************************************************************/

package main

import "fmt"
import "math/rand"
import "bioformats"
import "strdist"
import "tools"

const numOfSequences = 0
const ssl = 50
const fastaPath = "C:\\dev\\Go Workspace\\jo\\data\\fasta\\chrI.fa"
const netDiamDiff = 500

func dist(s1, s2 []byte) int {
	return strdist.EditDistanceBytes(s1, s2)
}

func netCovers(net [][]byte, sequence []byte, rad int) bool {
	for i := range net {
		d := dist(net[i], sequence)
		if d <= rad {return true}
	}
	
	return false
}

func getFarthest(sequence []byte, sequences [][]byte) []byte {
	// Initialize result
	var result []byte
	d := -1
	
	// Pick the farthest
	for i := range sequences {
		d2 := dist(sequences[i], sequence)
		if d < d2 {
			d = d2
			result = sequences[i]
		}
	}
	
	return result
}

func makeNet(sequences [][]byte, rad int) [][]byte {
	// Initialize result
	var net [][]byte
	
	// Pick random start point
	r := rand.Intn(len(sequences))
	for i := range sequences {
		// Sequence index
		ii := (i + r) % len(sequences)
		
		// Check if covered
		if !netCovers(net, sequences[ii], rad) {
			net = append(net, sequences[ii])
		}
	}
	
	return net
}

func splitByNet(sequences, net [][]byte) [][][]byte {
	// Distance array
	d := make([]int, len(net))
	
	// Result array
	result := make([][][]byte, len(net))
	
	// For each sequence
	for i := range sequences {
		for j := range net {
			d[j] = dist(net[j], sequences[i])
		}
		
		// Pick closest
		m := tools.ArgMinInt(d...)
		result[m] = append(result[m], sequences[i])
	}
	
	return result
}

func getDiameter(sequences [][]byte) int {
	// Pick random pivot
	pivot := sequences[rand.Intn(len(sequences))]

	// Calculate distances
	s1 := getFarthest(pivot, sequences)
	s2 := getFarthest(s1, sequences)
	
	return dist(s1, s2)
}

func main() {
	tools.Randomize()

	// Read fasta
	fmt.Println("reading fasta")
	fs := bioformats.NewFastaSubsequencer(bioformats.FastaFromFile(fastaPath))
	if fs == nil {panic("Failed to open fasta: " + fastaPath)}
	fs.SetSSLength(ssl)

	fmt.Printf("n=%d\n", fs.GetSSCount())
	
	// Create sequence array
	var seqs [][]byte
	if numOfSequences == 0 {seqs = make([][]byte, fs.GetSSCount())
	} else {seqs = make([][]byte, numOfSequences)}
	
	for i := range seqs {
		seqs[i] = fs.GetSS(i)
	}
	
	// Get diameter
	fmt.Println("calculating diameter")
	diam := getDiameter(seqs)
	fmt.Printf("diam=%d\n", diam)
	
	// Create a net
	fmt.Printf("creating net")
	netRad := diam * 3 / 4 //- netDiamDiff
	fmt.Printf(" [nr=%d]\n", netRad)
	net := makeNet(seqs, netRad)
	fmt.Printf("size=%d\n", len(net))
	
	// Check net size
	if len(net) == 1 {
		fmt.Println("net size is 1. aborting...")
		return
	}
	
	// Split sequences
	fmt.Println("splitting...")
	split := splitByNet(seqs, net)
	fmt.Println("diameters(size):")
	for i := range split {
		fmt.Printf("%d(%d), ", getDiameter(split[i]), len(split[i]))
	}
	fmt.Println("")
}







