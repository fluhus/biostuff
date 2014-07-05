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
	"runtime"
	"runtime/pprof"
)

// Short for fmt.Fprintln(stderr, ...)
func pe(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

// Short for fmt.Fprintf(stderr, ...)
func pef(s string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, s, a...)
}

// Returns the best scoring candidate positions from the search results.
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
const profiling = false

func main() {
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
	defer fastqFile.Close()
	
	// Open sam
	pe("creating sam...")
	samFile, err := os.Create(os.Args[3])
	if err != nil {
		pe("error creating output sam:", err.Error())
		return
	}
	defer samFile.Close()

	// Start profiling
	if profiling {
		profFile, _ := os.Create("jo.prof")
		defer profFile.Close()
		pprof.StartCPUProfile(profFile)
		defer pprof.StopCPUProfile()
	}
	
	// Create index
	pe("building index...")
	tools.Tic()
	idx, err := myindex.New(fa, 12, 1)
	if err != nil {
		pe("error building index:", err.Error())
	}
	pe("took", tools.Toc())
	
	// Buffers
	fastqBuf := bufio.NewReader(fastqFile)
	samBuf := bufio.NewWriter(samFile)
	
	// Prepare threads
	runtime.GOMAXPROCS(runtime.NumCPU())
	numberOfThreads := runtime.NumCPU()
	
	fastqChannel := make(chan *fastq.Fastq, numberOfThreads)
	samChannel := make(chan *sam.Sam, numberOfThreads)
	searcherDoneChannel := make(chan int, numberOfThreads)
	printerDoneChannel := make(chan int, 1)
	
	// Create new searcher threads
	for i := 0; i < numberOfThreads; i++ {
		go func() {
			for fq := range fastqChannel {
				// Search
				matches := idx.Search(fq.Sequence, 1, true)
		
				// Pick the best scoring positions
				leaders := scoreLeaders(matches, 1)
		
				// If mapped
				samLine := &sam.Sam{}
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
		
				samChannel <- samLine
			}
			
			searcherDoneChannel <- 1
		}()
	}
	
	// Create a new sam printer thread
	go func() {
		for samLine := range samChannel {
			fmt.Fprintln(samBuf, samLine)
		}
		
		samBuf.Flush()
		printerDoneChannel <- 1
	}()
	
	// Look up reads
	pef("looking up reads (%d threads)...\n", numberOfThreads)
	tools.Tic()
	var fq *fastq.Fastq
	
	// Read fastq
	for fq, err = fastq.ReadNext(fastqBuf); err == nil;
			fq, err = fastq.ReadNext(fastqBuf) {
		fastqChannel <- fq
	}
	close(fastqChannel)
	
	if err != io.EOF {
		pe("error reading fastq:", err)
		return
	}
	
	// Join searcher threads
	for i := 0; i < numberOfThreads; i++ {
		<-searcherDoneChannel
	}
	
	// Signal end of stream to printer and join
	close(samChannel)
	<-printerDoneChannel
	
	pe("took", tools.Toc())
}
