package main

// Handles signal file indexing.

import (
	"os"
	"fmt"
	"sort"
	"bioformats/bed/bedgraph"
)

// A single range in the index.
type tile struct {
	start int
	end int
	value float64
}

// Index type. Key is chromosome, value is a sorted list of tiles.
type index map[string][]*tile

// Creates an index from the given background file.
func makeIndex(path string) (index, error) {
	f, err := os.Open(path)
	if err != nil { return nil, err }
	scanner := bedgraph.NewScanner(f)

	result := make(map[string][]*tile)
	
	// Scan bed graph background.
	for scanner.Scan() {
		// Parse line.
		b := scanner.Bed()
		b.End--   // To avoid overlapping tiles.
		
		// Add chromosome.
		if _,ok := result[b.Chr]; !ok {
			result[b.Chr] = nil
		}
		
		result[b.Chr] = append(result[b.Chr], &tile{b.Start, b.End, b.Value})
	}
	
	// Sort index.
	for _,chr := range result {
		sort.Sort(tileSorter(chr))
		
		for i := range chr {
			if i != 0 && chr[i].start <= chr[i-1].end {
				return nil, fmt.Errorf("Overlapping tiles: %v, %v", *chr[i-1],
						*chr[i])
			}
		}
	}
	
	return result, nil
}

// Adds background values around pos to the given value slice.
func (idx index) collect(chr string, pos int, values []float64) {
	chrIndex, ok := idx[chr]
	if !ok { return }
	
	dist := (len(values) - 1) / 2
	
	// Find minimal greater tile
	i := sort.Search(len(chrIndex), func(i int) bool {
		return pos+dist < chrIndex[i].start
	}) - 1
	
	// Collect tiles that overlap with my range
	var tiles []*tile
	for i >= 0 && chrIndex[i].end >= pos-dist {
		tiles = append(tiles, chrIndex[i])
		i--
	}
	
	// Update values
	for _, t := range tiles {
		from := max(pos - dist, t.start)
		to := min(pos + dist, t.end)
		
		for i := from; i <= to; i++ {
			values[i - pos + dist] += t.value
		}
	}
}

// Counts how many tiles there are in the index.
func (idx index) size() int {
	result := 0
	
	for _, chr := range idx {
		result += len(chr)
	}
	
	return result
}

// Functions for sorting tiles.
type tileSorter []*tile
func (s tileSorter) Len() int {return len(s)}
func (s tileSorter) Less(i, j int) bool {return s[i].start < s[j].start}
func (s tileSorter) Swap(i, j int) {s[i], s[j] = s[j], s[i]}


