// Package mash provides sequence MinHashing and Mash distance calculation.
package mash

import (
	"bytes"
	"fmt"
	"iter"
	"math"

	"github.com/fluhus/biostuff/sequtil"
	"github.com/fluhus/gostuff/hashx"
	"github.com/fluhus/gostuff/minhash"
)

// A Masher min-hashes sequences and calculates Mash distances.
type Masher struct {
	mh   *minhash.MinHash[uint64]
	seed uint32
	k    int
}

// NewSeed returns a new Masher with n hashes, k-long subsequences
// and the given seed.
func NewSeed(n, k int, seed uint32) *Masher {
	return &Masher{
		mh:   minhash.New[uint64](n),
		seed: seed,
		k:    k,
	}
}

// New returns a new Masher with n hashes, k-long subsequences
// and seed 0.
func New(n, k int) *Masher {
	return NewSeed(n, k, 0)
}

// Add adds the given sequence to the MinHash.
func (m *Masher) Add(seq []byte) *Masher {
	h := hashx.NewSeed(m.seed)
	for b := range sequtil.CanonicalSubsequences(bytes.ToUpper(seq), m.k) {
		m.mh.Push(h.Bytes(b))
	}
	m.mh.Sort()
	return m
}

// AddIter adds the given sequences to the MinHash.
func (m *Masher) AddIter(seqs iter.Seq[[]byte]) *Masher {
	h := hashx.NewSeed(m.seed)
	for seq := range seqs {
		for b := range sequtil.CanonicalSubsequences(bytes.ToUpper(seq), m.k) {
			m.mh.Push(h.Bytes(b))
		}
	}
	m.mh.Sort()
	return m
}

// Distance returns the Mash distance between this Masher and another one.
//
// Both Mashers need to have the same k's and seeds.
func (m *Masher) Distance(other *Masher) float64 {
	if m.k != other.k {
		panic(fmt.Sprintf("mismatching k: this.k=%v, other.k=%v",
			m.k, other.k))
	}
	if m.seed != other.seed {
		panic(fmt.Sprintf("mismatching seed: this.seed=%v, other.seed=%v",
			m.seed, other.seed))
	}
	return FromJaccard(m.mh.Jaccard(other.mh), m.k)
}

// FromJaccard returns the Mash distance given a Jaccard similarity.
func FromJaccard(jac float64, k int) float64 {
	if jac == 0 {
		return 1
	}
	return min(-math.Log(2*jac/(1+jac))/float64(k), 1)
}
