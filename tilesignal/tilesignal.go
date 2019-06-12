// Averages signals along tiles
package main

import (
	"os"
	"fmt"
	"sort"
	"bufio"
	"runtime/pprof"

	"github.com/fluhus/biostuff/bioformats/bed"
	"github.com/fluhus/biostuff/bioformats/bed/bedgraph"
)

// If true, will generate profiling information.
const profiling = false

func main() {
	// Profiling stuff.
	if profiling {
		pout, _ := os.Create("tilesignal.prof")
		defer pout.Close()
		
		pprof.StartCPUProfile(pout)
		defer pprof.StopCPUProfile()
	}

	// Parse arguments.
	if len(os.Args) != 4 {
		fmt.Println("Averages signal value along tiles.")
		fmt.Println("\nWritten by Amit Lavon (amitlavon1@gmail.com).")
		fmt.Println("\nUsage:")
		fmt.Println("tilesignal <signals bedgraph> <in bed> <out bed>")
		os.Exit(1)
	}
	
	bg := os.Args[1]
	bedIn := os.Args[2]
	bedOut := os.Args[3]

	// Read background signals.
	fmt.Println("reading background (this may take a while)...")
	idx, err := newIndex(bg)
	
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(2)
	}
	
	// Process input bed.
	fmt.Println("Processing bed file...")
	err = processBed(bedIn, bedOut, idx)
	
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(2)
	}
	
	fmt.Println("Done!")
}


// ***** BACKGROUND INDEXING **************************************************

// A signle line in the index file.
type signalTile struct {
	start int
	end int
	value float64
}

// An index for fast retrieval of signal levels. Key is the chromosome name.
type signalIndex map[string][]*signalTile

// Creates a new index on the given bed-graph file.
func newIndex(path string) (idx signalIndex, err error) {
	// Open input file.
	f, err := os.Open(path)
	if err != nil { return nil, err }
	
	idx = make(map[string][]*signalTile)
	
	// Scan data.
	scanner := bedgraph.NewScanner(f)
	for scanner.Scan() {
		b := scanner.Bed()
		idx[b.Chr] = append(idx[b.Chr], &signalTile{b.Start, b.End-1, b.Value})
		// End-1 because it's exlusive.
	}
	
	// Sort tiles on each chromosome.
	for chrName, chr := range idx {
		sort.Sort(tileSorter(chr))
		
		// Check for overlaps.
		for i := range chr {
			if i > 0 && chr[i].start <= chr[i-1].end {
				return nil, fmt.Errorf("Overlapping tiles in %s: start=%d" +
						" end=%d", chrName, chr[i].start, chr[i-1].end)
			}
		}
	}
	
	if scanner.Err() != nil { return nil, err }
	
	return
}

// Returns the signal value at the given position. Positions with no signal
// information return 0.
func (idx signalIndex) valueAt(chr string, pos int) float64 {
	chrom, ok := idx[chr]
	if !ok { return 0 }

	i := sort.Search(len(chrom), func(n int) bool {
		return chrom[n].start > pos
	}) - 1
	
	if i == -1 || chrom[i].end < pos {
		return 0
	} else {
		return chrom[i].value
	}
}

// Functions for sorting tiles.
type tileSorter []*signalTile
func (s tileSorter) Len() int {return len(s)}
func (s tileSorter) Less(i, j int) bool {return s[i].start < s[j].start}
func (s tileSorter) Swap(i, j int) {s[i], s[j] = s[j], s[i]}


// ***** BED PROCESSING *******************************************************

func processBed(bedIn, bedOut string, idx signalIndex) error {
	// Open input and output files.
	fin, err := os.Open(bedIn)
	if err != nil { return err }
	defer fin.Close()
	
	fout, err := os.Create(bedOut)
	if err != nil { return err }
	defer fout.Close()
	
	scanner := bed.NewScanner(fin)
	bout := bufio.NewWriter(fout)
	defer bout.Flush()
	
	// Process each tile.
	for scanner.Scan() {
		b := scanner.Bed()
		
		// Measure signal on each base.
		signal := 0.0
		for i := b.Start; i <= b.End; i++ {
			signal += idx.valueAt(b.Chr, i)
		}
		
		// Average.
		signal /= float64(b.End - b.Start + 1)
		
		// Print to output file.
		fmt.Fprintf(bout, "%s\t%f\n", scanner.Text(), signal)
	}
	
	return scanner.Err()
}



