package main

// Handles interaction with python.

import (
	"os/exec"
	"fmt"
	"bytes"
	"math"
	"os"
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


