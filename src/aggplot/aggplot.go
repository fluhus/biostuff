package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"sort"
	"os/exec"
	"bytes"
	"flag"
	"io/ioutil"
	"math"
)


// ***** MAIN *****************************************************************

func main() {
	// Parse arguments
	if len(os.Args) == 1 {
		fmt.Println(usage)
		os.Exit(0)
	}
	
	parseArguments()
	if arguments.err != nil {
		fmt.Println("Error parsing arguments:", arguments.err)
		os.Exit(1)
	}
	
	// Index signals
	fmt.Println("Reading background...")
	idx, err := makeIndex(arguments.bgFile)
	if err != nil {
		fmt.Println("Error reading background:", err)
		os.Exit(2)
	}
	
	// Aggregate bed results
	fmt.Println("Reading region data...")
	var filesData [][]float64
	
	for _,file := range arguments.regionFiles {
		values, err := aggregate(file, idx, arguments.dist)
		if err != nil {
			fmt.Println("Error reading regions file " + file + ":", err)
			os.Exit(2)
		}
		filesData = append(filesData, values)
	}
	
	// Plot using python
	fmt.Println("Printing image to file...")
	plotWithPython(filesData, arguments.imgFile)
	
	fmt.Println("Done!")
}


// ***** ARGUMENT PARSING *****************************************************

// Holds parsed arguments
var arguments struct {
	bgFile string
	imgFile string
	regionFiles []string
	dist int
	err error
}

// Parses input arguments. arguments.err will hold the parsing error,
// if encountered.
func parseArguments() {
	// Create flag set
	flags := flag.NewFlagSet("", flag.ContinueOnError)
	flags.SetOutput(ioutil.Discard)
	
	// Register arguments
	flags.StringVar(&arguments.bgFile, "bg", "", "")
	flags.StringVar(&arguments.bgFile, "b", "", "")
	flags.StringVar(&arguments.imgFile, "img", "", "")
	flags.StringVar(&arguments.imgFile, "i", "", "")
	flags.IntVar(&arguments.dist, "range", 5000, "")
	flags.IntVar(&arguments.dist, "r", 5000, "")
	
	// Parse!
	arguments.err = flags.Parse(os.Args[1:])
	if arguments.err != nil { return }
	arguments.regionFiles = flags.Args()
	
	// Check argument validity
	if arguments.bgFile == "" {
		arguments.err = fmt.Errorf("No background file selected")
		return
	}
	
	if arguments.imgFile == "" {
		arguments.err = fmt.Errorf("No output image file selected")
		return
	}
	
	if len(arguments.regionFiles) == 0 {
		arguments.err = fmt.Errorf("No region files selected")
		return
	}
	
	if arguments.dist < 0 {
		arguments.err = fmt.Errorf("Bad range: %d, should be non-negative",
				arguments.dist)
		return
	}
}


// ***** BACKGROUND INDEXING **************************************************

// A single range in the index.
type tile struct {
	start int
	end int
	value float64
}

// Index type.
// Key is chromosome, value is a sorted list of tiles.
type index map[string][]*tile

// Creates an index from the given background file.
func makeIndex(path string) (index, error) {
	f, err := os.Open(path)
	if err != nil { return nil, err }
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	result := make(map[string][]*tile)
	
	for scanner.Scan() {
		
		// Parse line
		b, err := parseBedGraphLine(scanner.Text(), 0)
		if err != nil { return nil, err }
		b.end--   // to avoid overlapping tiles
		
		// Add chromosome
		if _,ok := result[b.chr]; !ok {
			result[b.chr] = nil
		}
		
		result[b.chr] = append(result[b.chr], &tile{b.start, b.end, b.value})
	}
	
	for _,chr := range result {
		sort.Sort(tileSorter(chr))
		
		for i := range chr {
			if i != 0 && chr[i].start <= chr[i-1].end {
				return nil, fmt.Errorf("Overlapping tiles: %v, %v", *chr[i-1], *chr[i])
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
	for _,t := range tiles {
		from := max(pos - dist, t.start)
		to := min(pos + dist, t.end)
		
		for i := from; i <= to; i++ {
			values[i - pos + dist] += t.value
		}
	}
}

// Functions for sorting tiles.
type tileSorter []*tile
func (s tileSorter) Len() int {return len(s)}
func (s tileSorter) Less(i, j int) bool {return s[i].start < s[j].start}
func (s tileSorter) Swap(i, j int) {s[i], s[j] = s[j], s[i]}


// ***** BED PARSING **********************************************************

// A single bed line.
type bed struct {
	chr string
	start int
	end int
}

// Creates a bed object from a bed line.
func parseBedLine(line string, offset int) (*bed, error) {
	fields := strings.Split(line, "\t")
	if len(fields) < offset + 3 {
		return nil, fmt.Errorf("Bad number of fields in bed: %d, expected %d",
				len(fields), offset + 3)
	}
	
	result := &bed{}
	
	result.chr = fields[0 + offset]
	
	var err error
	result.start, err = strconv.Atoi(fields[1 + offset])
	if err != nil { return nil, err }
	
	result.end, err = strconv.Atoi(fields[2 + offset])
	if err != nil { return nil, err }
	
	return result, nil
}

// A single bed-graph line.
type bedGraph struct {
	chr string
	start int
	end int
	value float64
}

// Creates a bed-graph object from a bed line.
func parseBedGraphLine(line string, offset int) (*bedGraph, error) {
	fields := strings.Split(line, "\t")
	if len(fields) < offset + 4 {
		return nil, fmt.Errorf("Bad number of fields in bed: %d, expected %d",
				len(fields), offset + 4)
	}
	
	result := &bedGraph{}
	
	result.chr = fields[0 + offset]
	
	var err error
	result.start, err = strconv.Atoi(fields[1 + offset])
	if err != nil { return nil, err }
	
	result.end, err = strconv.Atoi(fields[2 + offset])
	if err != nil { return nil, err }
	
	result.value, err = strconv.ParseFloat(fields[3 + offset], 64)
	if err != nil { return nil, err }
	
	return result, nil
}


// ***** AGGREGATION **********************************************************

// Creates an aggregation value slice for the given bed file.
func aggregate(path string, idx index, dist int) ([]float64, error) {
	f, err := os.Open(path)
	if err != nil { return nil, err }
	scanner := bufio.NewScanner(f)
	result := make([]float64, dist*2 + 1)
	
	scanner.Scan()
	lineNum := 1
	
	for scanner.Scan() {
		lineNum++
		b, err := parseBedLine(scanner.Text(), 0)
		if err != nil { return nil, fmt.Errorf("In line %d: %v", lineNum, err) }
		pos := (b.start + b.end) / 2
		
		idx.collect(b.chr, pos, result)
	}
	
	// Normalize by number of lines (average signal)
	for i := range result {
		result[i] /= float64(lineNum-1)
	}
	
	return result, nil
}


// ***** PYTHON INTERFACE *****************************************************

// Converts a value slice to a python list literal.
func valuesToText(values []float64) []byte {
	result := []byte("[")
	for _,v := range values {
		result = append(result, fmt.Sprintf("%f,", v)...)
	}
	result = append(result, "]"...)
	
	return result
}

func plotWithPython(filesData [][]float64, outFile string) {
	// Create imports
	src := []byte("import matplotlib.pyplot as plt\n")
	
	// Find min and max for axes
	minValue := math.MaxFloat64
	maxValue := -math.MaxFloat64
	for i := range filesData {
		for _,v := range filesData[i] {
			if v < minValue { minValue = v }
			if v > maxValue { maxValue = v }
		}
	}
	
	axesXMin := float64(-arguments.dist)
	axesXMax := float64(arguments.dist)
	axesYMin := minValue - 0.1*(maxValue-minValue)
	axesYMax := maxValue + 0.3*(maxValue-minValue)
	
	// Add x=0 marker
	src = append(src, fmt.Sprintf("plt.plot([0,0],[%f,%f],'--k')\n", axesYMin, axesYMax)...)
	
	// Add plot for each file
	for i,values := range filesData {
		src = append(src, fmt.Sprintf("plt.plot(range(%d,%d),[",
				-arguments.dist, arguments.dist+1)...)
		for _,v := range values {
			src = append(src, fmt.Sprintf("%f,", v)...)
		}
		src = append(src, ("],linewidth=2, label='" + arguments.regionFiles[i] +
				"')\n")...)
	}
	
	// Add figure settings
	src = append(src, fmt.Sprintf("plt.title('Aggregation plot')\n")...)
	src = append(src, fmt.Sprintf("plt.xlabel('Distance from region center')\n")...)
	src = append(src, fmt.Sprintf("plt.ylabel('Average signal')\n")...)
	src = append(src, fmt.Sprintf("plt.axis([%f,%f,%f,%f])\n",
			axesXMin, axesXMax, axesYMin, axesYMax)...)
	src = append(src, fmt.Sprintf("plt.legend(loc='upper right')\n")...)
	
	// Save to file command
	src = append(src, fmt.Sprintf("plt.savefig('%s',dpi=150)", outFile)...)
	
	cmd := exec.Command("python")
	cmd.Stdin = bytes.NewReader(src)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}


// ***** MISC. HELPERS ********************************************************

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

const usage =
`Creates aggregation plots of average signals around tiles.

Written by Amit Lavon (amitlavon1@gmail.com).

Usage:
aggplot [options] <regions file 1> <region file 2> <region file 3>...

Each region file should include a header and at least 3 columns - chromosome,
start and end. These columns should be the first ones.

Options:
	-b <path>
	-bg <path>
		Background signal file. Should be in bed-graph format, with a header.

	-i <path>
	-img <path>
		Output image file.

	-r <integer>
	-range <integer>
		Range around tiles to search. Will affect the width of the plot.
		Default: 5000.
`



