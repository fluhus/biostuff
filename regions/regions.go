package main

import (
	"fmt"
	"os"
	"sort"
	"bufio"
	"strings"

	"github.com/fluhus/biostuff/myflag"
	"github.com/fluhus/biostuff/bioformats/bed"
)

func main() {
	// Parse arguments
	err := parseArgs()
	if err != nil {
		fmt.Println("Error parsing arguments:", err)
		os.Exit(1)
	}
	if args.help {
		fmt.Println(help)
		fmt.Println(myflag.HelpString())
		os.Exit(1)
	}

	fmt.Println("Reading events...")
	e, err := eventsFromFile(args.eventFile)
	
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(2)
	}
	
	fmt.Println("Indexing...")
	idx := e.index()
	
	fmt.Println("Reading regions...")
	err = processFile(args.inFile, args.outFile, idx, args.prior)
	
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(2)
	}
	
	fmt.Println("Done!")
}


// ***** EVENT PARSING ********************************************************

type event struct {
	chr string
	pos int
	start bool    // true if event start, false if event end
	name string
}

type events []*event

func (e events) names() []string {
	m := make(map[string]struct{})
	for _, evt := range e {
		m[evt.name] = struct{}{}
	}
	
	var result []string
	for name := range m {
		result = append(result, name)
	}
	
	return result
}

func eventsFromFile(file string) (events, error) {
	f, err := os.Open(file)
	if err != nil { return nil, err }
	defer f.Close()
	
	scanner := bed.NewScanner(f)
	var result events
	
	for scanner.Scan() {
		b := scanner.Bed()
		name := scanner.Fields()[0]
		result = append(result, &event{ b.Chr, b.Start, true, name })
		result = append(result, &event{ b.Chr, b.End, false, name })
	}
	
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	
	// Sort
	result.sort()
	
	return result, nil
}


// ***** EVENT SORTING ********************************************************

func (e events) sort() {
	sort.Sort(e)
}

func (e events) Len() int {
	return len(e)
}

func (e events) Less(i, j int) bool {
	if e[i].chr != e[j].chr {
		return e[i].chr < e[j].chr
	}
	
	if e[i].pos != e[j].pos {
		return e[i].pos < e[j].pos
	}
	
	if e[i].start != e[j].start {
		return e[i].start == false  // end comes first
	}
	
	return true  // arbitrary
}

func (e events) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}


// ***** EVENT INDEXING *******************************************************

// Maps from event name to count.
type eventCounter map[string]int

// Holds names of events.
type eventSet map[string]struct{}

// Holds event names that are present starting from pos up to the next set.
type eventSetPos struct {
	chr string
	pos int
	names eventSet
}

type index []*eventSetPos

// Events are assumed to be sorted.
func (e events) index() index {
	var result []*eventSetPos
	var counter eventCounter

	for _, evt := range e {
		// Check if new chromosome
		if len(result) == 0 || result[len(result)-1].chr != evt.chr {
			// Reset counter
			counter = make(map[string]int)
		}
		
		// Start -> increment
		if evt.start {
			counter[evt.name]++
			
			// If created new event
			if counter[evt.name] == 1 {
				result = append(result, &eventSetPos{evt.chr, evt.pos,
						counter.set()})
			}
			
		// End -> decrement
		} else {
			counter[evt.name]--
			
			// If deleted an event
			if counter[evt.name] == 0 {
				result = append(result, &eventSetPos{evt.chr, evt.pos,
						counter.set()})
			}
			
			// If negative, we have a problem
			if counter[evt.name] == -1 {
				panic(fmt.Sprintf("-1 event count at (%s,%d): %s",
						evt.chr, evt.pos, evt.name))
			}
		}
	}
	
	return result
}

func (e eventCounter) set() eventSet {
	result := make(map[string]struct{})

	for name := range e {
		if e[name] > 0 {
			result[name] = struct{}{}
		}
	}
	
	return result
}

// Returns the names of all events that overlap 
func (idx index) search(chr string, start int, end int) eventSet {
	// Search
	i := sort.Search(len(idx), func(a int) bool {
		return idx[a].chr > chr || (idx[a].chr == chr && idx[a].pos > start)
	}) - 1
	
	if i == -1 {
		i = 0
	}
	
	result := make(map[string]struct{})
	
	for ; idx[i].chr < chr || (idx[i].chr == chr && idx[i].pos <= end); i++ {
		if idx[i].chr != chr { continue }
		
		for name := range idx[i].names {
			result[name] = struct{}{}
		}
	}
	
	return result
}

// Returns one arbitrary event from the set. If empty, returns an empty string.
func (e eventSet) event() string {
	for name := range e {
		return name
	}
	return ""
}


// ***** REGION FILE PROCESSING ***********************************************

func processFile(in string, out string, idx index, prior []string) error {
	// Buffered i/o.
	var bout *bufio.Writer
	var scanner *bed.Scanner

	// Open files
	if in == "" {
		scanner = bed.NewScanner(os.Stdin)
	} else {
		fin, err := os.Open(in)
		if err != nil { return err }
		defer fin.Close()
		
		scanner = bed.NewScanner(fin)
	}
	
	if out == "" {
		bout = bufio.NewWriter(os.Stdout)
	} else {
		fout, err := os.Create(out)
		if err != nil { return err }
		defer fout.Close()
		
		bout = bufio.NewWriter(fout)
	}
	
	defer bout.Flush()
	
	// Iterate over lines
	for scanner.Scan() {
		// Look up
		b := scanner.Bed()
		eSet := idx.search(b.Chr, b.Start, b.End)
		var name string
		found := false
		
		// Pick by priority
		for _, p := range prior {
			if _, ok := eSet[p]; ok {
				name = p
				found = true
				break
			}
		}
		
		if !found {
			name = eSet.event()
		}
		
		// Print
		fmt.Fprintf(bout, "%s\t%s\n", scanner.Text(), name)
	}
	
	// If no error, will return nil
	return scanner.Err()
}

// ***** ARGUMENTS *************************************************************

var args struct {
	eventFile string
	inFile    string
	outFile   string
	prior     []string
	extend    int
	help      bool
}

func parseArgs() error {
	// Parse command-line flags.
	events := myflag.String("events", "e", "path", "Input event file. Must " +
			"be set. Structure: chromosome, start, end, name.", "")
	in := myflag.String("in", "i", "path", "Input bed file. Default is " +
			"standard input.", "")
	out := myflag.String("out", "o", "path", "Output bed file. Default is " +
			"standard output.", "")
	extend := myflag.Int("extend", "x", "integer", "Extend each event by n " +
			"bases in each direction. Default is 0.", 0)
	prior := myflag.String("priority", "p", "list", "Optional. Comma-" +
			"separated events for when several overlap. The leftmost will " +
			"be returned. Priorities for exons/introns:" +
			" exon,intron,promoter,cpg_island",
			"")
	
	err := myflag.Parse()
	if err != nil {
		return err
	}
	
	if !myflag.HasAny() {
		args.help = true
		return nil
	}
	
	// Check arguments.
	if *events == "" {
		return fmt.Errorf("Event file not set.")
	}
	
	if len(myflag.Args()) != 0 {
		return fmt.Errorf("Unexpected argument: %s", myflag.Args()[0])
	}
	
	if *extend < 0 {
		return fmt.Errorf("Bad extension value: %d. Must be at least 0.",
				*extend)
	}
	
	if *prior != "" {
		args.prior = strings.Split(*prior, ",")
		for _, word := range args.prior {
			if word == "" {
				return fmt.Errorf("Empty words in priority are not allowed.")
			}
		}
	}
	
	args.eventFile = *events
	args.inFile = *in
	args.outFile = *out
	args.extend = *extend
	
	return nil
}

var help =
`Crosses region files with 'events' such as introns, exons, LINEs etc.

Written by Amit Lavon (amitlavon1@gmail.com).

Usage:
regions [options] -e <event file>

Accepted options:`



