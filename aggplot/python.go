package main

// Handles interaction with python.

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"os/exec"
)

// Converts a float slice to a python list literal.
func floatsToText(values []float64) []byte {
	result := []byte("[")
	for _, v := range values {
		result = append(result, fmt.Sprintf("%v,", v)...)
	}
	result = append(result, "]"...)

	return result
}

// Converts an int slice to a python list literal.
func intsToText(values []int) []byte {
	result := []byte("[")
	for _, v := range values {
		result = append(result, fmt.Sprintf("%v,", v)...)
	}
	result = append(result, "]"...)

	return result
}

// Plots the given data using python. An output file name equals to "show" will
// result in only showing the plot.
func plotWithPython(filesData [][]float64, xvals []int, binSize int,
	labels []string, baseFile string, outFile string) {
	src := bytes.NewBuffer(nil)

	// Create imports.
	fmt.Fprintf(src, "import matplotlib.pyplot as plt\n")
	fmt.Fprintf(src, "import matplotlib.cm as cm\n")
	fmt.Fprintf(src, "palette=cm.brg\n")

	// Find min and max for axes.
	minValue := math.MaxFloat64
	maxValue := -math.MaxFloat64
	for i := range filesData {
		for _, v := range filesData[i] {
			if v < minValue {
				minValue = v
			}
			if v > maxValue {
				maxValue = v
			}
		}
	}

	axesXMin := float64(-arguments.dist)
	axesXMax := float64(arguments.dist)
	axesYMin := minValue - 0.1*(maxValue-minValue)
	axesYMax := maxValue + 0.3*(maxValue-minValue)

	// Add x=0 marker
	fmt.Fprintf(src, "plt.plot([0,0],[%f,%f],'--k')\n", axesYMin, axesYMax)

	// Add plot for each file.
	for i, values := range filesData {
		fmt.Fprintf(src, "plt.plot(%s[::%d],%s[::%d],linewidth=2,label='%s',"+
			"c=palette(%f))\n",
			intsToText(xvals), binSize, floatsToText(values), binSize,
			labels[i], float64(i)/float64(len(filesData)-1))
	}

	// Add figure settings.
	fmt.Fprintf(src, "plt.title('Aggregation plot\\n%s')\n", baseFile)
	fmt.Fprintf(src, "plt.xlabel('Distance from region center')\n")
	fmt.Fprintf(src, "plt.ylabel('Average signal')\n")
	fmt.Fprintf(src, "plt.axis([%f,%f,%f,%f])\n",
		axesXMin, axesXMax, axesYMin, axesYMax)
	fmt.Fprintf(src, "plt.legend(loc='upper right')\n")
	fmt.Fprintf(src, "plt.gcf().set_size_inches(20, 10)\n")

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
