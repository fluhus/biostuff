// Aligner prototype.
package main

import (
	"os"
	"io"
	"fmt"
	"bufio"
	"tools"
	"myindex"
	"bioformats/sam"
	"bioformats/fasta"
	"bioformats/fastq"
	"runtime/pprof"
)

func pe(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func pef(s string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, s, a...)
}

func scoreLeaders(matches map[myindex.GenPos]int,
		howLess int) []myindex.GenPos {
	// Find max
	max := 0
	for _,score := range matches {
		if score > max {
			max = score
		}
	}
	
	var result []myindex.GenPos
	for match, score := range matches {
		if score >= max - howLess {
			result = append(result, match)
		}
	}
	
	return result
}

// If true, will generate a CPU profile
const profiling = true

func main() {
	if profiling {
		profFile, _ := os.Create("jo.prof")
		defer profFile.Close()
		pprof.StartCPUProfile(profFile)
		defer pprof.StopCPUProfile()
	}
	
	// Parse arguments
	if len(os.Args) != 4 {
		pe("Usage:\njo <reference fasta> <reads fastq> <output sam>")
		return;
	}

	// Load fasta
	pe("loading fasta...")
	tools.Tic()
	fa, err := fasta.FastaFromFile(os.Args[1])
	if err != nil {
		pe("error reading fasta:", err.Error())
		return
	}
	pe("took", tools.Toc())
	
	// Open fastq
	pe("opening fastq...")
	fastqFile, err := os.Open(os.Args[2])
	if err != nil {
		pe("error opening fastq:", err.Error())
		return
	}
	
	// Open sam
	pe("creating sam...")
	samFile, err := os.Create(os.Args[3])
	if err != nil {
		pe("error creating output sam:", err.Error())
		return
	}
	
	// Create index
	pe("building index...")
	tools.Tic()
	idx, err := myindex.New(fa, 12, 1)
	// pe(idx)
	if err != nil {
		pe("error building index:", err.Error())
	}
	pe("took", tools.Toc())
	
	// Buffers
	fastqBuf := bufio.NewReader(fastqFile)
	samBuf := bufio.NewWriter(samFile)
	
	// Look up reads
	pe("looking up reads...")
	tools.Tic()
	var fq *fastq.Fastq
	
	for fq, err = fastq.ReadNext(fastqBuf); err == nil;
			fq, err = fastq.ReadNext(fastqBuf) {
		// Search
		matches := idx.Search(fq.Sequence, 1, true)
		
		// Pick the best scoring positions
		leaders := scoreLeaders(matches, 1)
		
		// If mapped
		var samLine sam.Sam
		if len(leaders) == 1 {
			leader := leaders[0]
			samLine.Rname = string(fa[leader.Chr()].Title)
			samLine.Pos = leader.Pos() + 1 // sam are 1-based
			samLine.Mapq = 60  //TODO learn the actual mapq
		
		// If not mapped
		} else {
			samLine.Rname = "*"
			samLine.Pos = 0
			samLine.Mapq = 0
		}
		
		samLine.Qname = string(fq.Id)
		samLine.Cigar = "*"
		samLine.Rnext = "*"
		samLine.Pnext = "*"
		samLine.Seq =  string(fq.Sequence)
		samLine.Qual = string(fq.Quals)
		
		fmt.Fprintln(samBuf, samLine)
	}
	samBuf.Flush()
	
	if err != io.EOF {
		pe("error reading fastq:", err)
		return
	}
	
	pe("took", tools.Toc())
}
