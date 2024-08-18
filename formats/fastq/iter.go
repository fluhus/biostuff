package fastq

import (
	"io"
	"iter"

	"github.com/fluhus/gostuff/aio"
)

// Iter returns an iterator over fastq entries.
func (r *Reader) Iter() iter.Seq2[*Fastq, error] {
	return func(yield func(*Fastq, error) bool) {
		for {
			fq, err := r.Read()
			if err != nil {
				if err != io.EOF {
					yield(nil, err)
				}
				break
			}
			if !yield(fq, nil) {
				return
			}
		}
	}
}

// IterFile returns an iterator over fastq entries in a file.
func IterFile(file string) iter.Seq2[*Fastq, error] {
	return func(yield func(*Fastq, error) bool) {
		f, err := aio.Open(file)
		if err != nil {
			yield(nil, err)
			return
		}
		defer f.Close()
		for fq, err := range NewReader(f).Iter() {
			if !yield(fq, err) {
				break
			}
		}
	}
}
