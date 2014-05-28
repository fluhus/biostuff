package myindex

// A type for genomic positions.

import (
	"fmt"
)

// *** CONSTANTS **************************************************************

// Number of bits for chromosome number
const bitsChr = 16

// Number of bits for position
const bitsPos = 47

// Number of bits for strand
const bitsStrand = 1

// Offset of chromosome number
const offsetChr = bitsPos + bitsStrand

// Offset of position
const offsetPos = bitsStrand

// Offset of strand
const offsetStrand = 0

// Max value for chromosome.
const maxChr = 1 << bitsChr - 1

// Max value for position.
const maxPos = 1 << bitsPos - 1

// Max value for strand.
const maxStrand = 1 << bitsStrand - 1

// The bits that represent the chromosome
const maskChr uint64 = maxChr << offsetChr

// The bits that represent the position
const maskPos uint64 = maxPos << offsetPos

// The bits that represent the strand
const maskStrand uint64 = maxStrand << offsetStrand


// *** STRAND TYPE ************************************************************

// Represents the strand number.
type StrandType int

const (
	Plus  StrandType = 0
	Minus StrandType = 1
)


// *** TYPE *******************************************************************

// Holds a genomic position in a single integer.
type GenPos uint64

// Creates a new position variable (not on heap).
func NewGenPos(chr int, pos int, strand StrandType) GenPos {
	var g GenPos
	g.SetChr(chr)
	g.SetPos(pos)
	g.SetStrand(strand)
	return g
}

// Returns the chromosome number.
func (g GenPos) Chr() int {
	return int((uint64(g) & maskChr) >> offsetChr)
}

// Returns the position inside the chromosome.
func (g GenPos) Pos() int {
	return int((uint64(g) & maskPos) >> offsetPos)
}

// Returns the strand.
func (g GenPos) Strand() StrandType {
	return StrandType((uint64(g) & maskStrand) >> offsetStrand)
}

// Sets the chromosome number.
func (g *GenPos) SetChr(chr int) {
	// Check input
	if chr < 0 || chr > maxChr {
		panic(fmt.Sprintf("bad chromosome number: %d (0 <= n <= %d)",
				chr, maxChr))
	}
	
	*g = GenPos( (uint64(*g) & ^maskChr) | (uint64(chr) << offsetChr) )
}

// Sets the position.
func (g *GenPos) SetPos(pos int) {
	// Check input
	if pos < 0 || pos > maxPos {
		panic(fmt.Sprintf("bad position: %d (0 <= n <= %d)",
				pos, maxPos))
	}
	
	*g = GenPos( (uint64(*g) & ^maskPos) | (uint64(pos) << offsetPos) )
}

// Sets the strand.
func (g *GenPos) SetStrand(strand StrandType) {
	// Check input
	if strand < 0 || strand > maxStrand {
		panic(fmt.Sprintf("bad strand: %d (0 <= n <= %d)",
				strand, maxStrand))
	}
	
	*g = GenPos( (uint64(*g) & ^maskStrand) | (uint64(strand) << offsetStrand) )
}

// String representation of the position.
func (g GenPos) String() string {
	return fmt.Sprintf("(%d,%d,%d)", g.Chr(), g.Pos(), g.Strand())
}
