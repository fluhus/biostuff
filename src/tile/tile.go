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
	for _, file := range args.inFiles {
		fmt.Fprintf(os.Stderr, "\t%s\n", file)
		
		t, err := tileFile(file, args.tileSize)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
		files = append(files, t)
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
func tileFile(file string, tileSize int) (*fileTiles, error) {
	f, err := os.Open(file)
	if err != nil { return nil, err }
	defer f.Close()
	
	out := make(map[string]tiles)
	
	scanner := bufio.NewScanner(f)
	scanner.Scan() // Skip header line.
	
	for scanner.Scan() {
		// Split to fields.
		fields := strings.Split(scanner.Text(), "\t")
		if len(fields) != 12 {
			return nil, fmt.Errorf("Bad number of fields: %d", len(fields))
		}
		
		// Extract numbers from line.
		chr := fields[0]
		pos, err := strconv.Atoi(fields[1])
		if err != nil { return nil, err }
		total, err := strconv.ParseFloat(fields[5], 64)
		if err != nil { return nil, err }
		if total == 0 { continue }  // Avoid parsing 'NA'.
		ratio, err := strconv.ParseFloat(fields[4], 64)
		if err != nil { return nil, err }
		
		methd := int( total * ratio )
		
		// Create chromosome.
		if out[chr] == nil {
			out[chr] = make(map[int]*tile)
		}
		
		// Round position to tile.
		pos = pos / tileSize * tileSize
		
		// Create tile.
		if out[chr][pos] == nil {
			out[chr][pos] = &tile{}
		}
		
		out[chr][pos].total += int(total)
		out[chr][pos].methd += methd
	}
	
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	
	return &fileTiles{file, out}, nil
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
	inFiles []string
	outFile string
	tileSize int
	err error
}

func parseArguments() {
	out := myflag.String("out", "o", "path", "Output file. Default is" +
			" standard output.", "stdout")
	size := myflag.Int("size", "s", "integer", "Length of tile. Default is "+
			"100.", 100)
	
	args.err = myflag.Parse()
	if args.err != nil { return }
	
	args.inFiles = myflag.Args()
	if len(args.inFiles) == 0 {
		args.err = fmt.Errorf("No input files given.")
		return
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

Accepted options:
`







