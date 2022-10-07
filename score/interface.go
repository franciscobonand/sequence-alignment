package score

type Evaluation interface {
	Match(seq1, seq2 string) int
	Mismatch(seq1, seq2 string) int
	Gap() int
}
