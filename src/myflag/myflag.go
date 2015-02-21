// Handy functionality for flag parsing.
package myflag

import (
	"fmt"
	"flag"
	"bytes"
)

// Holds flags.
var flags *flag.FlagSet

// Help message for the flags. Accumulates help messages as flags are set.
var flagsHelp = bytes.NewBuffer(nil)

func init() {
	flags = flag.NewFlagSet("", flag.ContinueOnError)
}

// Registers a new int flag.
func Int(name string, shortName string, typ string, description string,
		dflt int) *int {
	// Register long name
	result := flags.Int(name, dflt, "")
	
	// Register short name
	if shortName != "" {
		flags.IntVar(result, shortName, dflt, "")
	}
	
	// Add help message
	if shortName != "" {
		fmt.Fprintf(flagsHelp, "\t-%s <%s>\n", shortName, typ)
	}
	fmt.Fprintf(flagsHelp, "\t-%s <%s>\n", name, typ)
	fmt.Fprintf(flagsHelp, "\t\t%s (default: %d)\n", description, dflt)
	
	return result
}
