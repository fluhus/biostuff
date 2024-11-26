package sam

// Bit values of the flag field.
const (
	FlagMultiple           Flag = 1 << iota // Template having multiple segments in sequencing
	FlagEach                                // Each segment properly aligned according to the aligner
	FlagUnmapped                            // Segment unmapped
	FlagUnmapped2                           // Next segment in the template unmapped
	FlagReverseComplement                   // SEQ being reverse complemented
	FlagReverseComplement2                  // SEQ of the next segment in the template being reverse complemented
	FlagFirst                               // The first segment in the template
	FlagLast                                // The last segment in the template
	FlagSecondary                           // Secondary alignment
	FlagNotPassing                          // Not passing filters, such as platform/vendor quality controls
	FlagDuplicate                           // PCR or optical duplicate
	FlagSupplementary                       // Supplementary alignment
)

// Flag is the bitwise flag field in a SAM entry.
type Flag int

// Multiple returns the FlagMultiple bit value of this flag.
func (f Flag) Multiple() bool {
	return f&FlagMultiple > 0
}

// Each returns the FlagEach bit value of this flag.
func (f Flag) Each() bool {
	return f&FlagEach > 0
}

// Unmapped returns the FlagUnmapped bit value of this flag.
func (f Flag) Unmapped() bool {
	return f&FlagUnmapped > 0
}

// Unmapped2 returns the FlagUnmapped2 bit value of this flag.
func (f Flag) Unmapped2() bool {
	return f&FlagUnmapped2 > 0
}

// ReverseComplement returns the FlagReverseComplement
// bit value of this flag.
func (f Flag) ReverseComplement() bool {
	return f&FlagReverseComplement > 0
}

// ReverseComplement2 returns the FlagReverseComplement2
// bit value of this flag.
func (f Flag) ReverseComplement2() bool {
	return f&FlagReverseComplement2 > 0
}

// First returns the FlagFirst bit value of this flag.
func (f Flag) First() bool {
	return f&FlagFirst > 0
}

// Last returns the FlagLast bit value of this flag.
func (f Flag) Last() bool {
	return f&FlagLast > 0
}

// Secondary returns the FlagSecondary bit value of this flag.
func (f Flag) Secondary() bool {
	return f&FlagSecondary > 0
}

// NotPassing returns the FlagNotPassing bit value of this flag.
func (f Flag) NotPassing() bool {
	return f&FlagNotPassing > 0
}

// Duplicate returns the FlagDuplicate bit value of this flag.
func (f Flag) Duplicate() bool {
	return f&FlagDuplicate > 0
}

// Supplementary returns the FlagSupplementary
// bit value of this flag.
func (f Flag) Supplementary() bool {
	return f&FlagSupplementary > 0
}

// SetMultiple sets the FlagMultiple bit value for this flag.
func (f *Flag) SetMultiple(value bool) {
	if value {
		*f |= FlagMultiple
	} else {
		*f &= ^FlagMultiple
	}
}

func (f *Flag) SetEach(value bool) {
	if value {
		*f |= FlagEach
	} else {
		*f &= ^FlagEach
	}
}

func (f *Flag) SetUnmapped(value bool) {
	if value {
		*f |= FlagUnmapped
	} else {
		*f &= ^FlagUnmapped
	}
}

func (f *Flag) SetUnmapped2(value bool) {
	if value {
		*f |= FlagUnmapped2
	} else {
		*f &= ^FlagUnmapped2
	}
}

func (f *Flag) SetReverseComplement(value bool) {
	if value {
		*f |= FlagReverseComplement
	} else {
		*f &= ^FlagReverseComplement
	}
}

func (f *Flag) SetReverseComplement2(value bool) {
	if value {
		*f |= FlagReverseComplement2
	} else {
		*f &= ^FlagReverseComplement2
	}
}

func (f *Flag) SetFirst(value bool) {
	if value {
		*f |= FlagFirst
	} else {
		*f &= ^FlagFirst
	}
}

func (f *Flag) SetLast(value bool) {
	if value {
		*f |= FlagLast
	} else {
		*f &= ^FlagLast
	}
}

func (f *Flag) SetSecondary(value bool) {
	if value {
		*f |= FlagSecondary
	} else {
		*f &= ^FlagSecondary
	}
}

func (f *Flag) SetNotPassing(value bool) {
	if value {
		*f |= FlagNotPassing
	} else {
		*f &= ^FlagNotPassing
	}
}

func (f *Flag) SetDuplicate(value bool) {
	if value {
		*f |= FlagDuplicate
	} else {
		*f &= ^FlagDuplicate
	}
}

func (f *Flag) SetSupplementary(value bool) {
	if value {
		*f |= FlagSupplementary
	} else {
		*f &= ^FlagSupplementary
	}
}
