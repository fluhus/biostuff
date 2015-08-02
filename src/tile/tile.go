package main

import (
	"os"
	"fmt"
	"math"
	"sort"
	"bufio"
	"myflag"
	"strconv"
	"strings"
)

func main() {
	// Handle arguments.
	parseArguments()
	if !myflag.HasAny() {
		fmt.Fprint(os.Stderr, usage)
		fmt.Fprint(os.Stderr, myflag.HelpString())
		os.Exit(1)
	} else if args.err != nil {
		fmt.Fprintln(os.Stderr, "Argument error:", args.err)
		os.Exit(1)
	}
	
	// Read input files.
	fmt.Fprintln(os.Stderr, "Reading input files...")
	var files []*fileTiles
	for i, group := range args.inFiles {
		fmt.Fprintf(os.Stderr, "Reading %s...\n", args.labels[i])
		
		current := &fileTiles{ args.labels[i], make(map[string]tiles) }
		files = append(files, current)
		
		for _, file := range group {
			fmt.Fprintf(os.Stderr, "\t%s\n", file)
		
			err := tileFile(file, args.tileSize, current)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}
		}
	}
	
	// Print output.
	fmt.Fprintln(os.Stderr, "Printing...")
	err := printTiles(args.outFile, files, args.tileSize)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}


// ***** TILING ***************************************************************

type tile struct {
	total int
	methd int
}

// Maps from position to tile.
type tiles map[int]*tile

// Maps from chromosome to tiles.
type fileTiles struct {
	name string
	tiles map[string]tiles
}

// Reads tiles from the given file.
func tileFile(file string, tileSize int, t *fileTiles) error {
	f, err := os.Open(file)
	if err != nil { return err }
	defer f.Close()
	
	scanner := bufio.NewScanner(f)
	scanner.Scan() // Skip header line.
	
	for scanner.Scan() {
		// Split to fields.
		fields := strings.Split(scanner.Text(), "\t")
		if len(fields) != 12 {
			return fmt.Errorf("Bad number of fields: %d", len(fields))
		}
		
		// Extract numbers from line.
		chr := fields[0]
		pos, err := strconv.Atoi(fields[1])
		if err != nil { return err }
		total, err := strconv.ParseFloat(fields[5], 64)
		if err != nil { return err }
		if total == 0 { continue }  // Avoid parsing 'NA'.
		ratio, err := strconv.ParseFloat(fields[4], 64)
		if err != nil { return err }
		
		methd := int( total * ratio )
		
		// Create chromosome.
		if t.tiles[chr] == nil {
			t.tiles[chr] = make(map[int]*tile)
		}
		
		// Round position to tile.
		pos = pos / tileSize * tileSize
		
		// Create tile.
		if t.tiles[chr][pos] == nil {
			t.tiles[chr][pos] = &tile{}
		}
		
		t.tiles[chr][pos].total += int(total)
		t.tiles[chr][pos].methd += methd
	}
	
	if scanner.Err() != nil {
		return scanner.Err()
	}
	
	return nil
}


// ***** OUTPUT ***************************************************************

func printTiles(outFile string, t []*fileTiles, tileSize int) error {
	var bout *bufio.Writer
	
	if outFile == "stdout" {
		bout = bufio.NewWriter(os.Stdout)
	} else {
		f, err := os.Create(outFile)
		if err != nil { return err }
		defer f.Close()
		bout = bufio.NewWriter(f)
	}
	
	defer bout.Flush()
	
	// Print header.
	fmt.Fprintf(bout, "chr\tstart\tend")
	for _, file := range t {
		fmt.Fprintf(bout, "\t%s", file.name)
	}
	fmt.Fprintf(bout, "\n")
	
	// Go over tiles.
	for _, chr := range collectChroms(t) {
		for _, pos := range collectPoss(t, chr) {
			// Print chromosome and position.
			fmt.Fprintf(bout, "%s\t%d\t%d", chr, pos, pos + tileSize)
		
			for _, file := range t {
				var value float64
				
				// Tile exists.
				if file.tiles[chr] != nil && file.tiles[chr][pos] != nil {
					currentTile := file.tiles[chr][pos]
					if currentTile.total == 0 {
						value = 0
					} else {
						value = float64(currentTile.methd) /
								float64(currentTile.total)
					}
				} else {
					value = math.NaN()
				}
				
				fmt.Fprintf(bout, "\t%f", value)
			}
			
			fmt.Fprintln(bout)
		}
	}
	
	return nil
}

// Returns a sorted list of all chromosomes that appear in the given tiles.
func collectChroms(t []*fileTiles) []string {
	chrmap := make(map[string]struct{})
	for _, tt := range t {
		for chr := range tt.tiles {
			chrmap[chr] = struct{}{}
		}
	}
	
	chrs := make([]string, 0, len(chrmap))
	for chr := range chrmap {
		chrs = append(chrs, chr)
	}
	
	sort.Sort(sort.StringSlice(chrs))
	
	return chrs
}

// Returns a sorted list of all positions that appear in the given tiles for
// the given chromosome.
func collectPoss(t []*fileTiles, chr string) []int {
	posmap := make(map[int]struct{})
	for _, tt := range t {
		if tt.tiles[chr] != nil {
			for pos := range tt.tiles[chr] {
				posmap[pos] = struct{}{}
			}
		}
	}
	
	poss := make([]int, 0, len(posmap))
	for pos := range posmap {
		poss = append(poss, pos)
	}
	
	sort.Sort(sort.IntSlice(poss))
	
	return poss
}


// ***** ARGUMENTS ************************************************************

var args struct {
	inFiles [][]string
	labels []string
	outFile string
	tileSize int
	err error
}

func parseArguments() {
	out := myflag.String("out", "o", "path", "Output file. Default is" +
			" standard output.", "stdout")
	labels := myflag.String("labels", "L", "", "Comma-separated labels of " +
			"columns. Default is file-names.", "")
	size := myflag.Int("size", "s", "integer", "Length of tile. Default is "+
			"100.", 100)
	
	args.err = myflag.Parse()
	if args.err != nil { return }
	
	// Split input files into groups.
	for _, list := range myflag.Args() {
		split := strings.Split(list, ",")
		for _, file := range split {
			if file == "" {
				args.err = fmt.Errorf("Empty file name in: %s", list)
				return
			}
		}
		args.inFiles = append(args.inFiles, split)
	}
	
	if len(args.inFiles) == 0 {
		args.err = fmt.Errorf("No input files given.")
		return
	}
	
	// Handle labels.
	if *labels == "" {
		// Create default labels.
		for _, group := range args.inFiles {
			label := ""
			for _, file := range group {
				if label == "" {
					label = file
				} else {
					label += "," + file
				}
			}
			
			args.labels = append(args.labels, label)
		}
	} else {
		// Use user-defined labels.
		split := strings.Split(*labels, ",")
		if len(split) != len(args.inFiles) {
			args.err = fmt.Errorf("Expected %d labels, but got %d.",
					len(args.inFiles), len(split))
			return
		}
		for _, label := range split {
			if label == "" {
				args.err = fmt.Errorf("Empty labels are not allowed.")
				return
			}
		}
		
		args.labels = split
	}
	
	args.tileSize = *size
	if args.tileSize < 1 {
		args.err = fmt.Errorf("Bad tile size: %d", args.tileSize)
		return
	}
	
	args.outFile = *out
}

var usage =
`Aggregates input meth files into tiles.

Written by Amit Lavon (amitlavon1@gmail.com).

Usage:
tiles [options] <meth 1> <meth 2> <meth 3> ...

Samples can be united to the same column using commas:
tiles [options] <meth 1>,<meth 2> <meth 3>,<meth 4> ...

Accepted options:
`







