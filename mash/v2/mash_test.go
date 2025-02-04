package mash

import (
	"math"
	"math/rand/v2"
	"slices"
	"testing"

	"github.com/fluhus/biostuff/sequtil"
)

func TestSequences(t *testing.T) {
	seqs := [][]byte{
		[]byte("AGATTTTTCTCCCAACGAAACTTTACAGCACGCTAGTTTACGGGCACTCC"),
		[]byte("GCCCTATCTGGGGAGAAATCGTAGTGAGAGACCGAGGTGGCCCCACGCAC"),
		[]byte("TACGACTGGAAGAGCCCATGCACGGATGCTGCTACTCGCATTGGTTTACG"),
	}
	m1 := New(5, 21).AddIter(slices.Values(seqs))
	m2 := New(5, 21).Add(seqs[2])
	m2.Add(seqs[1])
	m2.Add(seqs[0])

	v1 := m1.mh.View()
	v2 := m2.mh.View()
	if !slices.Equal(v1, v2) {
		t.Fatalf("AttIter(...) != Add(...), %v != %v", v1, v2)
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

	d := New(10000, 21).Add(seq1).Distance(New(10000, 21).Add(seq2))
	dif := math.Abs(d - 0.05)
	if dif > 0.005 {
		t.Fatalf("Distance(...)=%v, want %v", d, 0.05)
	}
}
