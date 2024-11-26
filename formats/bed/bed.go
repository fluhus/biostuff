// Package bed decodes and encodes BED files.
//
// This package uses the format described in:
// https://en.wikipedia.org/wiki/BED_(file_format)
//
// # Limitations
//
// Currently only tab delimiters are supported.
//
// Currently BED headers are not supported.
package bed

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Valid values for the strand field.
const (
	PlusStrand  = "+"
	MinusStrand = "-"
	NoStrand    = "."
)

// BED is a single line in a BED file.
type BED struct {
	N           int // Number of fields in this entry
	Chrom       string
	ChromStart  int // 0-based
	ChromEnd    int // 0-based exclusive
	Name        string
	Score       int
	Strand      string
	ThickStart  int
	ThickEnd    int
	ItemRGB     [3]byte
	BlockCount  int
	BlockSizes  []int // Length should match BlockCount
	BlockStarts []int // Length should match BlockCount
}

// Write writer the textual BED format representation of b to w.
// Encodes the first b.N fields, where b.N is between 3 and 12.
// Includes a trailing new line.
func (b *BED) Write(w io.Writer) error {
	if b.N < 3 || b.N > 12 {
		return fmt.Errorf("bad number of fields: %v, want 3-12", b.N)
	}
	if _, err := fmt.Fprintf(w, "%v\t%v\t%v",
		b.Chrom, b.ChromStart, b.ChromEnd); err != nil {
		return err
	}
	if b.N > 3 {
		if _, err := fmt.Fprintf(w, "\t%v", b.Name); err != nil {
			return err
		}
	}
	if b.N > 4 {
		if _, err := fmt.Fprintf(w, "\t%v", b.Score); err != nil {
			return err
		}
	}
	if b.N > 5 {
		if _, err := fmt.Fprintf(w, "\t%v", b.Strand); err != nil {
			return err
		}
	}
	if b.N > 6 {
		if _, err := fmt.Fprintf(w, "\t%v", b.ThickStart); err != nil {
			return err
		}
	}
	if b.N > 7 {
		if _, err := fmt.Fprintf(w, "\t%v", b.ThickEnd); err != nil {
			return err
		}
	}
	if b.N > 8 {
		if _, err := fmt.Fprintf(w, "\t%v,%v,%v",
			b.ItemRGB[0], b.ItemRGB[1], b.ItemRGB[2]); err != nil {
			return err
		}
	}
	if b.N > 9 {
		if _, err := fmt.Fprintf(w, "\t%v", b.BlockCount); err != nil {
			return err
		}
	}
	if b.N > 10 {
		if _, err := fmt.Fprintf(w, "\t"); err != nil {
			return err
		}
		for i, x := range b.BlockSizes {
			txt := "%v"
			if i > 0 {
				txt = ",%v"
			}
			if _, err := fmt.Fprintf(w, txt, x); err != nil {
				return err
			}
		}
	}
	if b.N > 11 {
		if _, err := fmt.Fprintf(w, "\t"); err != nil {
			return err
		}
		for i, x := range b.BlockStarts {
			txt := "%v"
			if i > 0 {
				txt = ",%v"
			}
			if _, err := fmt.Fprintf(w, txt, x); err != nil {
				return err
			}
		}
	}
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}
	return nil
}

// MarshalText returns the textual representation of b in BED format.
// Encodes the first b.N fields, where b.N is between 3 and 12.
// Includes a trailing new line.
func (b *BED) MarshalText() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	if err := b.Write(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Parses textual fields into a struct. Returns the number of parsed fields.
func parseLine(fields []string) (*BED, error) {
	n := len(fields)
	if n < 3 || n > 12 {
		return nil, fmt.Errorf("bad number of fields: %v, want 3-12", n)
	}

	// Force 12 fields to make parsing easy.
	fields = append(fields, make([]string, 12-n)...)
	bed := &BED{N: n}
	var err error

	// Mandatory fields.
	bed.Chrom = fields[0]
	if bed.ChromStart, err = strconv.Atoi(fields[1]); err != nil {
		return nil, fmt.Errorf("field 2: %v", err)
	}
	if bed.ChromEnd, err = strconv.Atoi(fields[2]); err != nil {
		return nil, fmt.Errorf("field 3: %v", err)
	}

	// Optional fields.
	bed.Name = fields[3]
	if fields[4] != "" {
		if bed.Score, err = strconv.Atoi(fields[4]); err != nil {
			return nil, fmt.Errorf("field 5: %v", err)
		}
	}
	if fields[5] != "" && fields[5] != PlusStrand &&
		fields[5] != MinusStrand && fields[5] != NoStrand {
		return nil, fmt.Errorf("field 6: bad strand: %q", fields[5])
	}
	bed.Strand = fields[5]
	if fields[6] != "" {
		if bed.ThickStart, err = strconv.Atoi(fields[6]); err != nil {
			return nil, fmt.Errorf("field 7: %v", err)
		}
	}
	if fields[7] != "" {
		if bed.ThickEnd, err = strconv.Atoi(fields[7]); err != nil {
			return nil, fmt.Errorf("field 8: %v", err)
		}
	}
	if fields[8] != "" {
		rgb := strings.Split(fields[8], ",")
		if len(rgb) != 3 {
			return nil, fmt.Errorf("field 9: bad RGB value: %q", fields[8])
		}
		for i := range rgb {
			a, err := strconv.ParseUint(rgb[i], 0, 8)
			if err != nil {
				return nil, fmt.Errorf("field 9: bad RGB value: %q", fields[8])
			}
			bed.ItemRGB[i] = byte(a)
		}
	}
	if fields[9] != "" {
		if bed.BlockCount, err = strconv.Atoi(fields[9]); err != nil {
			return nil, fmt.Errorf("field 10: %v", err)
		}
	}
	if fields[10] != "" {
		sizes := strings.Split(fields[10], ",")
		bed.BlockSizes = make([]int, len(sizes))
		for i := range sizes {
			bed.BlockSizes[i], err = strconv.Atoi(sizes[i])
			if err != nil {
				return nil, fmt.Errorf("field 11: %v", err)
			}
		}
	}
	if fields[11] != "" {
		starts := strings.Split(fields[11], ",")
		bed.BlockStarts = make([]int, len(starts))
		for i := range starts {
			bed.BlockStarts[i], err = strconv.Atoi(starts[i])
			if err != nil {
				return nil, fmt.Errorf("field 12: %v", err)
			}
		}
	}

	if len(bed.BlockSizes) != bed.BlockCount {
		return nil, fmt.Errorf("blockSizes has %v values but blockCount is %v",
			len(bed.BlockSizes), bed.BlockCount)
	}
	if len(bed.BlockStarts) != bed.BlockCount {
		return nil, fmt.Errorf("blockStarts has %v values but blockCount is %v",
			len(bed.BlockStarts), bed.BlockCount)
	}

	return bed, nil
}

// A reader reads and parses BED lines.
type reader struct {
	r *csv.Reader
}

// newReader returns a new BED reader that reads from r.
func newReader(r io.Reader) *reader {
	cr := csv.NewReader(r)
	cr.Comma = '\t'
	cr.Comment = '#'
	return &reader{cr}
}

// read returns the next BED line, and n as the number of fields that were found.
// The first n fields will be populated in the result BED, the rest will have zero
// values. n is always between 3 and 12.
//
// For example if n=5, then the populated fields are Chrom, ChromStart, ChromEnd,
// Name and Score.
func (r *reader) read() (b *BED, err error) {
	line, err := r.r.Read()
	if err != nil {
		return nil, err
	}
	return parseLine(line)
}
