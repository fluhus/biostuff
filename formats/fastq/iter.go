package fastq

import (
	"fmt"
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

// FilePaired returns an iterator over fastq entries in two paired files.
// Yields pairs of entries, one from each file.
// The files are expected to have the same number of reads.
func FilePaired(file1, file2 string) iter.Seq2[[]*Fastq, error] {
	return func(yield func([]*Fastq, error) bool) {
		next1, stop1 := iter.Pull2(File(file1))
		next2, stop2 := iter.Pull2(File(file2))
		defer stop1()
		defer stop2()

		for {
			fq1, err1, ok1 := next1()
			fq2, err2, ok2 := next2()
			if ok1 != ok2 { // One ran out before the other.
				if ok1 {
					yield(nil, fmt.Errorf("input 2 ran out before input 1"))
				} else {
					yield(nil, fmt.Errorf("input 1 ran out before input 2"))
				}
				return
			}
			if !ok1 { // All done.
				return
			}
			if err1 != nil {
				yield(nil, fmt.Errorf("input 1: %w", err1))
				return
			}
			if err2 != nil {
				yield(nil, fmt.Errorf("input 2: %w", err2))
				return
			}
			if !yield([]*Fastq{fq1, fq2}, nil) {
				return
			}
		}
	}
}
