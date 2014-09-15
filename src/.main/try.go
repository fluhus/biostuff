package main

import (
	"os"
	"fmt"
	"tools"
	"bioformats/fastq"
)

func pe(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func main() {
	pe("start")
	tools.Randomize()
	
	f := &fastq.Fastq{}
	f.Id = []byte("Amit")
	f.Sequence = []byte("AAAAAAAAAA")
	f.Quals = []byte("**********")
	
	f.ApplyQuals(fastq.Illumina18)
	pe(f)
	
	pe("end")
}










