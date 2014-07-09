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
		
		// If sequence is too long
		if uint64(len(sequence)) > maxPos+1 {
			return nil, errors.New(
					fmt.Sprintf("chromosome %d is too long: %d (max=%d)",
					chr, len(sequence), maxChr+1))
		}
		
		// Go over sequence
		it := newKmerIterator(sequence, kmerLength)
		pos := 0
		for it.hasNext() {
			kmer := it.next()
			
			// Add position to index?
			if pos % kmerInterval == 0 && kmer != -1 {
				index[kmer] = append(index[kmer], NewGenPos(chr, pos, Plus))
			}
			
			pos++
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
	it := newKmerIterator(query, idx.kmerLength)
	pos := 0
	for it.hasNext() {
		kmer := it.next()
		
		if kmer == -1 || pos % kmerInterval != 0 {
			continue
		}
		
		// Look up in index
		for _,amit := range idx.index[kmer] {
			// BUG( ) TODO find a better name for 'amit'.
			
			// If points to a negative position
			if amit.Pos() - pos < 0 {
				continue
			}
			
			candidate := NewGenPos(amit.Chr(), amit.Pos() - pos, Plus)
			posmap[candidate]++
		}
		
		pos++
	}
	
	// Search for reverse complement too
	if reverseComplement {
		rc := seqtools.ReverseComplement(query)
		it = newKmerIterator(rc, idx.kmerLength)
		pos = 0
		for it.hasNext() {
			kmer := it.next()
		
			if kmer == -1 || pos % kmerInterval != 0 {
				continue
			}
		
			// Look up in index
			for _,amit := range idx.index[kmer] {
				// BUG( ) TODO find a better name for 'amit'. Again.
				
				// If points to a negative position
				if amit.Pos() - pos < 0 {
					continue
				}
				
				candidate := NewGenPos(amit.Chr(), amit.Pos() - pos, Minus)
				posmap[candidate]++
			}
			
			pos++
		}
	}
	
	return posmap
}





