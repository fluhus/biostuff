package myindex

// A unit test for kmerIterator.

import (
	"fmt"
	"testing"
)

// Compares 2 int slices.
func intsEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	
	return true
}

func Test_Basic(t *testing.T) {
	seq := []byte("ACGT")
	it := newKmerIterator(seq, 1)
	
	var kmers []int
	for it.hasNext() {
		kmers = append(kmers, it.next())
	}
	
	if !intsEqual(kmers, []int{0,1,2,3}) {
		t.Error(fmt.Sprint("bad kmers:", kmers))
	}
}

func Test_Easy(t *testing.T) {
	seq := []byte("CACAC")
	it := newKmerIterator(seq, 2)
	
	var kmers []int
	for it.hasNext() {
		kmers = append(kmers, it.next())
	}
	
	if !intsEqual(kmers, []int{4,1,4,1}) {
		t.Error(fmt.Sprint("bad kmers:", kmers))
	}
}

func Test_Medium(t *testing.T) {
	seq := []byte("TAGTACvACC")
	//             302301#011
	it := newKmerIterator(seq, 3)
	
	var kmers []int
	for it.hasNext() {
		kmers = append(kmers, it.next())
	}
	
	if !intsEqual(kmers, []int{50,11,44,49,-1,-1,-1,5}) {
		t.Error(fmt.Sprint("bad kmers:", kmers))
	}
}
