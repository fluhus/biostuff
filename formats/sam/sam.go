// Package sam decodes and encodes SAM files.
//
// This package uses the format described in:
// https://en.wikipedia.org/wiki/SAM_(file_format)
package sam

import (
	"bytes"
	"fmt"
	"io"
	"strconv"

	"github.com/fluhus/gostuff/snm"
)

// SAM is a single line (alignment) in a SAM file.
type SAM struct {
	Qname string         // Query name
	Flag  Flag           // Bitwise flag
	Rname string         // Reference sequence name
	Pos   int            // Mapping position (1-based)
	Mapq  int            // Mapping quality
	Cigar string         // CIGAR string
	Rnext string         // Ref. name of the mate/next read
	Pnext int            // Position of the mate/next read
	Tlen  int            // Observed template length
	Seq   string         // Sequence
	Qual  string         // Phred qualities (ASCII)
	Tags  map[string]any // Typed optional tags.
}

// MarshalText returns the textual representation of s in SAM format.
// Includes a trailing new line.
func (s *SAM) MarshalText() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	s.Write(buf)
	return buf.Bytes(), nil
}

// Write writes this entry in textual SAM format to the given writer.
// Includes a trailing new line.
func (s *SAM) Write(w io.Writer) error {
	if _, err := fmt.Fprintf(w,
		"%s\t%d\t%s\t%d\t%d\t%s\t%s\t%d\t%d\t%s\t%s",
		s.Qname, s.Flag, s.Rname, s.Pos,
		s.Mapq, s.Cigar, s.Rnext, s.Pnext,
		s.Tlen, s.Seq, s.Qual,
	); err != nil {
		return err
	}
	for _, tag := range tagsToText(s.Tags) {
		if _, err := fmt.Fprintf(w, "\t%s", tag); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}
	return nil
}

func parseLine(line []string) (*SAM, error) {
	if len(line) < 11 {
		return nil, fmt.Errorf("too few fields: %v, want 11", len(line))
	}
	s := &SAM{}
	s.Qname = line[0]
	s.Rname = line[2]
	s.Cigar = line[5]
	s.Rnext = line[6]
	s.Seq = line[9]
	s.Qual = line[10]
	err := parseInts(snm.At(line, []int{1, 3, 4, 7, 8}),
		(*int)(&s.Flag), &s.Pos, &s.Mapq, &s.Pnext, &s.Tlen)
	if err != nil {
		return nil, err
	}
	s.Tags, err = parseTags(line[11:])
	if err != nil {
		return nil, err
	}
	return s, nil
}

func parseInts(strs []string, p ...*int) error {
	if len(strs) != len(p) {
		panic(fmt.Sprintf("mismatching lengths: %v, %v", len(strs), len(p)))
	}
	for i, s := range strs {
		n, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		*p[i] = n
	}
	return nil
}
