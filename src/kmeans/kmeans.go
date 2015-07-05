// Performs k-means analysis on data.
package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
	"math/rand"
	"time"
	"myflag"
)

func main() {
	// Seed random.
	rand.Seed(time.Now().UnixNano())

	// Parse arguments.
	err := parseArgs()
	if err != nil {
		pe("Bad arguments:", err)
		os.Exit(1)
	}
	if args.help {
		pe(help)
		pef("%s", myflag.HelpString())
		os.Exit(1)
	}

	var mat matrix
	var raw []string

	// Skip rows.
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < *args.skipRows; i++ {
		scanner.Scan()
		raw = append(raw, scanner.Text())
	}

	// Scan data.
	lineNum := *args.skipRows
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		raw = append(raw, line)
		
		vals := strings.Split(line, "\t")
		vals = vals[*args.skipCols:]  // Skip columns.

		// Check for number of columns.
		if mat.nrows() > 0 && len(vals) != mat.ncols() {
			pef("Error at line %d: expected %d columns and found %d.\n",
					lineNum, mat.ncols(), len(vals))
			os.Exit(2)
		}

		// Parse values.
		matRow := make([]float64, len(vals))
		for i := range vals {
			var err error
			matRow[i], err = strconv.ParseFloat(vals[i], 64)
			if err != nil {
				pef("Error at line %d: bad number: '%s'\n", lineNum, vals[i])
				os.Exit(2)
			}
		}

		mat.addRow(matRow)
	}

	// Flip if needed.
	if *args.cols {
		mat.flip()
	}

	// Go k-means!
	tags, d := mat.kmeans(*args.k)
	pe(d)

	// Print output.
	if *args.bare {
		// Bare -> just print the tags in one column.
		for i := range tags {
			fmt.Println(tags[i])
		}
	} else {
		// Not bare.
		if *args.cols {
			// Add a new row.
			for i := range raw {
				fmt.Println(raw[i])
			}
			for i := range tags {
				if i > 0 {
					fmt.Printf("\t%d", tags[i])
				} else {
					fmt.Printf("%d", tags[i])
				}
			}
			fmt.Println()
		} else {
			// Add a new column.
			rawHead := raw[:*args.skipRows]
			rawData := raw[*args.skipRows:]
			for i := range rawHead {
				fmt.Printf("%s\t\n", rawHead[i])
			}
			for i := range rawData {
				fmt.Printf("%s\t%d\n", rawData[i], tags[i])
			}
		}
	}
}

func pe(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func pef(s string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, s, a...)
}


// ----- ARGUMENTS -------------------------------------------------------------

var args struct {
	cols *bool
	k *int
	skipRows *int
	skipCols *int
	bare *bool
	help bool
}

func parseArgs() error {
	args.k = myflag.Int("k", "", "integer", "K parameter for K-means." +
			" Must be at least 1.", 0)
	args.skipRows = myflag.Int("skiprows", "sr", "integer", "Skip given " +
			"number of rows.", 0)
	args.skipCols = myflag.Int("skipcols", "sc", "integer", "Skip given " +
			"number of columns.", 0)
	args.bare = myflag.Bool("bare", "b", "Print bare tags instead of entire " +
			"table.", false)
	args.cols = myflag.Bool("cols", "c", "Cluster columns instead of rows.",
			false)

	err := myflag.Parse()
	if err != nil {
		return err
	}

	if !myflag.HasAny() {
		args.help = true
		return nil
	}

	if *args.k < 1 {
		return fmt.Errorf("Please supply K that is at least 1.")
	}
	if *args.skipRows < 0 || *args.skipCols < 0 {
		return fmt.Errorf("Number of rows/columns to skip must be "+
				"non-negative.")
	}

	return nil
}

var help =
`Performs k-means analysis on the rows of the given data.

The program reads data from the standard input and writes data with tags to the
standard output. The distortion will be printed to the standard error.

Written by Amit Lavon (amitlavon1@gmail.com)

Usage:
kmeans [options] < input_file > output_file

Accepted options:`


// ----- MATRIX ----------------------------------------------------------------

// A matrix type for convenience functions.
type matrix struct {
	data [][]float64
}

// Adds a row to the matrix. The first row sets the matrix's width permanently.
func (m *matrix) addRow(row []float64) {
	if len(m.data) > 0 && len(row) != len(m.data[0]) {
		panic(fmt.Sprintf("Bad number of elements: expected %d, got %d.",
				len(m.data[0]), len(row)))
	}
	m.data = append(m.data, row)
}

// Returns the number of rows in the matrix.
func (m *matrix) nrows() int {
	return len(m.data)
}

// Returns the number of columns in the matrix.
func (m *matrix) ncols() int {
	if len(m.data) == 0 {
		return 0
	} else {
		return len(m.data[0])
	}
}

// Transposes the matrix.
func (m *matrix) flip() {
	if m.nrows() == 0 || m.ncols() == 0 {
		return
	}

	data := make([][]float64, len(m.data[0]))
	for i := range data {
		data[i] = make([]float64, len(m.data))
		for j := range data[i] {
			data[i][j] = m.data[j][i]
		}
	}

	m.data = data
}


// ----- K-MEANS ---------------------------------------------------------------

// Performs a k-means analysis on the matrix.
func (m *matrix) kmeans(k int) (tags []int, dist float64) {
	// Must be at least 1.
	if k < 1 {
		panic(fmt.Sprint("Bad k:", k))
	}

	// If k is too large - that's ok just reduce to avoid out-of-range.
	if k > m.nrows() {
		k = m.nrows()
	}

	// Create initial centroids.
	initCents := rand.Perm(m.nrows())[:k]
	cents := make([][]float64, k)
	for i := range cents {
		cents[i] = make([]float64, m.ncols())
		copy(cents[i], m.data[initCents[i]])
	}

	// First tagging.
	tags = m.tag(cents)
	dist = m.distortion(cents, tags)
	distOld := 2 * dist

	// Iterate until converged.
	for dist > distOld || dist / distOld < 0.999 {
		distOld = dist
		cents = m.cent(tags, k)
		tags = m.tag(cents)
		dist = m.distortion(cents, tags)
	}

	return
}

// Tags each row with the index of its nearest centroid.
func (m *matrix) tag(cents [][]float64) []int {
	if len(cents) == 0 {
		panic("Cannot tag on 0 centroids.")
	}
	
	tags := make([]int, m.nrows())

	for i := range tags {
		// Find nearest centroid.
		tags[i] = 0
		d := distance(cents[0], m.data[i])
		for j := 1; j < len(cents); j++ {
			dj := distance(cents[j], m.data[i])
			if dj < d {
				d = dj
				tags[i] = j
			}
		}
	}

	return tags
}

// Calculates the distortion of the given tagging and centroids.
func (m *matrix) distortion(cents [][]float64, tags []int) float64 {
	if len(tags) != m.nrows() {
		panic(fmt.Sprintf("Non-matching lengths of matrix and tags: %d, %d",
				m.nrows(), len(tags)))
	}
	if m.nrows() == 0 {
		return 0
	}

	d := 0.0
	for i := range tags {
		d += distance(cents[tags[i]], m.data[i])
	}

	return d / float64(m.nrows())
}

// Calculates the new centroids, according to average of tagged rows in each
// group.
func (m *matrix) cent(tags []int, k int) [][]float64 {
	cents := make([][]float64, k)
	for i := range cents {
		cents[i] = make([]float64, m.ncols())
	}
	counts := make([]int, k)

	for i := range m.data {
		counts[tags[i]]++
		for j := range m.data[i] {
			cents[tags[i]][j] += m.data[i][j]
		}
	}

	for i := range cents {
		for j := range cents[i] {
			cents[i][j] /= float64(counts[tags[i]])
		}
	}

	return cents
}

// Returns the L1 (Manhattan) distance between 2 vectors.
func distance(a, b []float64) float64 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("Non-matching lengths: %d, %d",
				len(a), len(b)))
	}
	
	d := 0.0
	for i := range a {
		d += abs(a[i] - b[i])
	}

	return d
}

// Returns the absolute value of a number.
func abs(a float64) float64 {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
