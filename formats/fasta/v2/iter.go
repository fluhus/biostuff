package fasta

import (
	"io"
	"iter"

	"github.com/fluhus/gostuff/aio"
)

// Returns an iterator over fasta entries.
func (r *reader) iter() iter.Seq2[*Fasta, error] {
	return func(yield func(*Fasta, error) bool) {
		for {
			fa, err := r.read()
			if err != nil {
				if err != io.EOF {
					yield(nil, err)
				}
				break
			}
			if !yield(fa, nil) {
				return
			}
		}
	}
}

// File returns an iterator over fasta entries in a file.
func File(file string) iter.Seq2[*Fasta, error] {
	return func(yield func(*Fasta, error) bool) {
		f, err := aio.Open(file)
		if err != nil {
			yield(nil, err)
			return
		}
		defer f.Close()
		for fa, err := range Reader(f) {
			if !yield(fa, err) {
				break
			}
		}
	}
}

// Reader returns an iterator over fasta entries in a reader.
func Reader(r io.Reader) iter.Seq2[*Fasta, error] {
	return func(yield func(*Fasta, error) bool) {
		for fa, err := range newReader(r).iter() {
			if !yield(fa, err) {
				break
			}
		}
	}
}
