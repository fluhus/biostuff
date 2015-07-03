package bed

// Bed file indexing.

import (
	"sort"
	"strconv"
	"fmt"
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
		if counts[evt] < 1 {
			panic("Got non-positive count at '" + evt + "'")
		}
		
		// Add to set.
		set[evt] = struct{}{}
		
		// Add value.
		v, err := strconv.ParseFloat(evt, 64)
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
	names map[string]struct{}  // Set of overlapping event names.
}

// Compares a tile's values to another's. Returns true iff all fields except pos
// have equal values. Deep-checks event sets.
func (t *tile) valuesEqual(t2 *tile) bool {
	if t.value != t2.value { return false }
	if len(t.names) != len(t2.names) { return false }

	for e := range t.names {
		if _, ok := t2.names[e]; !ok { return false }
	}
	
	return true
}

// A slice of tiles, duh.
type tiles []*tile

// A bed index. Used to retrieve names of overlapping regions (genes, exons...)
// and values from bed-graph files.
//
// To create an index, use the IndexBuilder type.
type Index map[string]tiles

// Appends a tile at the given chromosome name, only if it has different values
// from the last tile in that chromosome. Tile position must be greater than
// last tile's position.
func (idx Index) add(chr string, t *tile) {
	ichr := idx[chr]
	if len(ichr) > 0 && ichr[len(ichr) - 1].pos >= t.pos {
		panic("Input tile position must be greater than last tile's.")
	}
	
	if len(ichr) == 0 || !ichr[len(ichr) - 1].valuesEqual(t) {
		idx[chr] = append(idx[chr], t)
	}
}

// Returns the value at the given position. Returns 0 if no value is registered.
func (idx Index) Value(chr string, pos int) float64 {
	ichr := idx[chr]
	
	// If no data, return empty.
	if len(ichr) == 0 {
		return 0
	}
	
	// Search for containing tile.
	i := sort.Search(len(ichr), func(j int) bool {
		return ichr[j].pos > pos
	}) - 1
	
	// Not found.
	if i == -1 {
		return 0
	}
	
	return ichr[i].value
}

func (idx Index) ValueRange(chr string, start, end int) []float64 {
	ichr := idx[chr]
	result := make([]float64, end - start)
	
	// If no data, return zeros.
	if len(ichr) == 0 {
		return result
	}
	
	// Search for containing tiles.
	i := sort.Search(len(ichr), func(j int) bool {
		return ichr[j].pos > start
	}) - 1
	
	if i == -1 { i = 0 }
	
	// Go over 
	for i < len(ichr) && ichr[i].pos < end {
		from := ichr[i].pos - start
		to := len(result)
		if i < len(ichr) - 1 {
			to2 := ichr[i + 1].pos - start
			if to2 < to { to = to2 }
		}
		
		for j := from; j < to; j++ {
			result[j] = ichr[i].value
		}

		i++
	}
	
	return result
}

// Returns a set of overlapping names at the given position. Modifying the set
// does not affect the index. Always returns non-nil.
func (idx Index) Names(chr string, pos int) map[string]struct{} {
	result := map[string]struct{}{}
	ichr := idx[chr]
	
	// If no data, return empty.
	if len(ichr) == 0 {
		return result
	}
	
	// Search for containing tile.
	i := sort.Search(len(ichr), func(j int) bool {
		return ichr[j].pos > pos
	}) - 1
	
	// Not found.
	if i == -1 {
		return result
	}
	
	for name := range ichr[i].names {
		result[name] = struct{}{}
	}
	
	return result
}

// Returns the name at the given position. If several overlap, returns one
// arbitrarily. If non found, returns an empty string.
func (idx Index) Name(chr string, pos int) string {
	for name := range idx.Names(chr, pos) {
		return name
	}
	return ""
}

// A string representation, for debugging.
func (idx Index) str() string {
	result := ""
	for chr := range idx {
		result += chr + "\n"
		for _, t := range idx[chr] {
			result += fmt.Sprintf("\t%d\t%f\t[", t.pos, t.value)
			for name := range t.names {
				result += name + ", "
			}
			result += "]\n"
		}
	}
	return result
}


// ----- INDEX BUILDER ---------------------------------------------------------

// Creates indexes from given bed entries.
type IndexBuilder map[string]events  // Maps chromosome to list of events.

// Returns a new index builder.
func NewIndexBuilder() IndexBuilder {
	return IndexBuilder{}
}

// Adds a bed entry to the builder.
func (b IndexBuilder) Add(chr string, start, end int, name string) {
	b[chr] = append(b[chr], &event{start, name, true},
			&event{end, name, false})
}

// Builds an index out of the builder. Builder keeps its state and can be used
// with more entries, keeping what it had before.
func (b IndexBuilder) Build() Index {
	result := Index{}

	for chr, bchr := range b {
		// Sort events.
		sort.Sort(bchr)
		
		// Create tiles.
		result[chr] = tiles{}
		counts := eventCounter{}
		for i := range bchr {
			// Create new tile if needed.
			if i > 0 && bchr[i].pos != bchr[i-1].pos {
				set, val := eventSet(counts)
				t := &tile{bchr[i-1].pos, val, set}
				result.add(chr, t)
			}
			
			// Update counter.
			if bchr[i].start {
				counts.inc(bchr[i].name)
			} else {
				counts.dec(bchr[i].name)
			}
		}
		
		// Create tile for last events (doesn't happen in the above loop).
		if len(bchr) > 0 {
			set, val := eventSet(counts)
			t := &tile{bchr[len(bchr) - 1].pos, val, set}
			result.add(chr, t)
		}
	}
	
	return result
}






