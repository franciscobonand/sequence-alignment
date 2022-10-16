package score

import blosum "github.com/franciscobonand/sequence-alignment/matrix/blossum"

type NeedlemanWunschScore struct {
	gap int
}

func NewNeedlemanWunschScore(gap int) *NeedlemanWunschScore {
	return &NeedlemanWunschScore{gap}
}

func (s *NeedlemanWunschScore) Match(seq1, seq2 string) int {
	if val, ok := blosum.Blosum62[seq1][seq2]; ok {
		return val
	}
	panic("invalid sequence character value")
}

func (s *NeedlemanWunschScore) Mismatch(seq1, seq2 string) int {
	if val, ok := blosum.Blosum62[seq1][seq2]; ok {
		return val
	}
	panic("invalid sequence character value")
}

func (s *NeedlemanWunschScore) Gap() int {
	return s.gap
}
