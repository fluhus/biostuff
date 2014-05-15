/* ****************************************************************************
 * Tests a mapping performance on a SAM file.
 * ****************************************************************************/

package main

import "fmt"
import "flag"
import "strconv"
import "bioformats"
import "strdist"

// Tests mapping performance on a SAM file.
// Expects to receive the path of the target SAM, path of the reference
// fasta and the subsequence length used when creating the reads.
// Returns the numbers of well mapped reads, badly mapped reads and unmapped
// reads. The 3 sum up to the total of lines in the sam.
func testSam(samPath string, fastaPath string, ssLength int) (
	goodMaps, badMaps, unMaps int) {
	// Open fasta
	fs := bioformats.NewFastaSubsequencer(bioformats.FastaFromFile(fastaPath))
	if fs == nil {
		panic(fmt.Sprintf("Could not open fasta: %s", fastaPath))
	}
	fs.SetSSLength(ssLength)

	// Open SAM
	sam := bioformats.NewSamReader(samPath)
	if sam == nil {
		panic(fmt.Sprintf("Could not open SAM: %s", samPath))
	}

	// Read each SAM line
	for line := sam.NextLine(); line != nil; line = sam.NextLine() {
		// Check if unmapped
		if line.Pos == 0 || line.Rname == "*" {
			unMaps++
			continue   // Nothing to check if unmapped, move to next read
		}

		// Extract subsequence number
		ssNumber, sserr := strconv.Atoi(line.Qname)
		if sserr != nil {
			panic("Bad subsequence index format: " + line.Qname)
		}

		// Extract real subsequence
		ssReal := string(fs.GetSS(ssNumber))

		// Extract mapped subsequence (-1 because it is 1 based)
		ssMap := string(fs.GetSS2(line.Rname, line.Pos - 1))

		// Compare distances; if mapped is at least as good as real, success!
		scores := strdist.BlastDefaultScores()
		if strdist.BlastDistance(ssMap, line.Seq, scores) <=
			strdist.BlastDistance(ssReal, line.Seq, scores) {
				goodMaps++
		} else {
			badMaps++
		}
	}

	return
}

func main() {
	// Parse command line arguments
	flag.Parse()
	args := flag.Args()

	if len(args) != 3 {
		const usage = "Tests mapping performance on a map file.\n\n" +
			"Usage:\nsamtester <sam> <fasta ref> <subsequence length>"
		fmt.Println(usage)
		return
	}

	// Extract subsequence length
	ssl, errssl := strconv.Atoi(args[2])
	if errssl != nil {
		panic(fmt.Sprintf("Bad number format for subsequence length: %s",
			args[2]))
	}

	// Test
	a, b, c := testSam(args[0], args[1], ssl)
	fmt.Printf("Good maps:\t%d\nBad maps:\t%d\nUnmaps:\t\t%d\n", a, b, c)
}





