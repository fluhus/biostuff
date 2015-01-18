package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"sort"
	"flag"
	"io/ioutil"
)

func main() {
	// Parse arguments
	if len(os.Args) == 1 {
		fmt.Println(usage)
		return
	}
	
	parseArguments()
	if arguments.err != nil {
		fmt.Println("Error parsing arguments:", arguments.err)
		os.Exit(1)
	}
	
	// Load genes
	fmt.Println("Loading genes from:", arguments.genesFile)
	idx, err := loadGenes(arguments.genesFile)
	if err != nil {
		fmt.Println("Error loading genes:", err)
		os.Exit(2)
	}
	
	// Attach to input file
	fmt.Println("Attaching nearest genes to:", arguments.inFile)
	err = attachGenes(arguments.inFile, arguments.outFile, idx, arguments.n)
	if err != nil {
		fmt.Println("Error attaching genes to file:", err)
		os.Exit(2)
	}
}

// ***** ARGUMENT PARSING *****************************************************

var arguments struct {
	genesFile string
	inFile string
	outFile string
	n int
	err error
}

// Parses input arguments. arguments.err will hold the parsing error,
// if encountered.
func parseArguments() {
	// Create flag set
	flags := flag.NewFlagSet("", flag.ContinueOnError)
	flags.SetOutput(ioutil.Discard)
	
	// Register arguments
	flags.StringVar(&arguments.genesFile, "genes", "", "")
	flags.StringVar(&arguments.genesFile, "g", "", "")
	flags.StringVar(&arguments.inFile, "in", "", "")
	flags.StringVar(&arguments.inFile, "i", "", "")
	flags.StringVar(&arguments.outFile, "out", "", "")
	flags.StringVar(&arguments.outFile, "o", "", "")
	flags.IntVar(&arguments.n, "n", 1, "")
	
	// Parse!
	arguments.err = flags.Parse(os.Args[1:])
	if arguments.err != nil { return }
	
	// Check argument validity
	if arguments.genesFile == "" {
		arguments.err = fmt.Errorf("No genes file selected")
		return
	}
	
	if arguments.inFile == "" {
		arguments.err = fmt.Errorf("No input file selected")
		return
	}
	
	if arguments.outFile == "" {
		arguments.err = fmt.Errorf("No output file selected")
		return
	}
	
	if arguments.n < 1 {
		arguments.err = fmt.Errorf("Number of genes must be positive, got %d",
				arguments.n)
	}
	
	if len(flags.Args()) > 0 {
		arguments.err = fmt.Errorf("Unknown argument: %s", flag.Args()[0])
		return
	}
}


// ***** GENES OBJECTS ********************************************************

// A single gene
type gene struct {
	name string
	start int
	end int
}

// Gene lists for a single chromosome. One sorted by start position and
// one by end.
type chromIndex struct {
	byStart []*gene
	byEnd []*gene
}

// From chromosome name to index.
type index map[string]*chromIndex

// Adds the given gene to the index.
func (idx index) add(chr string, name string, start int, end int) {
	// Create entry
	if idx[chr] == nil {
		idx[chr] = &chromIndex{}
	}
	
	idx[chr].byStart = append(idx[chr].byStart, &gene{name, start, end})
}

// Creates an index of genes.
func loadGenes(path string) (index, error) {
	// Open genes file
	fgenes, err := os.Open(path)
	if err != nil { return nil, err }
	defer fgenes.Close()
	scanner := bufio.NewScanner(fgenes)
	
	// Ignore header line
	scanner.Scan()
	
	// Load genes
	result := index(make(map[string]*chromIndex))
	for scanner.Scan() {
		// Split line
		fields := strings.Split(scanner.Text(), "\t")
		if len(fields) != 4 {
			return nil, fmt.Errorf("Bad number of fields: %d, expected 4",
					len(fields))
		}
		
		// Parse fields
		chr := fields[0]
		name := fields[3]
		
		start, err := strconv.Atoi(fields[1])
		if err != nil { return nil, err }
		end, err := strconv.Atoi(fields[2])
		if err != nil { return nil, err }
		
		result.add(chr, name, start, end)
	}
	
	result.sort()
	
	return result, nil
}


// ***** GENES INDEXING *******************************************************

// Sorts the entries in the index, which makes it available for use.
func (idx index) sort() {
	for _,chr := range idx {
		chr.sort()
	}
}

// Sorts the entries in the index, which makes it available for use.
func (idx *chromIndex) sort() {
	// Copy byStart entries
	idx.byEnd = make([]*gene, len(idx.byStart))
	copy(idx.byEnd, idx.byStart)
	
	// Sort index
	sort.Sort((geneStartSorter)(idx.byStart))
	sort.Sort((geneEndSorter)(idx.byEnd))
}

// Sorters
type geneStartSorter []*gene
func (g geneStartSorter) Len() int {
	return len(g)
}
func (g geneStartSorter) Less(i, j int) bool {
	return g[i].start < g[j].start
}
func (g geneStartSorter) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

type geneEndSorter []*gene
func (g geneEndSorter) Len() int {
	return len(g)
}
func (g geneEndSorter) Less(i, j int) bool {
	return g[i].end < g[j].end
}
func (g geneEndSorter) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}


// ***** BED TYPE *************************************************************

type bed struct {
	chr string
	start int
	end int
}

func parseBed(line string) (*bed, error) {
	// Split
	fields := strings.Split(line, "\t")
	if len(fields) < 3 {
		return nil, fmt.Errorf("Bad number of fields: %d, expected" +
				" at least 3", len(fields))
	}
	
	result := &bed{}
	
	var err error
	result.chr = fields[0]
	result.start, err = strconv.Atoi(fields[1])
	if err != nil { return nil, err }
	result.end, err = strconv.Atoi(fields[2])
	if err != nil { return nil, err }
	
	return result, nil
}


// ***** NEAREST GENE EXTRACTION **********************************************

// Returns the n nearest genes (if exist).
func (idx index) nearestGenes(tile *bed, n int) (genes []*gene,
		distances []int) {
	chrom, ok := idx[tile.chr];
	
	if !ok {
		return nil, nil
	} else {
		return chrom.nearestGenes(tile, n)
	}
}

// Returns the n nearest genes (if exist).
func (idx *chromIndex) nearestGenes(tile *bed, n int) (genes []*gene,
		distances []int) {
	// Set pointers - one going up and one going down
	nGenes := len(idx.byStart)
	
	up := sort.Search(nGenes, func(i int) bool {
		return idx.byStart[i].start > tile.end
	})
	
	down := sort.Search(nGenes, func(i int) bool {
		return idx.byEnd[i].end >= tile.start
	}) - 1
	
	// Set up to be the lowest gene that overlaps the tile, if exists
	for up >= 1 && idx.byStart[up-1].end >= tile.start {
		up--
	}
	
	// Look up genes
	for i := 0; i < n; i++ {
		// No genes left - return
		if down < 0 && up >= nGenes {
			break
		}
		
		// Only up
		if down < 0 {
			genes = append(genes, idx.byStart[up])
			distances = append(distances, distance(tile, idx.byStart[up]))
			up++
		
		// Down only
		} else if up >= nGenes {
			genes = append(genes, idx.byEnd[down])
			distances = append(distances, -distance(tile, idx.byEnd[down]))
			down--
		
		// Both
		} else {
			distUp := distance(tile, idx.byStart[up])
			distDown := distance(tile, idx.byEnd[down])
			
			if distUp < distDown {
				genes = append(genes, idx.byStart[up])
				distances = append(distances, distance(tile, idx.byStart[up]))
				up++
			} else {
				genes = append(genes, idx.byEnd[down])
				distances = append(distances, -distance(tile, idx.byEnd[down]))
				down--
			}
		}
	}
	
	return
}

// Distance between gene and tile. 0 if overlapping. Always positive.
func distance(b *bed, g *gene) int {
	if b.start > g.end {
		return b.start - g.end
	} else if g.start > b.end {
		return g.start - b.end
	} else {
		return 0
	}
}


// ***** BED FILE HANDLING ****************************************************

// Reads the given input bed file and spills the modified into the output
// file. Returns an error if encountered.
func attachGenes(in string, out string, idx index, n int) error {
	// Open input file
	fin, err := os.Open(in)
	if err != nil { return err }
	defer fin.Close()
	scanner := bufio.NewScanner(fin)
	
	// Open output file
	fout, err := os.Create(out)
	if err != nil { return err }
	defer fout.Close()
	bout := bufio.NewWriter(fout)
	defer bout.Flush()
	
	// Add header
	scanner.Scan()
	fmt.Fprint(bout, scanner.Text())
	for i := 1; i <= n; i++ {
		fmt.Fprintf(bout, "\tgene_%d\tdistance_%d", i, i)
	}
	fmt.Fprint(bout, "\n")
	
	// Iterate over lines
	for scanner.Scan() {
		// Parse bed tile
		tile, err := parseBed(scanner.Text())
		if err != nil { return err }
		
		// Scan genes
		genes, dists := idx.nearestGenes(tile, n)
		
		// Print
		fmt.Fprint(bout, scanner.Text())
		for i := 0; i < n; i++ {
			if i < len(genes) {
				fmt.Fprintf(bout, "\t%s\t%d", genes[i].name, dists[i])
			} else {
				fmt.Fprint(bout, "\t\t")
			}
		}
		
		fmt.Fprint(bout, "\n")
	}
	
	return nil
}

const usage =
`Attaches nearest genes to a bed file.

Written by Amit Lavon (amitlavon1@gmail.com).

Usage:
neargene [options]

Accepted options:
	-g <path>
	-genes <path>
		Backround gene file. Should include 4 columns - chromosome, start,
		end and gene name. First line should be a header.

	-i <path>
	-in <path>
		Input bed file. Should include at least 3 columns - chromosome, start
		and end. First line should be a header.

	-o <path>
	-out <path>
		Output file. Will be identical to the input, but each line will have
		its nearest genes attached.

	-n <int>
		Maximal number of genes to report. Default: 1.
`






