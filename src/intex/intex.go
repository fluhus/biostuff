package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
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
	fmt.Println("Loading genes...")
	genes, err := loadGenes(arguments.genesFile)
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
	
	// Attach region types to input file
	fmt.Println("Attaching region types...")
	err = attachRegions(arguments.inFile, arguments.outFile, idx)
	if err != nil {
		fmt.Println("Error attaching region types:", err)
		os.Exit(2)
	}
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
	if err != nil { return nil, fmt.Errorf("Bad gene start position: %s", err.Error()) }
	result.end, err = strconv.Atoi(fields[2])
	if err != nil { return nil, fmt.Errorf("Bad gene end position: %s", err.Error()) }
	
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

func (g genomicRegion) String() string {
	switch g {
	case intergenic: return "intergenic"
	case intron: return "intron"
	case exon: return "exon"
	default: panic("Unexpected region type")
	}
}

func maxRegion(r1, r2 genomicRegion) genomicRegion {
	if r1 > r2 {
		return r1
	} else {
		return r2
	}
}

// Fetches the highest region type that overlaps the given region.
func (idx eventIndex) regionType(b *bed) genomicRegion {
	chromEvents := idx[b.chr]
	
	// Find event at start
	i := sort.Search(len(chromEvents), func(i int) bool {
		return chromEvents[i].pos > b.start
	}) - 1
	
	// Can reach -1 if search returns 0
	if i == -1 {
		i = 0
	}
	
	result := intergenic
	
	for chromEvents[i].pos <= b.end {
		switch chromEvents[i].typ {
		case geneStart, exonEnd:
			result = maxRegion(result, intron)
		case geneEnd:
			result = maxRegion(result, intergenic)
		case exonStart:
			result = maxRegion(result, exon)
		}
		
		i++
	}
	
	return result
}


// ***** BED FILE INPUT *******************************************************

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

func attachRegions(in string, out string, idx eventIndex) error {
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
	fmt.Fprintf(bout, "%s\tregion_type\n", scanner.Text())
	
	// Iterate over lines
	for scanner.Scan() {
		// Parse bed tile
		tile, err := parseBed(scanner.Text())
		if err != nil { return err }
		
		// Find region type
		regType := idx.regionType(tile)
		
		// Print
		fmt.Fprintf(bout, "%s\t%v\n", scanner.Text(), regType)
	}
	
	return nil
}


// ***** ARGUMENTS ************************************************************

var arguments struct {
	genesFile string
	inFile string
	outFile string
	err error
}

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
	
	if len(flags.Args()) > 0 {
		arguments.err = fmt.Errorf("Unknown argument: %s", flag.Args()[0])
		return
	}
}

const usage =
`Attaches genomic region type (intro/exon/intergenic) to a bed file.

Written by Amit Lavon (amitlavon1@gmail.com).

Usage:
intex [options]

Accepted options:
	-g <path>
	-genes <path>
		Backround gene file. Should include at least 5 columns - chromosome,
		start, end, exon starts, exon ends. First line should be a header.
		Exon starts and ends should be comma separated.

	-i <path>
	-in <path>
		Input bed file. Should include at least 3 columns - chromosome, start
		and end. First line should be a header.

	-o <path>
	-out <path>
		Output file. Will be identical to the input, but each line will have
		its region type attached.
`




