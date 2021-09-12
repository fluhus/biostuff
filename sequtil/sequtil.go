// Package sequtil provides functions for processing genetic sequences.
//
// Data Type For Sequences
//
// In this repository, sequences are represented as byte slices. This is meant
// to keep them familiar and predictable. This package provides a buffet of functions
// for manipulating sequences, rather than having a dedicated sequence type with
// methods.
//
// Mutating Sequences
//
// Mutations can be made using basic slice operations.
//  // Substitution
//  seq[4] = 'G'
//
//  // Deletion of bases 10-12 (including 12)
//  copy(seq[10:], seq[13:])
//  seq = seq[:len(seq)-3]
//
//  // Insertion at position 4
//  insert := []byte{...}
//  seq = append(append(append([]byte{}, seq[:4]...), insert...), seq[4:]...)
//
// The U Nucleotide
//
// This package currently ignores the existence of uracil. Adding support for uracil
// means increasing the complexity of the API without adding new capabilities. The
// current solution is to substitute U's with T's before calling this package.
package sequtil

import (
	"fmt"
)

// Maps nucleotide byte value to its int value.
var ntoi []int

func init() {
	// Initialize ntoi values.
	ntoi = make([]int, 256)
	for i := range ntoi {
		ntoi[i] = -1
	}
	ntoi['a'], ntoi['A'] = 0, 0
	ntoi['c'], ntoi['C'] = 1, 1
	ntoi['g'], ntoi['G'] = 2, 2
	ntoi['t'], ntoi['T'] = 3, 3
}

// Ntoi converts a nucleotide to an int.
// Returns -1 for unknown nucleotides.
func Ntoi(nuc byte) int {
	return ntoi[nuc]
}

// Iton converts an int to a nucleotide character.
// Returns 'N' for any value not in {0,1,2,3}.
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

// ReverseComplement writes to dst the reverse complement of src.
// Characters not in "aAcCgGtTnN" will cause a panic.
func ReverseComplement(dst, src []byte) {
	if len(dst) < len(src) {
		panic(fmt.Sprintf("dst is too short: %v, want at least %v",
			len(dst), len(src)))
	}

	// Complement
	for i, b := range src {
		switch b {
		case 'a':
			dst[len(src)-1-i] = 't'
		case 'c':
			dst[len(src)-1-i] = 'g'
		case 'g':
			dst[len(src)-1-i] = 'c'
		case 't':
			dst[len(src)-1-i] = 'a'
		case 'A':
			dst[len(src)-1-i] = 'T'
		case 'C':
			dst[len(src)-1-i] = 'G'
		case 'G':
			dst[len(src)-1-i] = 'C'
		case 'T':
			dst[len(src)-1-i] = 'A'
		case 'N':
			dst[len(src)-1-i] = 'N'
		case 'n':
			dst[len(src)-1-i] = 'n'
		default:
			panic(fmt.Sprintf("Unexpected base value: %v, want aAcCgGtTnN", b))
		}
	}
}

// ReverseComplementString returns the reverse complement of s.
// Characters not in "aAcCgGtTnN" will cause a panic.
func ReverseComplementString(s string) string {
	result := make([]byte, len(s))
	ReverseComplement(result, []byte(s))
	return string(result)
}

// DNATo2Bit writes to dst the 2-bit representation of the DNA sequence in src.
// Any character not in "aAcCgGtT" will cause a panic.
func DNATo2Bit(dst, src []byte) {
	if len(dst) < (len(src)+3)/4 {
		panic(fmt.Sprintf("dst is too short: %v, want at least %v",
			len(dst), (len(src)+3)/4))
	}
	for i, b := range src {
		di := i / 4
		shift := 6 - i%4*2 // Make the first character the most significant.
		if shift == 6 {
			// Reset byte value before or'ing.
			dst[di] = 0
		}
		dbInt := Ntoi(b)
		if dbInt == -1 {
			panic(fmt.Sprintf("Unexpected base value: %v, want aAcCgGtT", b))
		}
		db := byte(dbInt) << shift
		dst[di] |= db
	}
}

// DNAFrom2Bit writes to dst the nucleotides represented in 2-bit in src.
// Only outputs characters in "ACGT".
func DNAFrom2Bit(dst, src []byte) {
	if len(dst) < len(src)*4-3 {
		panic(fmt.Sprintf("dst is too short: %v, want at least %v",
			len(dst), len(src)*4-3))
	}
	for i := 0; i < len(src); i++ {
		copy(dst[i*4:], dnaFrom2bit[src[i]][:])
	}
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
