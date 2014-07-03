// Creates indexes on fasta files.
package myindex

// Generates the index and searches in it.

import (
	"fmt"
	"errors"
	"seqtools"
	"bioformats/fasta"
)

// An index for a fasta file.
type Index struct {
	kmerLength int    // Length of indexed k-mers
	kmerInterval int  // Interval between indexed k-mers
	index [][]GenPos  // At index i are the positions of the i'th k-mer
}

// Builds an index for the given fasta. K-mer length and interval must be at
// least 1. Skips characters that are not in "acgtACGT".
func New(fa fasta.Fasta, kmerLength int, kmerInterval int) (*Index, error) {
	// Check input
	if kmerLength < 1 {
		panic( fmt.Sprint("bad kmer length:", kmerLength) )
	}
	
	if len(fa) > maxChr+1 {
		return nil, errors.New(fmt.Sprintf("too many chromosomes: %d (max=%d)",
				len(fa), maxChr+1))
	}
	
	// Create index
	index := make([][]GenPos, numOfKmers(kmerLength))
	
	// Go over chromosomes
	for chr := range fa {
		// Current sequence
		sequence   := fa[chr].Sequence
		numOfWords := len(sequence) - kmerLength + 1
		
		// If sequence is too long
		if len(sequence) > maxPos+1 {
			return nil, errors.New(
					fmt.Sprintf("chromosome %d is too long: %d (max=%d)",
					chr, len(sequence), maxChr+1))
		}
		
		// Go over sequence
		var kmer int
		var lastNonNucleotide int = -1
		numberOfKmers := numOfKmers(kmerLength)
		for pos := 0; pos < numOfWords; pos++ {
		// for pos := 0; pos < numOfWords; pos += kmerInterval {
			// Update kmer index
			if pos == 0 {
				kmer = kmerIndex(sequence[:kmerLength])
				
				// If met a non-nucleotide, find its index
				if kmer == -1 {
					kmer = 0
					for pos2 := 0; pos2 < kmerLength; pos2++ {
						if !isNucleotide(sequence[pos2]) {
							lastNonNucleotide = pos2
							// Don't break, there may be more ahead
						}
					}
				}
			} else {
				kmer = 4*kmer + nt2int[sequence[pos + kmerLength - 1]]
				kmer %= numberOfKmers
				if !isNucleotide(sequence[pos + kmerLength - 1]) {
					lastNonNucleotide = pos + kmerLength - 1
				}
			}
			
			// Add position to index?
			if pos % kmerInterval == 0 && lastNonNucleotide < pos {
				index[kmer] = append(index[kmer], NewGenPos(chr, pos, Plus))
			}
		}
	}
	
	return &Index{kmerLength, kmerInterval, index}, nil
}

// String representation of an index (for debugging).
func (idx *Index) String() string {
	// Count elements
	elements := 0
	for i := range idx.index {
		elements += len(idx.index[i])
	}
	
	return fmt.Sprintf("index:\n\tkmer length\t\t%d" +
			"\n\tsampling interval\t%d" +
			"\n\tnumber of kmers\t\t%d" +
			"\n\tnumber of records\t%d" +
			"\n\tavg. records per kmer\t%.1f",
			idx.kmerLength, idx.kmerInterval, numOfKmers(idx.kmerLength),
			elements, float64(elements) / float64(numOfKmers(idx.kmerLength)))
}

// Searches for the given sequence. Returns a map of genomic positions
// and the number of matching k-mers with the query sequence.
// If 'reverseComplement' is set to true, it will also search for the reverse
// complement of the given sequence.
// 'kmerInterval' denotes the k-mer sampling interval on the query sequence.
// If set to 1, will check for all the k-mers on the query. Must be at least
// 1.
// The function skips characters that are not in "acgtACGT".
func (idx *Index) Search(query []byte, kmerInterval int,
		reverseComplement bool) map[GenPos]int {
	// Check input
	if kmerInterval < 1 {
		panic(fmt.Sprint("bad k-mer interval: ", kmerInterval))
	}
	
	// This map counts the number of matches for each candidate position.
	posmap := map[GenPos]int{}
	
	// Break into k-mers
	numOfKmers := len(query) - idx.kmerLength + 1
	for i := 0; i < numOfKmers; i += kmerInterval {
		kmer := kmerIndex(query[i : i + idx.kmerLength])
		
		// Bad k-mer
		if kmer == -1 {
			continue
		}
		
		// Look up in index
		for _,amit := range idx.index[kmer] {
			// BUG( ) TODO find a better name for 'amit'.
			
			// If points to a negative position
			if amit.Pos() - i < 0 {
				continue
			}
			
			candidate := NewGenPos(amit.Chr(), amit.Pos() - i, Plus)
			posmap[candidate]++
		}
	}
	
	// Search for reverse complement too
	if reverseComplement {
		rc := seqtools.ReverseComplement(query)
		for i := 0; i < numOfKmers; i += kmerInterval {
			kmer := kmerIndex(rc[i : i + idx.kmerLength])
		
			// Bad k-mer
			if kmer == -1 {
				continue
			}
		
			// Look up in index
			for _,amit := range idx.index[kmer] {
				// BUG( ) TODO find a better name for 'amit'. Again.
				
				// If points to a negative position
				if amit.Pos() - i < 0 {
					continue
				}
				
				candidate := NewGenPos(amit.Chr(), amit.Pos() - i, Minus)
				posmap[candidate]++
			}
		}
	}
	
	return posmap
}





