//go:build go1.22

package sam

import (
	"io"
	"iter"

	"github.com/fluhus/gostuff/aio"
)

// Iter returns an iterator over SAM entries.
func (r *Reader) Iter() iter.Seq2[*SAM, error] {
	return func(yield func(*SAM, error) bool) {
		for {
			sm, err := r.Read()
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

// IterFile returns an iterator over SAM entries in a file.
func IterFile(file string) iter.Seq2[*SAM, error] {
	return func(yield func(*SAM, error) bool) {
		f, err := aio.Open(file)
		if err != nil {
			yield(nil, err)
			return
		}
		defer f.Close()
		r := NewReader(f)
		for sm, err := range r.Iter() {
			if !yield(sm, err) {
				break
			}
		}
	}
}
