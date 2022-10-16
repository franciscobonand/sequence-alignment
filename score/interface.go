package score

// Evaluation defines the interface for the scoring system
type Evaluation interface {
	Match(seq1, seq2 string) int
	Mismatch(seq1, seq2 string) int
	Gap() int
}
