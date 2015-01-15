package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"sort"
)

func main() {
	// Check arguments
	if len(os.Args) != 2 {
		fmt.Println("Error: Bad number of arguments.")
		os.Exit(1)
	}
	
	// Load genes
	fmt.Println("Loading genes from:", os.Args[1])
	idx, err := loadGenes(os.Args[1])
	if err != nil {
		fmt.Println("Error loading genes:", err)
		os.Exit(2)
	}
	
	fmt.Println(len(idx))
	
	b := &bed{"chr18", 75579801, 75579900}
	g, d := idx[b.chr].nearestGenes(b, 2)
	for i := range g {
		fmt.Println(g[i].name, g[i].start, g[i].end, d[i])
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
			return nil, fmt.Errorf("Bad number of fields: %d, expected 4.",
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


// ***** NEAREST GENE EXTRACTION **********************************************

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
			
			if distUp > distDown {
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




