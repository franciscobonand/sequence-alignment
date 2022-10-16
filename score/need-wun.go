package score

import blosum "github.com/franciscobonand/sequence-alignment/matrix/blossum"

// NeedlemanWunschScore defines a score for the Needleman-Wunsch algorithm
type NeedlemanWunschScore struct {
	gap int
}

// NewNeedlemanWunschScore creates a new NeedlemanWunschScore instance
func NewNeedlemanWunschScore(gap int) *NeedlemanWunschScore {
	return &NeedlemanWunschScore{gap}
}

// Match returns the match value
func (s *NeedlemanWunschScore) Match(seq1, seq2 string) int {
	if val, ok := blosum.Blosum62[seq1][seq2]; ok {
		return val
	}
	panic("invalid sequence character value")
}

// Mismatch returns the mismatch value
func (s *NeedlemanWunschScore) Mismatch(seq1, seq2 string) int {
	if val, ok := blosum.Blosum62[seq1][seq2]; ok {
		return val
	}
	panic("invalid sequence character value")
}

// Gap returns the gap value
func (s *NeedlemanWunschScore) Gap() int {
	return s.gap
}
