// Convinience functions for CPU profiling.
package mypprof

import (
	"runtime/pprof"
	"os"
	"bufio"
	"fmt"
)

var fout *os.File
var bout *bufio.Writer

// Starts CPU profiling and writes to the given file. Returns an error if
// something goes wrong.
func Start(file string) error {
	if fout != nil {
		return fmt.Errorf("Already profiling.")
	}

	var err error
	fout, err = os.Create(file)
	if err != nil {
		fout, bout = nil, nil
		return err
	}
	bout = bufio.NewWriter(fout)

	pprof.StartCPUProfile(bout)

	return nil
}

// Stops CPU profiling and closes the output file. If called without calling
// Start, does nothing.
func Stop() {
	if fout == nil {
		return
	}

	pprof.StopCPUProfile()
	bout.Flush()
	fout.Close()
	fout, bout = nil, nil
}
