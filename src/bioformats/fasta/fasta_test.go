package fasta

import (
	"testing"
	"bufio"
	"strings"
	"github.com/stretchr/testify/assert"
)

func TestReadFastaEntry_simple(t *testing.T) {
	assert := assert.New(t)
	
	input := ">foo\nAaTtGngcCaN"

	fa, err := ReadFastaEntry( bufio.NewReader( strings.NewReader(input) ) )
	assert.Nil(err, "Error reading fasta: %v", err)

	assertFasta(assert, fa, "foo", "AaTtGngcCaN")
}

func TestReadFastaEntry_noName(t *testing.T) {
	assert := assert.New(t)
	
	input := "AaTtGngcCaN"

	fa, err := ReadFastaEntry( bufio.NewReader( strings.NewReader(input) ) )
	assert.Nil(err, "Error reading fasta: %v", err)

	assertFasta(assert, fa, "", input)
}

func TestReadFastaEntry_multiline(t *testing.T) {
	assert := assert.New(t)
	
	input := ">foo\nAaTtGngcCaN\nGGgg\n>foo"

	fa, err := ReadFastaEntry( bufio.NewReader( strings.NewReader(input) ) )
	assert.Nil(err, "Error reading fasta: %v", err)

	assertFasta(assert, fa, "foo", "AaTtGngcCaNGGgg")
}

func TestReadFastaEntry_badChars(t *testing.T) {
	assert := assert.New(t)
	
	input := ">foo\naaaaaaKaaaaa"

	_, err := ReadFastaEntry( bufio.NewReader( strings.NewReader(input) ) )
	assert.NotNil(err, "Expected error for non-nucleotide character.")
}

func TestReadFasta_simple(t *testing.T) {
	assert := assert.New(t)
	
	input := ">foo\nAaTtGngcCaN\nGGgg\n>bar\naaaGcgnnNcTAtgGa\nAA\n\nGagaGNtCc"

	fas, err := ReadFasta( strings.NewReader(input) )
	assert.Nil(err, "Error reading fasta: %v", err)

	assert.Equal(2, len(fas), "Bad number of entries read.")
	assertFasta(assert, fas[0], "foo", "AaTtGngcCaNGGgg")
	assertFasta(assert, fas[1], "bar", "aaaGcgnnNcTAtgGaAAGagaGNtCc")
}

func assertFasta(assert *assert.Assertions, fa *FastaEntry, name string,
		sequence string) {
	assert.Equal(fa.Name(), name, "Bad name.")

	assert.Equal(len(sequence), fa.Length(), "Bad fasta length.")

	usequence := strings.ToUpper(sequence)
	for i := range usequence {
		assert.Equal(usequence[i], fa.At(i), "Bad character at index %d.", i)
	}

	for i := range usequence {
		for l := 0; l < len(usequence) - i; l++ {
			assert.Equal(usequence[i : i + l], string(fa.Subsequence(i, l)),
					"Bad subsequence starting at %d length %d.", i, l)
		}
	}
}

