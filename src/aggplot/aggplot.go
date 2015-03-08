package main

import (
	"os"
	"fmt"
	"math"
	"sort"
	"bufio"
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
		fmt.Print(usage + myflag.HelpString())
		os.Exit(1)
	}
	
	var data [][]float64  // Numbers to plot.
	var labels []string   // Labels for plot legend.
	
	// Choose strategy.
	if len(arguments.bedgraphs) == 1 {
		// Compare 1 bedgraph to many beds.
		labels = arguments.beds
		
		// Create index.
		fmt.Printf("Reading bed-graph '%s'...\n", arguments.bedgraphs[0])
		idx, err := makeIndex(arguments.bedgraphs[0])
		if err != nil {
			fmt.Println("Error reading bed-graph:", err)
			os.Exit(2)
		}
		
		// Go over bed files.
		for _, file := range arguments.beds {
			fmt.Printf("Reading bed '%s'...\n", file)
			values, err := aggregate(file, []index{idx}, arguments.dist)
			if err != nil {
				fmt.Println("Error reading bed:", err)
				os.Exit(2)
			}
			
			data = append(data, values[0])
		}
		
	} else {
		// Compare 1 bed to many bedgraphs.
		if len(arguments.beds) != 1 {
			panic("Must have either 1 bed or 1 bed-graph.")
		}
		
		labels = arguments.bedgraphs
		
		// Create indexes.
		var idxs []index
		for _, file := range arguments.bedgraphs {
			fmt.Printf("Reading bed-graph '%s'...\n", file)
			idx, err := makeIndex(file)
			if err != nil {
				fmt.Println("Error reading bed-graph:", err)
				os.Exit(2)
			}
			
			idxs = append(idxs, idx)
		}
		
		// Process bed file.
		fmt.Printf("Reading bed '%s'...\n", arguments.beds[0])
		var err error
		data, err = aggregate(arguments.beds[0], idxs, arguments.dist)
		if err != nil {
			fmt.Println("Error reading bed:", err)
			os.Exit(2)
		}
		
		// Signals from different bed-graphs should be normalized.
		fmt.Println("Normalizing...")
		//normalize(data)
	}
	
	// Generate bins.
	xvals := make([]int, 2 * arguments.dist + 1)
	for i := range xvals {
		xvals[i] = i - arguments.dist
	}
	
	// Output text file.
	if arguments.txt != "" {
		fmt.Println("Printing to text file...")
		printData(data, xvals, labels, arguments.txt)
	}
	
	// Plot using python
	if arguments.img != "" {
		fmt.Println("Generating image...")
		plotWithPython(data, xvals, labels, arguments.img)
	}
	
	fmt.Println("Done!")
}


// ***** ARGUMENT PARSING *****************************************************

// Holds parsed arguments.
var arguments struct {
	bedgraphs []string   // Input bed-graph files.
	beds      []string   // Input bed files.
	img       string     // Output image file.
	txt       string     // Output text file.
	dist      int        // Distance around tile center.
	bin       int        // Bin size.
	err       error      // Parsing error.
}

// Parses input arguments. arguments.err will hold the parsing error,
// if encountered. Caller should check for myflag.HasAny().
func parseArguments() {
	// Register arguments.
	bedgraphFile := myflag.String("bedgraph", "bg", "path",
			"Bed graph file for 1 bed-graph to many beds plot.", "")
	bedFile := myflag.String("bed", "b", "path",
			"Bed file for 1 bed to many bed-graphs plot.", "")
	img := myflag.String("img", "i", "path",
			"Output image file. Give 'show' to show the image without saving " +
			"it.", "")
	txt := myflag.String("text", "t", "path",
			"Output text file. If not given, no text output will be " +
			"generated.", "")
	dist := myflag.Int("range", "r", "integer",
			"Range around tile center to plot. Default is 5000.", 5000)
	bin := myflag.Int("bin", "", "integer",
			"Size of each bin in the histogram. Default is 1.", 1)
	
	// Parse!
	arguments.err = myflag.Parse()
	if arguments.err != nil { return }
	if !myflag.HasAny() { return }
	
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
	arguments.txt = *txt
	
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


// ***** AGGREGATION **********************************************************

// Creates an aggregation value slice for the given bed file.
func aggregate(path string, idx []index, dist int) ([][]float64, error) {
	f, err := os.Open(path)
	if err != nil { return nil, err }
	scanner := bed.NewScanner(f)
	
	result := make([][]float64, len(idx))
	for i := range result {
		result[i] = make([]float64, dist*2 + 1)
	}
	
	lineCount := 0
	
	for scanner.Scan() {
		lineCount++
		b := scanner.Bed()
		pos := (b.Start + b.End) / 2
		
		for i := range idx {
			idx[i].collect(b.Chr, pos, result[i])
		}
	}
	
	// Normalize by number of lines (average signal).
	for i := range result {
		for j := range result[i] {
			result[i][j] /= float64(lineCount)
		}
	}
	
	return result, nil
}


// ***** PYTHON INTERFACE *****************************************************

// Converts a float slice to a python list literal.
func floatsToText(values []float64) []byte {
	result := []byte("[")
	for _,v := range values {
		result = append(result, fmt.Sprintf("%v,", v)...)
	}
	result = append(result, "]"...)
	
	return result
}

// Converts an int slice to a python list literal.
func intsToText(values []int) []byte {
	result := []byte("[")
	for _,v := range values {
		result = append(result, fmt.Sprintf("%v,", v)...)
	}
	result = append(result, "]"...)
	
	return result
}

// Plots the given data using python. An empty output file name will result in
// only showing the plot.
func plotWithPython(filesData [][]float64, xvals []int, labels []string,
		outFile string) {
	src := bytes.NewBuffer(nil)
	
	// Create imports.
	fmt.Fprintf(src, "import matplotlib.pyplot as plt\n")
	
	// Find min and max for axes.
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
	fmt.Fprintf(src, "plt.plot([0,0],[%f,%f],'--k')\n", axesYMin, axesYMax)
	
	// Add plot for each file.
	for i,values := range filesData {
		fmt.Fprintf(src, "plt.plot(%s,%s,linewidth=2,label='%s')\n",
				intsToText(xvals), floatsToText(values), labels[i])
	}
	
	// Add figure settings.
	fmt.Fprintf(src, "plt.title('Aggregation plot')\n")
	fmt.Fprintf(src, "plt.xlabel('Distance from region center')\n")
	fmt.Fprintf(src, "plt.ylabel('Average signal')\n")
	fmt.Fprintf(src, "plt.axis([%f,%f,%f,%f])\n",
			axesXMin, axesXMax, axesYMin, axesYMax)
	fmt.Fprintf(src, "plt.legend(loc='upper right')\n")
	
	// Save to file command.
	if outFile == "show" {
		fmt.Fprintf(src, "plt.show()")
	} else {
		fmt.Fprintf(src, "plt.savefig('%s',dpi=150)", outFile)
	}
	
	cmd := exec.Command("python")
	cmd.Stdin = src
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}


// ***** TEXT OUTPUT **********************************************************

func printData(data [][]float64, xvals []int, labels []string,
		file string) error {
	// Input assertions.
	if len(labels) == 0 {
		panic("Empty label set is invalid.")
	}
	if len(xvals) == 0 {
		panic("Empty x-value set is invalid.")
	}
	if len(data) == 0 {
		panic("Empty data is invalid.")
	}
	if len(data) != len(labels) {
		panic("Data and labels are of different lengths.")
	}
	if len(data[0]) != len(xvals) {
		panic(fmt.Sprintf("Data and x-values are of different lengths: %v, %v",
				len(data[0]), len(xvals)))
	}

	// Open output file.
	fout, err := os.Create(file)
	if err != nil { return err }
	defer fout.Close()
	
	bout := bufio.NewWriter(fout)
	defer bout.Flush()
	
	// Print labels.
	fmt.Fprint(bout, "distance")
	for i := range labels {
		fmt.Fprintf(bout, "\t%s", labels[i])
	}
	
	fmt.Fprint(bout, "\n")
	
	// Print data.
	for i := range data[0] {
		fmt.Fprintf(bout, "%v", xvals[i])
		
		for j := range data {
			fmt.Fprintf(bout, "\t%v", data[j][i])
		}
		
		fmt.Fprint(bout, "\n")
	}
	
	return nil
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

// Normalize such that the medians are all 1.
func normalize(data [][]float64) {
	for i := range data {
		med := minFloat(data[i])
		if med == 0 { continue }
		
		for j := range data[i] {
			data[i][j] /= med
		}
	}
}

// Returns the minimal value.
func minFloat(values []float64) float64 {
	result := values[0]
	for _, v := range values {
		if v < result { result = v }
	}
	return result
}

// Returns the median value.
func median(values []float64) float64 {
	sorted := make([]float64, len(values))
	copy(sorted, values)
	sort.Sort(sort.Float64Slice(sorted))
	
	return sorted[len(sorted) / 2]
}

const usage =
`Creates aggregation plots of average signals around tiles.

Written by Amit Lavon (amitlavon1@gmail.com).

Usage:
aggplot [options] <bed/graph file 1> <bed/graph file 2> <bed/graph file 3>...

Choose either 1 bed-graph to many beds using '-bedgraph', or 1 bed to many
bedgraphs using '-bed'.

Options:
`



