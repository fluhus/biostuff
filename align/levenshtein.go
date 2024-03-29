// Concrete substitution matrices.

package align

// Initializes Levenshtein.
func init() {
	Levenshtein = make(SubstitutionMatrix, 256*256)
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			if i != j {
				Levenshtein[[2]byte{byte(i), byte(j)}] = -1
			} else {
				Levenshtein[[2]byte{byte(i), byte(j)}] = 0
			}
		}
	}
}
