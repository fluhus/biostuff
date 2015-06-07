package myindex

// A unit test for GenPos.

import (
	"fmt"
	"testing"
)

func Test_Zeros(t *testing.T) {
	g := NewGenPos(0, 0, Plus)
	
	if g.Chr() != 0 {
		t.Error(fmt.Sprintf("error in Chr(): expected 0 got %d", g.Chr()))
	}
	if g.Pos() != 0 {
		t.Error(fmt.Sprintf("error in Pos(): expected 0 got %d", g.Pos()))
	}
	if g.Strand() != 0 {
		t.Error(fmt.Sprintf("error in Strand(): expected 0 got %d", g.Strand()))
	}
}

func Test_Ones(t *testing.T) {
	g := NewGenPos(1, 1, 1)
	
	if g.Chr() != 1 {
		t.Error(fmt.Sprintf("error in Chr(): expected 1 got %d", g.Chr()))
	}
	if g.Pos() != 1 {
		t.Error(fmt.Sprintf("error in Pos(): expected 1 got %d", g.Pos()))
	}
	if g.Strand() != 1 {
		t.Error(fmt.Sprintf("error in Strand(): expected 1 got %d", g.Strand()))
	}
}

func Test_Maxs(t *testing.T) {
	g := NewGenPos(maxChr, maxPos, maxStrand)
	
	if g.Chr() != maxChr {
		t.Error(fmt.Sprintf("error in Chr(): expected %d got %d",
				maxChr, g.Chr()))
	}
	if g.Pos() != maxPos {
		t.Error(fmt.Sprintf("error in Pos(): expected %d got %d",
				maxPos, g.Pos()))
	}
	if g.Strand() != maxStrand {
		t.Error(fmt.Sprintf("error in Strand(): expected %d got %d",
				maxStrand, g.Strand()))
	}
}

func Test_Set(t *testing.T) {
	g := NewGenPos(maxChr, maxPos, maxStrand)
	
	g.SetChr(0)
	g.SetPos(0)
	g.SetStrand(0)
	
	if g.Chr() != 0 {
		t.Error(fmt.Sprintf("error in Chr(): expected 0 got %d", g.Chr()))
	}
	if g.Pos() != 0 {
		t.Error(fmt.Sprintf("error in Pos(): expected 0 got %d", g.Pos()))
	}
	if g.Strand() != 0 {
		t.Error(fmt.Sprintf("error in Strand(): expected 0 got %d", g.Strand()))
	}
	
	g.SetChr(maxChr)
	g.SetPos(maxPos)
	g.SetStrand(maxStrand)
	
	if g.Chr() != maxChr {
		t.Error(fmt.Sprintf("error in Chr(): expected %d got %d",
				maxChr, g.Chr()))
	}
	if g.Pos() != maxPos {
		t.Error(fmt.Sprintf("error in Pos(): expected %d got %d",
				maxPos, g.Pos()))
	}
	if g.Strand() != maxStrand {
		t.Error(fmt.Sprintf("error in Strand(): expected %d got %d",
				maxStrand, g.Strand()))
	}
}
