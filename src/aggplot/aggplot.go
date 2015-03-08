package main

import (
	"os"
	"fmt"
	"sort"
	"bufio"
	"myflag"
	"bioformats/bed"
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
		normalize(data)
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

// Normalize such that the mins are all 1.
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



