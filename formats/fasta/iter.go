//go:build go1.22

package fasta

import (
	"io"
	"iter"

	"github.com/fluhus/gostuff/aio"
)

// Iter returns an iterator over fasta entries.
func (r *Reader) Iter() iter.Seq2[*Fasta, error] {
	return func(yield func(*Fasta, error) bool) {
		for {
			fa, err := r.Read()
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

// IterFile returns an iterator over fasta entries in a file.
func IterFile(file string) iter.Seq2[*Fasta, error] {
	return func(yield func(*Fasta, error) bool) {
		f, err := aio.Open(file)
		if err != nil {
			yield(nil, err)
			return
		}
		defer f.Close()
		r := NewReader(f)
		for fa, err := range r.Iter() {
			if !yield(fa, err) {
				break
			}
		}
	}
}
