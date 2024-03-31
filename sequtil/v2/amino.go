// Functions for amino acids.

package sequtil

import (
	"fmt"
)

const (
	// AminoAcids holds the single-letter amino acid symbols.
	// These are the valid inputs to AminoName.
	AminoAcids = "ABCDEFGHIKLMNPQRSTVWXYZ*"
)

// Translate translates the nucleotides in src to amino acids, appends the result to
// dst and returns the new slice. Nucleotides should be in "aAcCgGtT". Length of src
// should be a multiple of 3.
func Translate(dst, src []byte) []byte {
	if len(src)%3 != 0 {
		panic(fmt.Sprintf("length of src should be a multiple of 3, got %v",
			len(src)))
	}
	var buf [3]byte
	for i := 0; i < len(src); i += 3 {
		copy(buf[:], src[i:i+3])
		for j := range buf {
			if buf[j] >= 'a' {
				buf[j] -= 'a' - 'A'
			}
		}
		aa := codonToAmino[buf]
		if aa == 0 {
			panic(fmt.Sprintf("bad codon at position %v: %q", i, src[i:i+3]))
		}
		dst = append(dst, aa)
	}
	return dst
}

// TranslateReadingFrames returns the translation of the 3 reading frames of seq.
// Nucleotides should be in "aAcCgGtT". seq can be of any length.
func TranslateReadingFrames(seq []byte) [3][]byte {
	var result [3][]byte
	for i := 0; i < 3; i++ {
		sub := seq[i:]
		sub = sub[:len(sub)/3*3]
		result[i] = Translate(nil, sub)
	}
	return result
}

// AminoName returns the 3-letter code and the full name of the amino acid with the
// given letter. Input may be uppercase or lowercase.
func AminoName(aa byte) (string, string) {
	if aa >= 'a' && aa <= 'z' {
		aa -= 'a' - 'A'
	}
	names, ok := aminoToName[aa]
	if !ok {
		panic(fmt.Sprintf("bad amino acid code: '%c'", aa))
	}
	return names[0], names[1]
}

var codonToAmino = map[[3]byte]byte{
	{'A', 'A', 'A'}: 'K',
	{'A', 'A', 'C'}: 'N',
	{'A', 'A', 'G'}: 'K',
	{'A', 'A', 'T'}: 'N',
	{'A', 'C', 'A'}: 'T',
	{'A', 'C', 'C'}: 'T',
	{'A', 'C', 'G'}: 'T',
	{'A', 'C', 'T'}: 'T',
	{'A', 'G', 'A'}: 'R',
	{'A', 'G', 'C'}: 'S',
	{'A', 'G', 'G'}: 'R',
	{'A', 'G', 'T'}: 'S',
	{'A', 'T', 'A'}: 'I',
	{'A', 'T', 'C'}: 'I',
	{'A', 'T', 'G'}: 'M',
	{'A', 'T', 'T'}: 'I',
	{'C', 'A', 'A'}: 'Q',
	{'C', 'A', 'C'}: 'H',
	{'C', 'A', 'G'}: 'Q',
	{'C', 'A', 'T'}: 'H',
	{'C', 'C', 'A'}: 'P',
	{'C', 'C', 'C'}: 'P',
	{'C', 'C', 'G'}: 'P',
	{'C', 'C', 'T'}: 'P',
	{'C', 'G', 'A'}: 'R',
	{'C', 'G', 'C'}: 'R',
	{'C', 'G', 'G'}: 'R',
	{'C', 'G', 'T'}: 'R',
	{'C', 'T', 'A'}: 'L',
	{'C', 'T', 'C'}: 'L',
	{'C', 'T', 'G'}: 'L',
	{'C', 'T', 'T'}: 'L',
	{'G', 'A', 'A'}: 'E',
	{'G', 'A', 'C'}: 'D',
	{'G', 'A', 'G'}: 'E',
	{'G', 'A', 'T'}: 'D',
	{'G', 'C', 'A'}: 'A',
	{'G', 'C', 'C'}: 'A',
	{'G', 'C', 'G'}: 'A',
	{'G', 'C', 'T'}: 'A',
	{'G', 'G', 'A'}: 'G',
	{'G', 'G', 'C'}: 'G',
	{'G', 'G', 'G'}: 'G',
	{'G', 'G', 'T'}: 'G',
	{'G', 'T', 'A'}: 'V',
	{'G', 'T', 'C'}: 'V',
	{'G', 'T', 'G'}: 'V',
	{'G', 'T', 'T'}: 'V',
	{'T', 'A', 'A'}: '*',
	{'T', 'A', 'C'}: 'Y',
	{'T', 'A', 'G'}: '*',
	{'T', 'A', 'T'}: 'Y',
	{'T', 'C', 'A'}: 'S',
	{'T', 'C', 'C'}: 'S',
	{'T', 'C', 'G'}: 'S',
	{'T', 'C', 'T'}: 'S',
	{'T', 'G', 'A'}: '*',
	{'T', 'G', 'C'}: 'C',
	{'T', 'G', 'G'}: 'W',
	{'T', 'G', 'T'}: 'C',
	{'T', 'T', 'A'}: 'L',
	{'T', 'T', 'C'}: 'F',
	{'T', 'T', 'G'}: 'L',
	{'T', 'T', 'T'}: 'F',
}

var aminoToName = map[byte][2]string{
	'A': {"Ala", "Alanine"},
	'B': {"Asx", "Asparagine"},
	'C': {"Cys", "Cysteine"},
	'D': {"Asp", "Aspartic"},
	'E': {"Glu", "Glutamic"},
	'F': {"Phe", "Phenylalanine"},
	'G': {"Gly", "Glycine"},
	'H': {"His", "Histidine"},
	'I': {"Ile", "Isoleucine"},
	'K': {"Lys", "Lysine"},
	'L': {"Leu", "Leucine"},
	'M': {"Met", "Methionine"},
	'N': {"Asn", "Asparagine"},
	'P': {"Pro", "Proline"},
	'Q': {"Gln", "Glutamine"},
	'R': {"Arg", "Arginine"},
	'S': {"Ser", "Serine"},
	'T': {"Thr", "Threonine"},
	'V': {"Val", "Valine"},
	'W': {"Trp", "Tryptophan"},
	'X': {"X", "Any codon"},
	'Y': {"Tyr", "Tyrosine"},
	'Z': {"Glx", "Glutamine"},
	'*': {"*", "Stop codon"},
}
