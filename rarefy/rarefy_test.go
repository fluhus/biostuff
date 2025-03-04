package rarefy

import (
	"fmt"
	"slices"
	"testing"

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
