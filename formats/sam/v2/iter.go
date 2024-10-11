package sam

import (
	"io"
	"iter"

	"github.com/fluhus/gostuff/aio"
)

// Returns an iterator over SAM entries.
func (r *reader) iter() iter.Seq2[*SAM, error] {
	return func(yield func(*SAM, error) bool) {
		for {
			sm, err := r.read()
			if err != nil {
				if err != io.EOF {
					yield(nil, err)
				}
				break
			}
			if !yield(sm, nil) {
				return
			}
		}
	}
}

// Reader returns an iterator over SAM entries in a reader.
func Reader(r io.Reader) iter.Seq2[*SAM, error] {
	return func(yield func(*SAM, error) bool) {
		for sm, err := range newReader(r).iter() {
			if !yield(sm, err) {
				break
			}
		}
	}
}

// File returns an iterator over SAM entries in a file.
func File(file string) iter.Seq2[*SAM, error] {
	return func(yield func(*SAM, error) bool) {
		f, err := aio.Open(file)
		if err != nil {
			yield(nil, err)
			return
		}
		defer f.Close()
		for sm, err := range newReader(f).iter() {
			if !yield(sm, err) {
				break
			}
		}
	}
}
