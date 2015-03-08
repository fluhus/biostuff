package main

import (
	"os"
	"fmt"
	"math"
	"sort"
	"bytes"
	"myflag"
	"os/exec"
	//"io/ioutil"
	"bioformats/bed"
	"bioformats/bed/bedgraph"
)


// ***** MAIN *****************************************************************

func main() {
	// Parse arguments.
	parseArguments()
	if arguments.err != nil {
		fmt.Println("Error parsing arguments:", arguments.err)
		os.Exit(1)
	} else if !myflag.HasAny() {
		fmt.Println(usage)
		os.Exit(1)
	}
	
	var data [][]float64  // Numbers to plot.
	var labels []string   // Labels for plot legend.
	
	// Choose strategy.
	if len(arguments.bedgraphs) == 1 {
		// Compare 1 bedgraph to many beds.
		labels = arguments.beds
		
		// Create index.
		idx, err := makeIndex(arguments.bedgraphs[0])
		if err != nil {
			fmt.Println("Error reading bedgraph:", err)
			os.Exit(2)
		}
		
		// Go over bed files.
		for _, file := range arguments.beds {
			values, err := aggregate(file, idx, arguments.dist)
			if err != nil {
				fmt.Println("Error reading regions file " + file + ":", err)
				os.Exit(2)
			}
			
			data = append(data, values)
		}
		
	} else {
		// Compare 1 bed to many bedgraphs.
		if len(arguments.beds) != 1 {
			panic("Must have either 1 bed or 1 bed-graph.")
		}
	}
	
	// Plot using python
	fmt.Println("Printing image to file...")
	plotWithPython(data, labels, arguments.img)
	
	fmt.Println("Done!")
}


// ***** ARGUMENT PARSING *****************************************************

// Holds parsed arguments.
var arguments struct {
	bedgraphs []string
	beds      []string
	img       string
	dist      int
	bin       int
	err       error
}

// Parses input arguments. arguments.err will hold the parsing error,
// if encountered. Caller should check for myflag.HasAny().
func parseArguments() {
	// Register arguments.
	bedgraphFile := myflag.String("bed-graph", "bg", "path",
			"Bed graph file for 1 bed-graph to many beds plot.", "")
	bedFile := myflag.String("bed", "b", "path",
			"Bed file for 1 bed to many bed-graphs plot.", "")
	img := myflag.String("img", "i", "path",
			"Output image file. If not given, image will be opened.", "")
	dist := myflag.Int("range", "r", "integer",
			"Range around tile center to plot. Default is 5000.", 5000)
	bin := myflag.Int("bin", "", "integer",
			"Size of each bin in the histogram. Default is 1.", 1)
	
	// Parse!
	arguments.err = myflag.Parse()
	if arguments.err != nil { return }
	
	// Check argument validity
	if *bedgraphFile != "" && *bedFile != "" {
		arguments.err = fmt.Errorf("Only one common file may be set." +
				"Please choose either bed or bedgraph.")
		return
	}
	
	if *bedgraphFile == "" && *bedFile == "" {
		arguments.err = fmt.Errorf("No common file was set." +
				"Please choose either bed or bedgraph.")
		return
	}
	
	if len(myflag.Args()) == 0 {
		arguments.err = fmt.Errorf("No query files.")
		return
	}
	
	if *dist < 0 {
		arguments.err = fmt.Errorf("Bad range: %d, should be non-negative.",
				*dist)
		return
	}
	
	if *bin < 0 {
		arguments.err = fmt.Errorf("Bad bin size: %d, should be non-negative.",
				*bin)
		return
	}
	
	// Assign to arguments.
	arguments.dist = *dist
	arguments.bin = *bin
	arguments.img = *img
	
	if *bedFile != "" {
		arguments.beds = []string{ *bedFile }
		arguments.bedgraphs = myflag.Args()
	} else {
		arguments.bedgraphs = []string{ *bedgraphFile }
		arguments.beds = myflag.Args()
	}
}


// ***** BACKGROUND INDEXING **************************************************

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


// ***** AGGREGATION **********************************************************

// Creates an aggregation value slice for the given bed file.
func aggregate(path string, idx index, dist int) ([]float64, error) {
	f, err := os.Open(path)
	if err != nil { return nil, err }
	scanner := bed.NewScanner(f)
	result := make([]float64, dist*2 + 1)
	
	lineNum := 1
	
	for scanner.Scan() {
		lineNum++
		b := scanner.Bed()
		pos := (b.Start + b.End) / 2
		
		idx.collect(b.Chr, pos, result)
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

func plotWithPython(filesData [][]float64, labels []string, outFile string) {
	panic("check outfile for empty!")
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
		src = append(src, ("],linewidth=2, label='" + labels[i] +
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



