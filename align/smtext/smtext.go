// Package smtext handles text representations of substitution matrices.
//
// NCBI Format
//
// This is the format used in ftp://ftp.ncbi.nih.gov/blast/matrices/.
// Example:
//
//  # Optional comment line.
//  # Another optional comment line.
//     A  C  G  T
//  A  1 -2 -1 -2
//  C -2  1 -2 -1
//  G -1 -2  1 -2
//  T -2 -1 -2  1
//
// Specifically:
//
// Empty lines and lines beginning with '#' are ignored. Each line should have its
// values separated by whitespaces (any kind and any amount). The first line should
// have n letters. The others need to have n+1 values, where the first is a letter
// and the rest are float-parseable numbers.
package smtext

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strconv"

	"github.com/fluhus/biostuff/align"
)

// ReadNCBI decodes an NCBI-format substitution matrix from the given reader.
func ReadNCBI(r io.Reader) (align.SubstitutionMatrix, error) {
	sc := bufio.NewScanner(r)
	re := regexp.MustCompile(`\S+`)
	m := align.SubstitutionMatrix{}
	var chars []byte
	for sc.Scan() {
		row := sc.Text()
		if row == "" || row[0] == '#' {
			continue
		}
		if chars == nil { // First row
			charStrs := re.FindAllString(row, -1)
			for _, char := range charStrs {
				b, err := extractSingleChar(char)
				if err != nil {
					return nil, err
				}
				chars = append(chars, b)
			}
			continue
		}

		valStrs := re.FindAllString(row, -1)
		if len(valStrs) != len(chars)+1 {
			return nil, fmt.Errorf("bad number of values in %q: %v, want %v",
				row, len(valStrs), len(chars)+1)
		}
		c, err := extractSingleChar(valStrs[0])
		if err != nil {
			return nil, err
		}
		for i, val := range valStrs[1:] {
			x, err := strconv.ParseInt(val, 0, 0)
			if err != nil {
				return nil, fmt.Errorf("could not parse score: %v", err)
			}
			m[[2]byte{c, chars[i]}] = int(x)
		}
	}
	if sc.Err() != nil {
		return nil, sc.Err()
	}
	return m, nil
}

// Checks that a string is a single character and turns * into a gap.
func extractSingleChar(s string) (byte, error) {
	if len(s) != 1 {
		return 0, fmt.Errorf("expected a single character, got %q", s)
	}
	if s == "*" {
		return align.Gap, nil
	}
	return s[0], nil
}

// GoString returns a Go-literal of the given matrix.
func GoString(m align.SubstitutionMatrix) []byte {
	buf := bytes.NewBuffer(nil)
	fmt.Fprintln(buf, "SubstitutionMatrix{")
	var sorted [][]byte
	for k := range m {
		sorted = append(sorted, []byte{k[0], k[1]})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return bytes.Compare(sorted[i], sorted[j]) < 0
	})
	for _, k := range sorted {
		fmt.Fprintf(buf, "{%s,%s}:%v,\n",
			charOrGap(k[0]), charOrGap(k[1]), m.Get(k[0], k[1]))
	}
	fmt.Fprintln(buf, "}")
	return buf.Bytes()
}

// Returns a quoted char, or the constant Gap for a gap.
func charOrGap(c byte) string {
	if c == align.Gap {
		return "Gap"
	}
	return fmt.Sprintf("%q", c)
}
