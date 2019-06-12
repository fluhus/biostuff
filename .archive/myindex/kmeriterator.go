package myindex

// Iterates over kmer numbers in a sequence.
type kmerIterator struct {
	sequence []byte    // Sequence on which to iterate
	k int              // Kmer length
	numberOfKmers int  // Number of k-long possible kmers
	position int       // Position of NEXT nucleotide
	lastResult int     // Value of last encountered kmer
	lastInvalid int    // Position of last encountered non-nucleotide
}
// BUG( ) TODO find a better name for 'lastResult'.

// Returns a new iterator.
func newKmerIterator(sequence []byte, k int) *kmerIterator {
	return &kmerIterator{sequence, k, numOfKmers(k), 0, 0, -1}
}

// Returns true iff there exists a next kmer.
func (it *kmerIterator) hasNext() bool {
	return it.position < len(it.sequence)
}

// Adds next nucleotide to the result and advances the counters.
func (it *kmerIterator) advance() {
	// Advance to next value
	it.lastResult = 4*it.lastResult + nt2int[it.sequence[it.position]]
	it.lastResult %= it.numberOfKmers
	
	// Update if invalid
	if !isNucleotide(it.sequence[it.position]) {
		it.lastInvalid = it.position
	}
	
	it.position++
}

func (it *kmerIterator) next() int {
	// Make sure there is next
	if !it.hasNext() {
		panic("invalid call, hasNext() is false")
	}
	
	// If first time, calculate initial kmer
	if it.position == 0 {
		for i := 0; i < it.k - 1; i++ {
			it.advance()
		}
	}
	
	it.advance()
	
	if it.lastInvalid < it.position - it.k {
		return it.lastResult
	} else {
		return -1
	}
}



