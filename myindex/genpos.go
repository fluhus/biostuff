package myindex

// A type for genomic positions.

import "fmt"

// *** CONSTANTS **************************************************************

// Max value for chromosome.
const maxChr = 2 << 8 - 1

// Max value for position.
const maxPos = 2 << 56 - 1

// The bits that represent the chromosome.
const maskChr uint64 = 255 << 56

// The bits that represent the position.
const maskPos uint64 = ^maskChr

// *** TYPE *******************************************************************

// Holds a genomic position in a single integer.
// 8 bytes are used for the chromosome and 56 bytes for
// the position in the chromosome.
type GenPos uint64

// Creates a new position variable (not on heap).
func NewGenPos(chr, pos int) GenPos {
	// Check input
	if chr < 0 {
		panic(fmt.Sprint("bad chromosome index: ", chr, " (must be >= 0)"))
	}
	if pos < 0 {
		panic(fmt.Sprint("bad position: ", pos, " (must be >= 0)"))
	}
	
	var g GenPos
	g.SetChr(chr)
	g.SetPos(pos)
	return g
}

// Returns the chromosome number.
func (g GenPos) Chr() int {
	return int(g >> 56)
}

// Returns the position inside the chromosome.
func (g GenPos) Pos() int {
	return int( uint64(g) & maskPos )
}

// Sets the chromosome number. Max value is (2^8 - 1 = 255).
func (g *GenPos) SetChr(chr int) {
	*g = GenPos( (uint64(*g) & maskPos) + (uint64(chr) << 56) )
}

// Sets the position inside the chromosome. Max value is (2^56 - 1 = ???).
func (g *GenPos) SetPos(pos int) {
	*g = GenPos( (uint64(*g) & maskChr) + uint64(pos) )
}

// String representation of the position.
func (g GenPos) String() string {
	return fmt.Sprintf("(%d,%d)", g.Chr(), g.Pos())
}
