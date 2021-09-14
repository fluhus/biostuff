// Package bed parses BED files.
//
// This package uses the format described in:
// https://en.wikipedia.org/wiki/BED_(file_format)
//
// Limitations
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

// Returns the first n fields as strings.
func (b *BED) toStrings(n int) (fields []string) {
	fields = make([]string, n)
	fields[0] = b.Chrom
	fields[1] = fmt.Sprint(b.ChromStart)
	fields[2] = fmt.Sprint(b.ChromEnd)
	if n == 3 {
		return
	}
	fields[3] = b.Name
	if n == 4 {
		return
	}
	fields[4] = fmt.Sprint(b.Score)
	if n == 5 {
		return
	}
	fields[5] = b.Strand
	if n == 6 {
		return
	}
	fields[6] = fmt.Sprint(b.ThickStart)
	if n == 7 {
		return
	}
	fields[7] = fmt.Sprint(b.ThickEnd)
	if n == 8 {
		return
	}
	fields[8] = fmt.Sprintf("%v,%v,%v", b.ItemRGB[0], b.ItemRGB[1], b.ItemRGB[2])
	if n == 9 {
		return
	}
	fields[9] = fmt.Sprint(b.BlockCount)
	if n == 10 {
		return
	}
	strs := make([]string, len(b.BlockSizes))
	for i := range strs {
		strs[i] = fmt.Sprint(b.BlockSizes[i])
	}
	fields[10] = strings.Join(strs, ",")
	if n == 11 {
		return
	}
	strs = make([]string, len(b.BlockStarts))
	for i := range strs {
		strs[i] = fmt.Sprint(b.BlockStarts[i])
	}
	fields[11] = strings.Join(strs, ",")
	return
}

// Text returns the textual representation of b in BED format. Encodes the first n
// fields, where n is between 3 and 12. Includes a trailing new line.
func (b *BED) Text(n int) []byte {
	if n < 3 || n > 12 {
		panic(fmt.Sprintf("bad n: %v, should be 3-12", n))
	}
	strs := b.toStrings(n)
	buf := bytes.NewBuffer(nil)
	w := csv.NewWriter(buf)
	w.Comma = '\t'
	w.Write(strs)
	w.Flush()
	return buf.Bytes()
}

// Parses textual fields into a struct. Returns the number of parsed fields.
func parseLine(fields []string) (*BED, int, error) {
	n := len(fields)
	if n < 3 || n > 12 {
		return nil, 0, fmt.Errorf("bad number of fields: %v, want 3-12", n)
	}

	// Force 12 fields to make parsing easy.
	fields = append(fields, make([]string, 12-n)...)
	bed := &BED{}
	var err error

	// Mandatory fields.
	bed.Chrom = fields[0]
	if bed.ChromStart, err = strconv.Atoi(fields[1]); err != nil {
		return nil, 0, fmt.Errorf("field 2: %v", err)
	}
	if bed.ChromEnd, err = strconv.Atoi(fields[2]); err != nil {
		return nil, 0, fmt.Errorf("field 3: %v", err)
	}

	// Optional fields.
	bed.Name = fields[3]
	if fields[4] != "" {
		if bed.Score, err = strconv.Atoi(fields[4]); err != nil {
			return nil, 0, fmt.Errorf("field 5: %v", err)
		}
	}
	if fields[5] != "" && fields[5] != PlusStrand &&
		fields[5] != MinusStrand && fields[5] != NoStrand {
		return nil, 0, fmt.Errorf("field 6: bad strand: %q", fields[5])
	}
	bed.Strand = fields[5]
	if fields[6] != "" {
		if bed.ThickStart, err = strconv.Atoi(fields[6]); err != nil {
			return nil, 0, fmt.Errorf("field 7: %v", err)
		}
	}
	if fields[7] != "" {
		if bed.ThickEnd, err = strconv.Atoi(fields[7]); err != nil {
			return nil, 0, fmt.Errorf("field 8: %v", err)
		}
	}
	if fields[8] != "" {
		rgb := strings.Split(fields[8], ",")
		if len(rgb) != 3 {
			return nil, 0, fmt.Errorf("field 9: bad RGB value: %q", fields[8])
		}
		for i := range rgb {
			a, err := strconv.ParseUint(rgb[i], 0, 8)
			if err != nil {
				return nil, 0, fmt.Errorf("field 9: bad RGB value: %q", fields[8])
			}
			bed.ItemRGB[i] = byte(a)
		}
	}
	if fields[9] != "" {
		if bed.BlockCount, err = strconv.Atoi(fields[9]); err != nil {
			return nil, 0, fmt.Errorf("field 10: %v", err)
		}
	}
	if fields[10] != "" {
		sizes := strings.Split(fields[10], ",")
		bed.BlockSizes = make([]int, len(sizes))
		for i := range sizes {
			bed.BlockSizes[i], err = strconv.Atoi(sizes[i])
			if err != nil {
				return nil, 0, fmt.Errorf("field 11: %v", err)
			}
		}
	}
	if fields[11] != "" {
		starts := strings.Split(fields[11], ",")
		bed.BlockStarts = make([]int, len(starts))
		for i := range starts {
			bed.BlockStarts[i], err = strconv.Atoi(starts[i])
			if err != nil {
				return nil, 0, fmt.Errorf("field 12: %v", err)
			}
		}
	}

	if len(bed.BlockSizes) != bed.BlockCount {
		return nil, 0, fmt.Errorf("blockSizes has %v values but blockCount is %v",
			len(bed.BlockSizes), bed.BlockCount)
	}
	if len(bed.BlockStarts) != bed.BlockCount {
		return nil, 0, fmt.Errorf("blockStarts has %v values but blockCount is %v",
			len(bed.BlockStarts), bed.BlockCount)
	}

	return bed, n, nil
}

// A Reader reads and parses BED lines.
type Reader struct {
	r *csv.Reader
}

// NewReader returns a new BED reader that reads from r.
func NewReader(r io.Reader) *Reader {
	cr := csv.NewReader(r)
	cr.Comma = '\t'
	cr.Comment = '#'
	return &Reader{cr}
}

// Next returns the next BED line, and n as the number of fields that were found.
// The first n fields will be populated in the result BED, the rest will have zero
// values. n is always between 3 and 12.
//
// For example if n=5, then the populated fields are Chrom, ChromStart, ChromEnd,
// Name and Score.
func (r *Reader) Next() (b *BED, n int, err error) {
	line, err := r.r.Read()
	if err != nil {
		return nil, 0, err
	}
	return parseLine(line)
}
