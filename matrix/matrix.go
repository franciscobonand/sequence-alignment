package matrix

type Matrix [][]Cell

type Direction string

const (
	Left     Direction = "Left"
	Top      Direction = "Top"
	Diagonal Direction = "Diag"
)

type Cell struct {
	Value int
	From  Direction
	X, Y  int
}

func New(seq1, seq2 string) Matrix {
	matrix := generateZeroMatrix(len(seq1), len(seq2))
	return matrix
}

func generateZeroMatrix(lines, cols int) Matrix {
	matrix := make([][]Cell, lines+1)

	for i := range matrix {
		matrix[i] = make([]Cell, cols+1)
	}

	return matrix
}
