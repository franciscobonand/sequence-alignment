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

func New(lines, cols int) Matrix {
	matrix := generateZeroMatrix(lines, cols)
	return matrix
}

func generateZeroMatrix(lines, cols int) Matrix {
	matrix := make([][]Cell, cols+1)

	for i := range matrix {
		matrix[i] = make([]Cell, lines+1)
	}

	return matrix
}
