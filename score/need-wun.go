package score

import blosum "github.com/franciscobonand/sequence-alignment/matrix/blossum"

// Blosum62Score defines a score for the Needleman-Wunsch algorithm
type Blosum62Score struct {
	gap int
}

// NewBlosum62Score creates a new Blosum62Score instance
func NewBlosum62Score(gap int) *Blosum62Score {
	return &Blosum62Score{gap}
}

// Match returns the match value
func (s *Blosum62Score) Match(seq1, seq2 string) int {
	if val, ok := blosum.Blosum62[seq1][seq2]; ok {
		return val
	}
	panic("invalid sequence character value")
}

// Mismatch returns the mismatch value
func (s *Blosum62Score) Mismatch(seq1, seq2 string) int {
	if val, ok := blosum.Blosum62[seq1][seq2]; ok {
		return val
	}
	panic("invalid sequence character value")
}

// Gap returns the gap value
func (s *Blosum62Score) Gap() int {
	return s.gap
}
