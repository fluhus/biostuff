package sam

import (
	"encoding/csv"
	"io"
	"iter"
	"strings"

	"github.com/fluhus/gostuff/aio"
	"github.com/fluhus/gostuff/iterx"
)

// ReaderHeader iterates over SAM or header entries in a reader.
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

// Reader iterates over SAM entries in a reader.
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

// File iterates over SAM entries in a file.
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

// FileHeader iterates over SAM or header entries in a file.
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

// SAMOrHeader holds either a SAM or a header entry from a SAM-formatted
// input.
// If there is no error, exactly one of the fields will be non-nil.
type SAMOrHeader struct {
	H *string // Header line, including the '@' sign.
	S *SAM    // SAM entry.
}
