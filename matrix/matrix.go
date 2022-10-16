package matrix

// Matrix defines a matrix
type Matrix [][]Cell

// Direction defines the direction of the path to be taken in the matrix
type Direction string

const (
	Left     Direction = "Left"
	Top      Direction = "Top"
	Diagonal Direction = "Diag"
)

// Cell defines a cell in the matrix
type Cell struct {
	Value int
	From  Direction
	X, Y  int
}

// New creates a new matrix
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
