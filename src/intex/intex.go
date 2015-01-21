package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"sort"
)

func main() {
	// Load genes
	fmt.Println("Loading genes...")
	genes, err := loadGenes("ens_genes_raw.txt")
	if err != nil {
		fmt.Println("Error loading genes:", err)
		os.Exit(2)
	}
	
	// Create indexes
	fmt.Println("Indexing genes...")
	idx, err := aggregateRawEvents( genesToRawEvents(genes) )
	if err != nil {
		fmt.Println("Error indexing genes:", err)
		os.Exit(2)
	}
	
	fmt.Println(len(idx))
}


// ***** GENE PARSER **********************************************************

type gene struct {
	chr string
	start int
	end int
	exonStarts []int
	exonEnds []int
}

func loadGenes(path string) ([]*gene, error) {
	fin, err := os.Open(path)
	if err != nil { return nil, err }
	defer fin.Close()
	
	scanner := bufio.NewScanner(fin)
	scanner.Scan()  // skip header
	var result []*gene
	
	for scanner.Scan() {
		g, err := parseGene(scanner.Text())
		if err != nil { return nil, err }
		
		result = append(result, g)
	}
	
	return result, nil
}

func parseGene(line string) (*gene, error) {
	// Split to fields
	fields := strings.Split(line, "\t")
	if len(fields) < 5 {
		return nil, fmt.Errorf("Bad number of fields, got %d instead of 5 at least",
				len(fields))
	}
	
	result := &gene{}
	var err error
	
	// Copy strings
	result.chr = fields[0]
	
	// Parse numeric fields
	result.start, err = strconv.Atoi(fields[1])
	if err != nil { return nil, fmt.Errorf("Bad start position: %s", err.Error()) }
	result.end, err = strconv.Atoi(fields[2])
	if err != nil { return nil, fmt.Errorf("Bad end position: %s", err.Error()) }
	
	starts := strings.Split(fields[3], ",")
	starts = starts[: len(starts) - 1]  // last field is empty
	result.exonStarts = make([]int, len(starts))
	for i := range starts {
		result.exonStarts[i], err = strconv.Atoi(starts[i])
		if err != nil { return nil, fmt.Errorf("Bad exon start position: %s",
				err.Error()) }
	}
	
	ends := strings.Split(fields[4], ",")
	ends = ends[: len(ends) - 1]  // last field is empty
	result.exonEnds = make([]int, len(ends))
	for i := range ends {
		result.exonEnds[i], err = strconv.Atoi(ends[i])
		if err != nil { return nil, fmt.Errorf("Bad exon end position: %s",
				err.Error()) }
	}
	
	return result, nil
}


// ***** EVENT TYPE ***********************************************************

type eventType int
const (
	exonEnd eventType = iota
	geneEnd
	geneStart
	exonStart
)

type event struct {
	pos int
	typ eventType
}

// Sorts events by position and type.
type eventSorter []*event

func (s eventSorter) Len() int {
	return len(s)
}

func (s eventSorter) Less(i, j int) bool {
	return s[i].pos < s[j].pos || (s[i].pos == s[j].pos && s[i].typ < s[j].typ)
}

func (s eventSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}


// ***** EVENT INDEXING *******************************************************

type eventIndex map[string]eventSorter

// Converts parsed genes to a map with sorted slices of raw events (keys are
// chromosomes).
func genesToRawEvents(genes []*gene) (events eventIndex) {
	events = make(map[string]eventSorter)

	// Go over genes
	for _, g := range genes {
		// Add gene start and end
		events[g.chr] = append(events[g.chr], &event{g.start, geneStart})
		events[g.chr] = append(events[g.chr], &event{g.end, geneEnd})
		
		// Add exon starts
		for _, start := range g.exonStarts {
			events[g.chr] = append(events[g.chr], &event{start, exonStart})
		}
		
		// Add exon ends
		for _, end := range g.exonEnds {
			events[g.chr] = append(events[g.chr], &event{end, exonEnd})
		}
	}
	
	for _, sorter := range events {
		sort.Sort(sorter)
	}
	
	return
}

// Input events should be sorted.
func aggregateRawEvents(rawEvents eventIndex) (eventIndex, error) {
	result := make(map[string]eventSorter)

	// For each chromosome
	for chr := range rawEvents {
		geneCount := 0
		exonCount := 0
		
		for _, evt := range rawEvents[chr] {
			switch evt.typ {
			
			case geneStart:
				geneCount++
				if geneCount == 1 {
					result[chr] = append(result[chr], &event{evt.pos, evt.typ})
				}
			
			case geneEnd:
				geneCount--
				if geneCount == 0 {
					result[chr] = append(result[chr], &event{evt.pos, evt.typ})
				}
				if geneCount < 0 {
					return nil, fmt.Errorf("Reached a negative number of genes")
				}
				
			case exonStart:
				exonCount++
				if exonCount == 1 {
					result[chr] = append(result[chr], &event{evt.pos, evt.typ})
				}
				
			case exonEnd:
				exonCount--
				if exonCount == 0 {
					result[chr] = append(result[chr], &event{evt.pos, evt.typ})
				}
				if exonCount < 0 {
					return nil, fmt.Errorf("Reached a negative number of exons")
				}
			}
		}
	}
	
	return result, nil
}

// Genomic region annotation type
type genomicRegion int

const (
	intergenic genomicRegion = iota
	intron
	exon
)

// Fetches the highest region type that overlaps the given region.
func (idx eventIndex) regionType(chr string, start int, end int) genomicRegion {
	chromEvents := idx[chr]
	
	// Find event at start
	startEvent := sort.Search(len(chromEvents), func(i int) bool {
		return chromEvents[i].pos > start
	}) - 1
	
	
}




