package score

type SimpleScore struct {
	match, mismatch, gap int
}

func NewSimpleScore(match, mismatch, gap int) *SimpleScore {
	return &SimpleScore{match, mismatch, gap}
}

func (s *SimpleScore) Match(seq1, seq2 string) int {
	return s.match
}

func (s *SimpleScore) Mismatch(seq1, seq2 string) int {
	return s.mismatch
}

func (s *SimpleScore) Gap() int {
	return s.gap
}
