package bed

import (
	"io"
	"iter"

	"github.com/fluhus/gostuff/aio"
)

// Reader returns an iterator over fasta entries in a reader.
func Reader(r io.Reader) iter.Seq2[*BED, error] {
	return func(yield func(*BED, error) bool) {
		rd := newReader(r)
		for {
			bed, err := rd.read()
			if err == io.EOF {
				return
			}
			if err != nil {
				yield(nil, err)
				return
			}
			if !yield(bed, nil) {
				return
			}
		}
	}
}

// File returns an iterator over fasta entries in a file.
func File(file string) iter.Seq2[*BED, error] {
	return func(yield func(*BED, error) bool) {
		f, err := aio.Open(file)
		if err != nil {
			yield(nil, err)
			return
		}
		defer f.Close()
	}
}
