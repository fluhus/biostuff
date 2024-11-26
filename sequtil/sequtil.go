// Package sequtil provides functions for processing genetic sequences.
//
// # Data Type For Sequences
//
// In this repository, sequences are represented as byte slices. This is meant
// to keep them familiar and predictable. This design favors having a buffet of
// functions for manipulating basic types, over having a dedicated sequence type with
// methods.
//
// # Mutating Sequences
//
// Mutations can be made using basic slice operations.
//
//	// Substitution
//	seq[4] = 'G'
//
//	// Deletion of bases 10-12 (including 12)
//	copy(seq[10:], seq[13:])
//	seq = seq[:len(seq)-3]
//
//	// Insertion at position 4
//	insert := []byte{...}
//	seq = append(append(append([]byte{}, seq[:4]...), insert...), seq[4:]...)
//
// # The U Nucleotide
//
// This package currently ignores the existence of uracil. Adding support for uracil
// means increasing the complexity of the API without adding new capabilities. The
// current solution is to substitute U's with T's before calling this package.
//
// This may change in the future, keeping backward compatibility.
package sequtil

import (
	"bytes"
	"fmt"
	"iter"
	"strings"
)

// Maps nucleotide byte value to its int value.
var ntoi []int

// Maps nucleotide byte value to its complement.
var complementBytes []byte

func init() {
	// Initialize ntoi.
	ntoi = make([]int, 256)
	for i := range ntoi {
		ntoi[i] = -1
	}
	ntoi['a'], ntoi['A'] = 0, 0
	ntoi['c'], ntoi['C'] = 1, 1
	ntoi['g'], ntoi['G'] = 2, 2
	ntoi['t'], ntoi['T'] = 3, 3

	// Initialize complementBytes.
	complementBytes = make([]byte, 256)
	complementBytes['a'], complementBytes['A'] = 't', 'T'
	complementBytes['c'], complementBytes['C'] = 'g', 'G'
	complementBytes['g'], complementBytes['G'] = 'c', 'C'
	complementBytes['t'], complementBytes['T'] = 'a', 'A'
	complementBytes['n'], complementBytes['N'] = 'n', 'N'
}

// Ntoi converts a nucleotide to an int.
// Aa:0 Cc:1 Gg:2 Tt:3. Other values return -1.
func Ntoi(nuc byte) int {
	return ntoi[nuc]
}

// Iton converts an int to a nucleotide character.
// 0:A 1:C 2:G 3:T. Other values return N.
func Iton(num int) byte {
	switch num {
	case 0:
		return 'A'
	case 1:
		return 'C'
	case 2:
		return 'G'
	case 3:
		return 'T'
	default:
		return 'N'
	}
}

// Converts a single base to its complement.
func complementByte(b byte) byte {
	rc := complementBytes[b]
	if rc == 0 {
		panic(fmt.Sprintf("unexpected base value: %q, want aAcCgGtTnN", b))
	}
	return rc
}

// ReverseComplement appends to dst the reverse complement of src and returns
// the new dst. Characters not in "aAcCgGtTnN" will cause a panic.
func ReverseComplement(dst, src []byte) []byte {
	for i := len(src) - 1; i >= 0; i-- {
		dst = append(dst, complementByte(src[i]))
	}
	return dst
}

// ReverseComplementString returns the reverse complement of s.
// Characters not in "aAcCgGtTnN" will cause a panic.
func ReverseComplementString(s string) string {
	builder := &strings.Builder{}
	builder.Grow(len(s))
	for i := len(s) - 1; i >= 0; i-- {
		builder.WriteByte(complementByte(s[i]))
	}
	return builder.String()
}

// DNATo2Bit appends to dst the 2-bit representation of the DNA sequence in src,
// and returns the new dst. Characters not in "aAcCgGtT" will cause a panic.
func DNATo2Bit(dst, src []byte) []byte {
	dn := len(dst)
	for i, b := range src {
		di := dn + i/4
		shift := 6 - i%4*2 // Make the first character the most significant.
		if shift == 6 {
			// Starting a new byte.
			dst = append(dst, 0)
		}
		dbInt := Ntoi(b)
		if dbInt == -1 {
			panic(fmt.Sprintf("unexpected base value: %q, want aAcCgGtT", b))
		}
		db := byte(dbInt) << shift
		dst[di] |= db
	}
	return dst
}

// DNAFrom2Bit appends to dst the nucleotides represented in 2-bit in src and
// returns the new dst. Only outputs characters in "ACGT".
func DNAFrom2Bit(dst, src []byte) []byte {
	for i := 0; i < len(src); i++ {
		dst = append(dst, dnaFrom2bit[src[i]][:]...)
	}
	return dst
}

// Maps 2-bit value to its expanded representation.
var dnaFrom2bit = make([][4]byte, 256)

// Initializes the dnaFrom2bit slice.
func init() {
	val := make([]byte, 4)
	for i := 0; i < 256; i++ {
		for j := 0; j < 4; j++ {
			// First base is the most significant digit.
			val[3-j] = Iton((i >> (2 * j)) & 3)
		}
		copy(dnaFrom2bit[i][:], val)
	}
}

// CanonicalSubsequences iterates over canonical k-long subsequences of seq.
// A canonical sequence is the lexicographically lesser out of a sequence and
// its reverse complement.
// Makes one call to ReverseComplement.
func CanonicalSubsequences(seq []byte, k int) iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		rc := ReverseComplement(make([]byte, 0, len(seq)), seq)
		nk := len(seq) - k + 1
		for i := range nk {
			kmer := seq[i : i+k]
			kmerRC := rc[len(rc)-i-k : len(rc)-i]
			if bytes.Compare(kmer, kmerRC) == 1 {
				kmer = kmerRC
			}
			if !yield(kmer) {
				return
			}
		}
	}
}
