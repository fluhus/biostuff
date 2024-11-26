package sam

import (
	"encoding/csv"
	"io"
	"iter"
	"strings"

	"github.com/fluhus/gostuff/aio"
	"github.com/fluhus/gostuff/iterx"
)

func ReaderHeader(r io.Reader) iter.Seq2[SAMOrHeader, error] {
	return func(yield func(SAMOrHeader, error) bool) {
		csvReader := iterx.CSVReader(r, func(r *csv.Reader) {
			r.Comma = '\t'
			r.FieldsPerRecord = -1 // Allow variable number of fields.
			r.LazyQuotes = true
		})
		for line, err := range csvReader {
			// Error case.
			if err != nil {
				if !yield(SAMOrHeader{}, err) {
					break
				}
				continue
			}
			// Header line case.
			if len(line) > 0 && strings.HasPrefix(line[0], "@") {
				h := strings.Join(line, "\t")
				if !yield(SAMOrHeader{H: &h}, nil) {
					break
				}
				continue
			}
			// SAM line case.
			s, err := parseLine(line)
			if !yield(SAMOrHeader{S: s}, err) {
				break
			}
		}
	}
}

// Reader returns an iterator over SAM entries in a reader.
func Reader(r io.Reader) iter.Seq2[*SAM, error] {
	return func(yield func(*SAM, error) bool) {
		for sh, err := range ReaderHeader(r) {
			if err != nil {
				if !yield(nil, err) {
					break
				}
				continue
			}
			if sh.S == nil {
				continue
			}
			if !yield(sh.S, nil) {
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
		for sm, err := range Reader(f) {
			if !yield(sm, err) {
				break
			}
		}
	}
}

func FileHeader(file string) iter.Seq2[SAMOrHeader, error] {
	return func(yield func(SAMOrHeader, error) bool) {
		f, err := aio.Open(file)
		if err != nil {
			yield(SAMOrHeader{}, err)
			return
		}
		defer f.Close()
		for sh, err := range ReaderHeader(f) {
			if !yield(sh, err) {
				break
			}
		}
	}
}

type SAMOrHeader struct {
	H *string
	S *SAM
}
