// Package bed handles BED I/O.
//
// This package uses the format described in:
// https://en.wikipedia.org/wiki/BED_(file_format)
package bed

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO(amit): Overhaul the package.
// TODO(amit): Add writing.

const (
	PlusStrand  = "+"
	MinusStrand = "-"
	NoStrand    = "."
)

type BED struct {
	Chrom       string
	ChromStart  int
	ChromEnd    int
	Name        string
	Score       int
	Strand      string
	ThickStart  int
	ThickEnd    int
	ItemRGB     [3]byte
	BlockCount  int
	BlockSizes  []int
	BlockStarts []int
}

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
