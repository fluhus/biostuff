package bed

// Bed file indexing.

import (
	"fmt"
	"sort"
)

// ----- EVENT TYPE ------------------------------------------------------------

// Raw event for indexing. Each row in a bed file creates 2 events - one for
// start and one for end.
type event struct {
	pos   int    // Position along chromosome.
	name  string // Name of the event (4'th column).
	start bool   // True is event is starting, of false if not.
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

// Makes an event set from the given count-map. All counts are assumed to be
// at least 1. Panics if not.
func eventSet(counts eventCounter) map[string]struct{} {
	set := map[string]struct{}{}

	for evt := range counts {
		if counts[evt] < 1 {
			panic("Got non-positive count at '" + evt + "'")
		}

		// Add to set.
		set[evt] = struct{}{}
	}

	return set
}

// ----- INDEX -----------------------------------------------------------------

// A single tile in the index.
type tile struct {
	pos   int                 // Start position (0-based).
	names map[string]struct{} // Set of overlapping event names.
}

// Compares a tile's values to another's. Returns true iff all fields except pos
// have equal values. Deep-checks event sets.
func (t *tile) valuesEqual(t2 *tile) bool {
	if len(t.names) != len(t2.names) {
		return false
	}

	for e := range t.names {
		if _, ok := t2.names[e]; !ok {
			return false
		}
	}

	return true
}

// A slice of tiles, duh.
type tiles []*tile

// A bed index. Used to retrieve names of overlapping regions
// (genes, exons...).
//
// To create an index, use the IndexBuilder type.
type Index map[string]tiles

// Appends a tile at the given chromosome name, only if it has different values
// from the last tile in that chromosome. Tile position must be greater than
// last tile's position.
func (idx Index) add(chr string, t *tile) {
	ichr := idx[chr]
	if len(ichr) > 0 && ichr[len(ichr)-1].pos >= t.pos {
		panic("Input tile position must be greater than last tile's.")
	}

	if len(ichr) == 0 || !ichr[len(ichr)-1].valuesEqual(t) {
		idx[chr] = append(idx[chr], t)
	}
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
			result += fmt.Sprintf("\t%d\t[", t.pos)
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
type IndexBuilder map[string]events // Maps chromosome to list of events.

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
// again with more entries, keeping what it had before.
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
				set := eventSet(counts)
				t := &tile{bchr[i-1].pos, set}
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
			set := eventSet(counts)
			t := &tile{bchr[len(bchr)-1].pos, set}
			result.add(chr, t)
		}
	}

	return result
}
