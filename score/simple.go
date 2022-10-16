package score

// SimpleScore is a simple scoring system that only takes into account the
// match, mismatch and gap values informed by the user
type SimpleScore struct {
	match, mismatch, gap int
}

// NewSimpleScore creates a new SimpleScore instance
func NewSimpleScore(match, mismatch, gap int) *SimpleScore {
	return &SimpleScore{match, mismatch, gap}
}

// Match returns the match value
func (s *SimpleScore) Match(seq1, seq2 string) int {
	return s.match
}

// Mismatch returns the mismatch value
func (s *SimpleScore) Mismatch(seq1, seq2 string) int {
	return s.mismatch
}

// Gap returns the gap value
func (s *SimpleScore) Gap() int {
	return s.gap
}
