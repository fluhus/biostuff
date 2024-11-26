// Package mash provides sequence MinHashing and Mash distance calculation.
package mash

import (
	"bytes"
	"math"

	"github.com/fluhus/biostuff/sequtil"
	"github.com/fluhus/gostuff/minhash"
	"github.com/spaolacci/murmur3"
)

// Seed is the hash seed.
// Affects subsequent calls to [Sequences] and [Add].
var Seed uint32 = 0

// Add adds the given sequences to an existing MinHash
// using subsequences of length k.
// Equivalent to calling [Sequences] on the old and new sequences together.
func Add(mh *minhash.MinHash[uint64], k int, seqs ...[]byte) {
	h := murmur3.New64WithSeed(Seed)
	for _, seq := range seqs {
		for b := range sequtil.CanonicalSubsequences(bytes.ToUpper(seq), k) {
			h.Reset()
			h.Write(b)
			mh.Push(h.Sum64())
		}
	}
	mh.Sort()
}

// Sequences returns a single MinHash for seqs with n elements and for
// subsequences of length k.
func Sequences(n, k int, seqs ...[]byte) *minhash.MinHash[uint64] {
	mh := minhash.New[uint64](n)
	Add(mh, k, seqs...)
	return mh
}

// Distance returns the Mash distance between two MinHash collections.
func Distance(mh1, mh2 *minhash.MinHash[uint64], k int) float64 {
	return FromJaccard(mh1.Jaccard(mh2), k)
}

// FromJaccard returns the Mash distance given a Jaccard similarity.
func FromJaccard(jac float64, k int) float64 {
	if jac == 0 {
		return 1
	}
	return min(-math.Log(2*jac/(1+jac))/float64(k), 1)
}
