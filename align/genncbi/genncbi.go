// Command genncbi generates substitution matrices from NCBI text files.
package main

import (
	"flag"
	"fmt"
	"go/format"
	"os"

	"github.com/fluhus/biostuff/align"
	"github.com/fluhus/biostuff/formats/smtext"
)

var (
	varName = flag.String("v", "", "Name of the variable")
)

func main() {
	flag.Parse()
	if *varName == "" {
		fmt.Fprintln(os.Stderr, "Please provide a variable name with -v.")
		os.Exit(2)
	}
	fmt.Fprintln(os.Stderr, "Reading from stdin.")
	m, err := smtext.ReadNCBI(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to parse matrix:", err)
		os.Exit(2)
	}
	m[[2]byte{align.Gap, align.Gap}] = 0
	src := []byte(fmt.Sprintf("package align\n\nfunc init() {\n%s = %#v}",
		*varName, m))
	src, err = format.Source(src)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to format source:", err)
		os.Exit(2)
	}
	fmt.Println(string(src))
}
