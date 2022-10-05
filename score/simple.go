package score

type SimpleScore struct {
}

func NewSimpleScore() *SimpleScore {
	return &SimpleScore{}
}

func (s *SimpleScore) Match() int {
	return 1
}

func (s *SimpleScore) Mismatch() int {
	return 0
}

func (s *SimpleScore) Gap() int {
	return 0
}
