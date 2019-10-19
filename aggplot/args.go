package main

// Handles argument parsing.

import (
	"flag"
	"fmt"
)

// Holds parsed arguments.
var arguments struct {
	bedgraphs []string // Input bed-graph files.
	beds      []string // Input bed files.
	img       string   // Output image file.
	txt       string   // Output text file.
	dist      int      // Distance around tile center.
	bin       int      // Bin size.
	err       error    // Parsing error.
}

// Parses input arguments. arguments.err will hold the parsing error,
// if encountered. Caller should check for myflag.HasAny().
func parseArguments() {
	// Register arguments.
	bedgraphFile := flag.String("bedgraph", "",
		"Bed-graph file for 1 bed-graph to many beds plot.")
	bedFile := flag.String("bed", "",
		"Bed file for 1 bed to many bed-graphs plot.")
	img := flag.String("img", "",
		"Output image file. "+
			"Give 'show' to show the image without saving it.")
	txt := flag.String("out", "",
		"Output text file. "+
			"If not given, no text output will be generated.")
	dist := flag.Int("range", 5000,
		"Range around tile center to plot.")
	bin := flag.Int("bin", 1,
		"Size of bins on the x-axis. Default is 1.")

	// Check argument validity
	if *bedgraphFile != "" && *bedFile != "" {
		arguments.err = fmt.Errorf("only one common file may be set;" +
			"Please choose either bed or bedgraph")
		return
	}

	if *bedgraphFile == "" && *bedFile == "" {
		arguments.err = fmt.Errorf("no common file was set;" +
			"please choose either bed or bedgraph")
		return
	}

	if len(flag.Args()) == 0 {
		arguments.err = fmt.Errorf("no query files")
		return
	}

	if *dist < 0 {
		arguments.err = fmt.Errorf("bad range: %d, should be non-negative",
			*dist)
		return
	}

	if *bin < 0 {
		arguments.err = fmt.Errorf("bad bin size: %d, should be positive",
			*bin)
		return
	}

	// Assign to arguments.
	arguments.dist = *dist
	arguments.img = *img
	arguments.txt = *txt
	arguments.bin = *bin

	if *bedFile != "" {
		arguments.beds = []string{*bedFile}
		arguments.bedgraphs = flag.Args()
	} else {
		arguments.bedgraphs = []string{*bedgraphFile}
		arguments.beds = flag.Args()
	}
}

// Usage help message.
const usage = `Creates aggregation plots of average signals around tiles.

Written by Amit Lavon (amitlavon1@gmail.com).

Usage:
aggplot [options] -bedgraph <bedgraph> <bed 1> <bed 2> <bed 3>...
or
aggplot [options] -bed <bed> <bedgraph 1> <bedgraph 2> <bedgraph 3>...

Choose either 1 bed-graph to many beds using '-bedgraph', or 1 bed to many
bedgraphs using '-bed'.

Options:
`
