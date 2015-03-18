// Annotates differentially methylated regions.
package main

import (
	"os"
	"fmt"
	"sort"
	"bufio"
	"myflag"
	"strings"
	"strconv"
	"runtime/pprof"
)

func main() {
	// Parse arguments.
	parseArguments()
	if !myflag.HasAny() {
		fmt.Println(usage)
		fmt.Println(myflag.HelpString())
		os.Exit(1)
	}

	if arguments.err != nil {
		fmt.Println("Error parsing arguments:", arguments.err)
		os.Exit(1)
	}

	// Load input files.
	fmt.Printf("Loading %s:\n", arguments.label1)
	t1 := make(map[string]tiles)
	for _, f := range arguments.inputs1 {
		fmt.Printf("\t%s\n", f)
		err := tileFile(f, t1)
		if err != nil {
			fmt.Printf("Error loading %s: %v\n", f, err)
			os.Exit(2)
		}
	}

	fmt.Printf("Loading %s:\n", arguments.label2)
	t2 := make(map[string]tiles)
	for _, f := range arguments.inputs2 {
		fmt.Printf("\t%s\n", f)
		err := tileFile(f, t2)
		if err != nil {
			fmt.Printf("Error loading %s: %v\n", f, err)
			os.Exit(2)
		}
	}

	// Open output file.
	fmt.Println("Diffing...")
	fout, err := os.Create(arguments.output)
	if err != nil {
		fmt.Println("Error opening output file:", err)
		os.Exit(2)
	}
	defer fout.Close()
	
	bout := bufio.NewWriter(fout)
	defer bout.Flush()

	fmt.Fprintf(bout, "chromosome\tposition\tmeth_ratio_%s\tmeth_ratio_%s\t" +
			"p_value\n", arguments.label1, arguments.label2)
	
	// Sort chromosomes, for sorted output.
	chrs := make([]string, 0, len(t1))
	for chr := range t1 {
		chrs = append(chrs, chr)
	}
	sort.Sort(chrSorter(chrs))

	// Compare methylation rates.
	for _, chr := range chrs {
		if t2[chr] == nil { continue }

		// Sort positions, for sorted output.
		poss := make([]int, 0, len(t1[chr]))
		for pos := range t1[chr] {
			poss = append(poss, pos)
		}
		sort.Sort(sort.IntSlice(poss))
		
		for _, pos := range poss {
			tile1 := t1[chr][pos]
			tile2 := t2[chr][pos]
			
			if tile2 == nil { continue }
			
			r1 := float64(tile1.methd) / float64(tile1.total)
			r2 := float64(tile2.methd) / float64(tile2.total)
			
			fmt.Fprintf(bout, "%s\t%d\t%f\t%f\t%f\n", chr, pos, r1, r2,
					tilediff(tile1, tile2))
		}
	}
}


// ***** TILING ***************************************************************

type tile struct {
	total int
	methd int
}

// Maps from position to tile.
type tiles map[int]*tile

// Adds the file's contents to the given map.
func tileFile(file string, out map[string]tiles) error {
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
		if out[chr] == nil {
			out[chr] = make(map[int]*tile)
		}
		
		// Round position to tile.
		pos = pos / 100 * 100
		
		// Create tile.
		if out[chr][pos] == nil {
			out[chr][pos] = &tile{}
		}
		
		out[chr][pos].total += int(total)
		out[chr][pos].methd += methd
	}
	
	if scanner.Err() != nil {
		return scanner.Err()
	}
	
	return nil
}

// Returns the p-value for the null hypothesis, that the 2 tiles come from the
// same distribution.
func tilediff(tile1, tile2 *tile) float64 {
	return bindiff(tile1.total, tile1.methd, tile2.total, tile2.methd)
}


// ***** PROFILING ************************************************************

var profOut *os.File

func startProfiling() {
	var err error
	profOut, err = os.Create("profile")
	if err != nil { return }
	err = pprof.StartCPUProfile(profOut)
	if err != nil {
		profOut.Close()
		profOut = nil
	}
}

func stopProfiling() {
	if profOut == nil { return }
	pprof.StopCPUProfile()
	profOut.Close()
}


// ***** CHROMOSOME NAME SORTER ***********************************************

type chrSorter []string
func (c chrSorter) Len() int {
	return len(c)
}
func (c chrSorter) Less(i, j int) bool {
	if len(c[i]) != len(c[j]) {
		return len(c[i]) < len(c[j])
	}
	return c[i] < c[j]
}
func (c chrSorter) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}


// ***** ARGUMENTS ************************************************************

var arguments struct {
	label1 string
	label2 string
	inputs1 []string
	inputs2 []string
	output string
	err error
}

func parseArguments() {
	label1 := myflag.String("label1", "L1", "string",
			"Label of the first group.", "group_1")
	label2 := myflag.String("label2", "L2", "string",
			"Label of the second group.", "group_2")
	in1 := myflag.String("in1", "1", "paths",
			"Comma separated meth files of the first group.", "")
	in2 := myflag.String("in2", "2", "paths",
			"Comma separated meth files of the second group.", "")
	out := myflag.String("out", "o", "path", "Output file.", "")

	arguments.err = myflag.Parse()
	if arguments.err != nil { return }

	// Handle labels.
	if *label1 == "" {
		arguments.err = fmt.Errorf("Label 1 is empty.")
		return
	}
	if *label2 == "" {
		arguments.err = fmt.Errorf("Label 2 is empty.")
		return
	}
	arguments.label1 = *label1
	arguments.label2 = *label2

	// Handle output.
	if *out == "" {
		arguments.err = fmt.Errorf("No output file given.")
		return
	}
	arguments.output = *out

	// Handle inputs.
	if *in1 == "" {
		arguments.err = fmt.Errorf("No input files for group 1.")
	}
	if *in2 == "" {
		arguments.err = fmt.Errorf("No input files for group 2.")
	}

	arguments.inputs1 = strings.Split(*in1, ",")
	arguments.inputs2 = strings.Split(*in2, ",")

	for _, f := range arguments.inputs1 {
		if f == "" {
			arguments.err = fmt.Errorf("Empty input file path in group 1.")
			return
		}
	}
	for _, f := range arguments.inputs2 {
		if f == "" {
			arguments.err = fmt.Errorf("Empty input file path in group 2.")
			return
		}
	}
}

var usage =
`Compares methylation levels of files generated with BSMap's methratio.

Written by Amit Lavon (amitlavon1@gmail.com).

Usage:
methdiff -1 wt_1.meth,wt_2.meth -2 mut_1.meth,mut_2.meth -o my_file.txt

Options:`
