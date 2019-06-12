package bedgraph

// Bed file indexing.

import (
	"sort"
	"fmt"
)


// ----- EVENT TYPE ------------------------------------------------------------

// Raw event for indexing. Each row in a bed file creates 2 events - one for
// start and one for end.
type event struct {
	pos   int      // Position along chromosome.
	value float64  // Value of the event (4'th column).
	start bool     // True is event is starting, of false if not.
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


// ----- INDEX -----------------------------------------------------------------

// A single tile in the index.
type tile struct {
	pos   int                  // Start position (0-based).
	value float64              // Value for bed-graph.
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
	
	if len(ichr) == 0 || ichr[len(ichr) - 1].value != t.value {
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
		if from < 0 { from = 0 }
		
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

// A string representation, for debugging.
func (idx Index) str() string {
	result := ""
	for chr := range idx {
		result += chr + "\n"
		for _, t := range idx[chr] {
			result += fmt.Sprintf("\t%d\t%f\n", t.pos, t.value)
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
func (b IndexBuilder) Add(chr string, start, end int, value float64) {
	b[chr] = append(b[chr], &event{start, value, true},
			&event{end, value, false})
}

// Builds an index out of the builder, using the given number of threads.
// Builder keeps its state and can be used with more entries, keeping what it
// had before.
func (b IndexBuilder) build(numThreads int) Index {
	result := Index{}
	
	chrChan := make(chan string, numThreads)
	go func() {
		// First build the map, to keep later access thread-safe.
		for chr := range b {
			result[chr] = tiles{}
		}
		for chr := range b {
			chrChan <- chr
		}
		close(chrChan)
	}()

	done := make(chan int, numThreads)
	for th := 0; th < numThreads; th++ {
		go func() {
			for chr := range chrChan {
				bchr := b[chr]
				
				// Sort events.
				sort.Sort(bchr)
				
				// Create tiles.
				value := 0.0
				for i := range bchr {
					// Create new tile if needed.
					if i > 0 && bchr[i].pos != bchr[i-1].pos {
						t := &tile{bchr[i-1].pos, value}
						result.add(chr, t)
					}
		
					// Update value.
					if bchr[i].start {
						value += bchr[i].value
					} else {
						value -= bchr[i].value
					}
				}
				
				// Create tile for last events (doesn't happen in the above loop).
				if len(bchr) > 0 {
					t := &tile{bchr[len(bchr) - 1].pos, value}
					result.add(chr, t)
				}
			}
			done <- 1
		}()
	}
	
	for th := 0; th < numThreads; th++ {
		<-done
	}
	close(done)
	
	return result
}

// Builds an index out of the builder. Builder keeps its state and can be used
// with more entries, keeping what it had before.
func (b IndexBuilder) Build() Index {
	return b.build(1)
}

// Builds an index out of the builder, using the given number of threads.
// Builder keeps its state and can be used with more entries, keeping what it
// had before.
func (b IndexBuilder) BuildThreads(numThreads int) Index {
	return b.build(numThreads)
}




