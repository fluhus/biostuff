package main

// Handles signal file indexing.

import (
	"os"
	"runtime"

	"github.com/fluhus/biostuff/bioformats/bed/bedgraph"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// Creates an index from the given background file.
func makeIndex(path string) (bedgraph.Index, error) {
	f, err := os.Open(path)
	if err != nil { return nil, err }
	defer f.Close()
	scanner := bedgraph.NewScanner(f)

	builder := bedgraph.NewIndexBuilder()

	// Scan bed graph background.
	for scanner.Scan() {
		// Parse line.
		b := scanner.Bed()
		builder.Add(b.Chr, b.Start, b.End, b.Value)
	}
		
	return builder.BuildThreads(runtime.NumCPU()), nil
}

// Adds background values around pos to the given value slice.
func collect(idx bedgraph.Index, chr string, pos int, values []float64) {
	vals := idx.ValueRange(chr, pos - len(values)/2, pos + len(values)/2 + 1)
	for i := range values {
		values[i] += vals[i]
	}
}


