package score

type Evaluation interface {
	Match() int
	Mismatch() int
	Gap() int
}
