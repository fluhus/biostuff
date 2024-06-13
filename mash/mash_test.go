package mash

import (
	"math"
	"math/rand/v2"
	"slices"
	"testing"

	"github.com/fluhus/biostuff/sequtil/v2"
)

func TestSequences(t *testing.T) {
	seqs := [][]byte{
		[]byte("AGATTTTTCTCCCAACGAAACTTTACAGCACGCTAGTTTACGGGCACTCC"),
		[]byte("GCCCTATCTGGGGAGAAATCGTAGTGAGAGACCGAGGTGGCCCCACGCAC"),
		[]byte("TACGACTGGAAGAGCCCATGCACGGATGCTGCTACTCGCATTGGTTTACG"),
	}
	mh1 := Sequences(5, 21, seqs...)
	mh2 := Sequences(5, 21, seqs[0])
	Add(mh2, 21, seqs[1])
	Add(mh2, 21, seqs[2])

	v1 := mh1.View()
	v2 := mh2.View()
	if !slices.Equal(v1, v2) {
		t.Fatalf("Sequences(...) != Add(...), %v != %v", v1, v2)
	}
}

func TestDistance(t *testing.T) {
	const n = 1000000

	// Create sequence.
	seq1 := make([]byte, n)
	for i := range seq1 {
		seq1[i] = sequtil.Iton(rand.N(4))
	}

	// Create mutated sequence.
	seq2 := slices.Clone(seq1)
	for _, i := range rand.Perm(n)[:n/20] {
		seq2[i] = sequtil.Iton((sequtil.Ntoi(seq2[i]) + 1 + rand.N(3)) % 4)
	}

	d := Distance(Sequences(10000, 21, seq1), Sequences(10000, 21, seq2), 21)
	dif := math.Abs(d - 0.05)
	if dif > 0.001 {
		t.Fatalf("Distance(...)=%v, want %v", d, 0.05)
	}
}
