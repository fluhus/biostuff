package main

import (
	"time"
	"math/rand"
	"os"
	"fmt"
	"bioformats/fasta"
	"strdist"
)

func pe(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func pef(s string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, s, a...)
}

func main() {
	pe("start")
	
	
	
	pe("end")
}










