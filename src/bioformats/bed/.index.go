package bed

// Bed file indexing.

import (
	"sort"
	"strconv"
)


// ----- EVENT TYPE ------------------------------------------------------------

// Raw event for indexing. Each row in a bed file creates 2 events - one for
// start and one for end.
type event struct {
	pos   int     // Position along chromosome.
	name  string  // Name / value of the event (4'th column).
	start bool    // True is event is starting, of false if not.
}

// Sorting interface.
type events []*event

func (a events) Len() int {
	return len(a)
}

func (a events) Less(i, j int) bool {
	// Not on the same position
	if a[i].pos != a[j].pos {
		return a[i].pos < a[j].pos
	}

	// End comes before start.
	return a[i].start == false
}

func (a events) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}


// ----- EVENT SET/COUNTER -----------------------------------------------------

// Counts overlapping counts of events.
type eventCounter map[string]int

// Adds 1 to a name's count.
func (e eventCounter) inc(name string) {
	e[name]++
}

// Removes 1 from a name's count.
func (e eventCounter) dec(name string) {
	e[name]--
	if e[name] == 0 {
		delete(e, name)
	}
}

// Makes an event set and a value from the given count-map. All counts are
// assumed to be at least 1. Panics if not.
func eventSet(counts eventCounter) (map[string]struct{}, float64) {
	set := map[string]struct{}{}
	value := 0.0

	for evt := range counts {
		if counts < 1 {
			panic("Got non-positive count at '" + evt + "'")
		}
		
		// Add to set.
		set[evt] = struct{}{}
		
		// Add value.
		v, err := strconv.ParseFloat(evt)
		if err == nil {
			value += float64(counts[evt]) * v
		}
	}
	
	return set, value
}


// ----- INDEX -----------------------------------------------------------------

// A single tile in the index.
type tile struct {
	pos   int                  // Start position (0-based).
	value float64              // Value for bed-graph.
	evts  map[string]struct{}  // Set of overlapping event names.
}

// A slice of tiles, duh.
type tiles []*tile

// A bed index. Used to retrieve names of overlapping regions (genes, exons...)
// and values from bed-graph files.
//
// To create an index, use the IndexBuilder type.
type Index struct {
	evts map[string]tiles
}


// ----- INDEX BUILDER ---------------------------------------------------------

// Creates indexes from given bed entries.
type IndexBuilder map[string]events  // Maps chromosome to list of events.

// Adds a bed entry to the builder.
func (b IndexBuilder) Add(chr string, start, end int, name string) {
	b[name] = append(b[name], &event{start, name, true},
			&events{end, name, false})
}

func (b IndexBuilder) Build() *Index {
	for chr := range b {
		// Sort events.
		sort.Sort(b[chr])
	}
}






