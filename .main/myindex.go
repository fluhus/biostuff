package main

import (
	"fmt"
	"time"
	"tools"
	"myindex"
	// "seqtools"
	// "math/rand"
	"bioformats/fasta"
)

var ticTime time.Time
func tic() {
	ticTime = time.Now()
}
func toc() time.Duration {
	return time.Since(ticTime)
}

func init() {
	tools.Randomize()
}

func main() {
	p := fmt.Println

	// read fasta
	fa, err := fasta.FastaFromFile("fasta/Yeast.fa")
	if err != nil { panic("could not open fasta: " + err.Error()) }
	// f = f[0:2]
	// f[0].Sequence = f[0].Sequence[0:4]
	// f[1].Sequence = f[1].Sequence[0:4]
	
	// make index
	p("building index...")
	tic()
	idx, err := myindex.New(fa, 9)
	p("took", toc())
	p(idx)
	p("err:", err)
	
	// stam search
	seq, seqchrom, seqpos := fa.Subsequence(50, 1234)
	p("seq", string(seq))
	p("len", len(seq))
	p("at chromosome", seqchrom, "position", seqpos)
	matches := idx.Search(seq)
	p("matches", len(matches))
	p("true score", matches[myindex.NewGenPos(0,0)])
	
	// score distribution
	dist := make([]int, 11)
	for pos, score := range matches {
		dist[score]++
		if score == 5 || score == 4 {
			p("match", score, string(fa[pos.Chr()].Sequence[pos.Pos() : pos.Pos() + 50]))
			p("at", pos)
		}
	}
	
	for i := range dist {
		fmt.Println(i, "\t", dist[i])
	}
}
