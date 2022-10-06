package matrix

type Matrix [][]Cell

type Cell struct {
	Value int
	From  string
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
