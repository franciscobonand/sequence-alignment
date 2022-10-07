package score

type NeedlemanWunschScore struct {
	gap int
}

func NewNeedlemanWunschScore(gap int) *NeedlemanWunschScore {
	return &NeedlemanWunschScore{gap}
}

func (s *NeedlemanWunschScore) Match(seq1, seq2 string) int {
	return 1
}

func (s *NeedlemanWunschScore) Mismatch(seq1, seq2 string) int {
	return 0
}

func (s *NeedlemanWunschScore) Gap() int {
	return s.gap
}
