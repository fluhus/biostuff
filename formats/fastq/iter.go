package fastq

import (
	"io"
	"iter"

	"github.com/fluhus/gostuff/aio"
)

// Returns an iterator over fastq entries.
func (r *reader) iter() iter.Seq2[*Fastq, error] {
	return func(yield func(*Fastq, error) bool) {
		for {
			fq, err := r.read()
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

// File returns an iterator over fastq entries in a file.
func File(file string) iter.Seq2[*Fastq, error] {
	return func(yield func(*Fastq, error) bool) {
		f, err := aio.Open(file)
		if err != nil {
			yield(nil, err)
			return
		}
		defer f.Close()
		for fq, err := range Reader(f) {
			if !yield(fq, err) {
				break
			}
		}
	}
}

// Reader returns an iterator over fastq entries in a reader.
func Reader(r io.Reader) iter.Seq2[*Fastq, error] {
	return func(yield func(*Fastq, error) bool) {
		for fa, err := range newReader(r).iter() {
			if !yield(fa, err) {
				break
			}
		}
	}
}
