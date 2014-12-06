package main

import (
	"testing"
	"bioformats/fastq"
)

func Test_NoMatch(t *testing.T) {
	fq := &fastq.Fastq{}
	fq.Sequence = []byte("ACTAGGTTCA")
	fq.Quals = []byte("IIIIIIIIII")
	initialLength := len(fq.Sequence)
	
	trimAdapterEnd(fq, []byte("CCCA"))
	
	if len(fq.Sequence) != initialLength {
		t.Errorf("expected no trim, got length %d instead of %d",
				len(fq.Sequence), initialLength)
	}
}
