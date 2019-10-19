// Genetic sequence processing functions.
package seqtools

import (
	"fmt"
	"math/rand"
)

// Converts a nucleotide to an int.
// Returns -1 for unknown nucleotides.
func Ntoi(nuc byte) int {
	switch nuc {
	case 'A', 'a':
		return 0
	case 'C', 'c':
		return 1
	case 'G', 'g':
		return 2
	case 'T', 't':
		return 3
	default:
		return -1
	}
}

// Converts an int to a nucleotide character.
// Returns 'N' for any value not from {0,1,2,3}.
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

// Returns the n-gram count vector for the given sequence.
func NgramVector(n int, sequence []byte) []int {
	// Check input
	if n < 1 || n > 10 {
		panic(fmt.Sprintf("Bad n: %d", n))
	}

	// Calculate size of result
	rsize := 1
	for i := 0; i < n; i++ {
		rsize *= 4 // For 4 letters in DNA
	}

	// Initialize result
	result := make([]int, rsize)

	// Count n-grams
	ngram := 0
	lastBad := n // Distance to last bad nucleotide
	for i := range sequence {
		// Nucleotide index
		ntoi := Ntoi(sequence[i])

		// Check if bad
		if ntoi == -1 {
			lastBad = 0
			ntoi = 0
		}

		// n-gram index
		ngram = (ngram*4)%rsize + Ntoi(sequence[i])

		// Increment only if went over a whole n-gram and no bad nucleotide
		if i >= (n-1) && lastBad >= n {
			result[ngram]++
		}

		lastBad++
	}

	return result
}

// Returns a random (uniform) nucleotide.
func RandNuc() byte {
	return Iton(rand.Intn(4))
}

// Generates a random (uniform) DNA sequence.
func RandSeq(length int) []byte {
	// Create result array
	result := make([]byte, length)

	// Randomize characters
	for i := range result {
		result[i] = RandNuc()
	}

	return result
}

// Returns a copy of the sequence, with exactly n SNPs.
func MutateSNP(sequence []byte, n int) (mutant []byte) {
	// Check input
	if n > len(sequence) {
		panic(fmt.Sprintf("n is greater than sequence length: %d > %d",
			n, len(sequence)))
	}
	if n < 0 {
		panic(fmt.Sprintf("n cannot be negative: %d", n))
	}

	// Create result slice
	mutant = make([]byte, len(sequence))
	copy(mutant, sequence)

	// Pick positions to mutate
	positions := rand.Perm(len(sequence))[0:n]

	// Mutate
	for _, i := range positions {
		// Generate random nucleotide that's different from current
		nuc := RandNuc()
		for nuc == mutant[i] {
			nuc = RandNuc()
		}

		mutant[i] = nuc
	}

	return
}

// Returns a new sequence with an insertion with the given size, at a random
// position along the sequence.
func MutateIns(sequence []byte, size int) (mutant []byte) {
	// Check input
	if size < 0 {
		panic(fmt.Sprintf("bad size: %d", size))
	}

	// Create result slice
	mutant = make([]byte, len(sequence)+size)

	// Pick a random position
	pos := rand.Intn(len(sequence) + 1)

	// Copy pre-mutation bytes
	copy(mutant[:pos], sequence)

	// Copy post-mutation bytes
	copy(mutant[pos+size:], sequence[pos:])

	// Generate insertion
	for i := pos; i < (pos + size); i++ {
		mutant[i] = RandNuc()
	}

	return
}

// Deletes a random subsequence of the given size.
func MutateDel(sequence []byte, size int) (mutant []byte) {
	// Check input
	if size > len(sequence) {
		panic(fmt.Sprintf("Size (%d) is greater than sequence length(%d)",
			size, len(sequence)))
	}

	if size < 0 {
		panic(fmt.Sprintf("Bad size: %d", size))
	}

	// Create result slice
	mutant = make([]byte, len(sequence)-size)

	// Pick a random position
	pos := rand.Intn(len(sequence) - size + 1)

	// Copy pre-deletion bytes
	copy(mutant[:pos], sequence)

	// Copy post-deletion bytes
	copy(mutant[pos:], sequence[pos+size:])

	return
}

// Returns the reverse complement of the given sequence. Only complements
// characters in "acgtACGT", other characters remain the same.
func ReverseComplement(sequence []byte) []byte {
	result := make([]byte, len(sequence))

	// Complement
	for i := range sequence {
		switch sequence[i] {
		case 'a':
			result[i] = 't'
		case 'c':
			result[i] = 'g'
		case 'g':
			result[i] = 'c'
		case 't':
			result[i] = 'a'
		case 'A':
			result[i] = 'T'
		case 'C':
			result[i] = 'G'
		case 'G':
			result[i] = 'C'
		case 'T':
			result[i] = 'A'
		default:
			result[i] = sequence[i]
		}
	}

	// Reverse
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] =
			result[len(result)-i-1], result[i]
	}

	return result
}

// Returns the reverse complement of the given sequence. Only complements
// characters in "acgtACGT", other characters remain the same.
func ReverseComplementString(sequence string) string {
	return string(ReverseComplement([]byte(sequence)))
}