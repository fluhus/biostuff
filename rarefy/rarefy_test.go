package rarefy

import (
	"fmt"
	"math"
	"slices"
	"testing"

	"github.com/fluhus/gostuff/gnum"
	"github.com/fluhus/gostuff/snm"
)

func TestChunkShuffle(t *testing.T) {
	input := snm.Slice(100, func(i int) int { return i })
	want := snm.TightClone(input)

	for range 10 {
		got := chunkShuffle(input)
		if !slices.Equal(input, want) {
			t.Fatalf("chunkShuffle(%v): input changed to %v", want, input)
		}
		if slices.Equal(got, input) {
			t.Fatalf("chunkShuffle(%v): not shuffled", want)
		}
		if !slices.Equal(snm.Sorted(got), want) {
			t.Fatalf("chunkShuffle(%v)=%v", want, got)
		}
	}
}

func BenchmarkChunkShuffle(b *testing.B) {
	for _, n := range []int{1000, 10000, 100000, 1000000, 10000000} {
		b.Run(fmt.Sprint("chunkShuffle,", n), func(b *testing.B) {
			a := snm.Slice(n, func(i int) int { return i })
			for b.Loop() {
				chunkShuffle(a)
			}
		})
		b.Run(fmt.Sprint("snm.Shuffle,", n), func(b *testing.B) {
			a := snm.Slice(n, func(i int) int { return i })
			for b.Loop() {
				snm.Shuffle(a)
			}
		})
	}
}

func TestSteps(t *testing.T) {
	tests := []struct {
		sum, step int
		want      []int
	}{
		{50, 10, []int{10, 20, 30, 40, 50}},
		{55, 10, []int{10, 20, 30, 40, 50, 55}},
		{15, 10, []int{10, 15}},
		{10, 10, []int{10}},
		{7, 10, []int{7}},
	}
	for _, test := range tests {
		got := slices.Collect(steps(test.sum, test.step))
		if !slices.Equal(got, test.want) {
			t.Errorf("steps(%v,%v)=%v, want %v",
				test.sum, test.step, got, test.want)
		}
	}
}

func TestLogFactorial(t *testing.T) {
	tests := [][]int{
		{0, 1}, {1, 1}, {2, 2}, {3, 6}, {4, 24}, {5, 120}, {6, 720},
	}
	for _, test := range tests {
		got := math.Exp(logFactorial(test[0]))
		want := float64(test[1])
		if gnum.Diff(got, want) > want*0.1 {
			t.Errorf("lf(%v)=%v, want %v", test[0], got, test[1])
		}
	}
}
