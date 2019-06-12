package main

import (
	"testing"

	"github.com/fluhus/biostuff/bioformats/fastq"
)

func Test_Simple(t *testing.T) {
	fq := &fastq.Fastq{}
	fq.Sequence = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fq.Quals = []byte{10, 9, 11, 8, 15, 16, 9, 10, 9}

	trimQual(fq, 0, 10)

	newSequence := []byte{5, 6}
	newQuals := []byte{15, 16}
	if string(fq.Sequence) != string(newSequence) {
		t.Errorf("bad sequence: %v, expected: %v", fq.Sequence, newSequence)
	}
	if string(fq.Quals) != string(newQuals) {
		t.Errorf("bad qualities: %v, expected: %v", fq.Quals, newQuals)
	}
}

func Test_BadQuals(t *testing.T) {
	fq := &fastq.Fastq{}
	fq.Sequence = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fq.Quals = []byte{7, 9, 7, 8, 1, 10, 9, 10, 9}

	trimQual(fq, 0, 10)

	if len(fq.Sequence) > 0 {
		t.Errorf("bad sequence: %v, expected empty", fq.Sequence)
	}
	if len(fq.Quals) > 0 {
		t.Errorf("bad qualities: %v, expected empty", fq.Quals)
	}
}

func Test_Empty(t *testing.T) {
	fq := &fastq.Fastq{}

	trimQual(fq, 0, 10)

	if len(fq.Sequence) > 0 {
		t.Errorf("bad sequence: %v, expected empty", fq.Sequence)
	}
	if len(fq.Quals) > 0 {
		t.Errorf("bad qualities: %v, expected empty", fq.Quals)
	}
}
