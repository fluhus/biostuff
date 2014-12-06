package main

import (
	"testing"
	"bioformats/fastq"
)

func Test_NoMatch(t *testing.T) {
	fq := &fastq.Fastq{}
	fq.Sequence = []byte("ACTAGGTTCA")
	fq.Quals = []byte("IIIIIIIIII")
	
	trimAdapterEnd(fq, []byte("CCCA"), 5)
	
	if string(fq.Sequence) != "ACTAGGTTCA" {
		t.Errorf("bad trimming: got '%s' expected '%s'",
				string(fq.Sequence), "ACTAGGTTCA")
	}
}

func Test_Match(t *testing.T) {
	fq := &fastq.Fastq{}
	fq.Sequence = []byte("ACTAGGTTCA")
	fq.Quals = []byte("ABCDEFGHIJ")
	
	trimAdapterEnd(fq, []byte("TCAAAAAA"), 5)
	
	if string(fq.Sequence) != "ACTAGGT" {
		t.Errorf("bad trimming: got '%s' expected '%s'",
				string(fq.Sequence), "ACTAGGT")
	}
	if string(fq.Quals) != "ABCDEFG" {
		t.Errorf("bad trimming: got quals '%s' expected '%s'",
				string(fq.Quals), "ABCDEFG")
	}
}

func Test_MatchWithErrors(t *testing.T) {
	fq := &fastq.Fastq{}
	fq.Sequence = []byte("ACTAGGTTCATTTAGCGCTTAA")
	fq.Quals =    []byte("1234567891234567891234")
	
	trimAdapterEnd(fq, []byte("TAGCCCTTAGGTAAT"), 5)
	
	if string(fq.Sequence) != "ACTAGGTTCATT" {
		t.Errorf("bad trimming: got '%s' expected '%s'",
				string(fq.Sequence), "ACTAGGTTCATT")
	}
	if string(fq.Quals) != "123456789123" {
		t.Errorf("bad trimming: got quals '%s' expected '%s'",
				string(fq.Quals), "123456789123")
	}
}

func Test_MatchWithTooManyErrors(t *testing.T) {
	fq := &fastq.Fastq{}
	fq.Sequence = []byte("ACTAGGTTCATTTAGCGCTTAA")
	fq.Quals =    []byte("1234567891234567891234")
	
	trimAdapterEnd(fq, []byte("GAGCCCTTAGGTAAT"), 5)
	
	if string(fq.Sequence) != "ACTAGGTTCATTTAGCGCTTAA" {
		t.Errorf("bad trimming: got '%s' expected '%s'",
				string(fq.Sequence), "ACTAGGTTCATTTAGCGCTTAA")
	}
	if string(fq.Quals) != "1234567891234567891234" {
		t.Errorf("bad trimming: got quals '%s' expected '%s'",
				string(fq.Quals), "1234567891234567891234")
	}
}

func Test_TrimStart(t *testing.T) {
	fq := &fastq.Fastq{}
	fq.Sequence = []byte("ACTAGGTTCATTTAGCGCTTAA")
	fq.Quals =    []byte("1234567891234567891234")
	
	trimAdapterStart(fq, []byte("GGCACTA"), 5)
	
	if string(fq.Sequence) != "GGTTCATTTAGCGCTTAA" {
		t.Errorf("bad trimming: got '%s' expected '%s'",
				string(fq.Sequence), "GGTTCATTTAGCGCTTAA")
	}
	if string(fq.Quals) != "567891234567891234" {
		t.Errorf("bad trimming: got quals '%s' expected '%s'",
				string(fq.Quals), "567891234567891234")
	}
}

