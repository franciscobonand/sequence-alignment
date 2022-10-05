package matrix

type Matrix [][]int

func New(seq1, seq2 string) Matrix {
	matrix := generateZeroMatrix(len(seq1), len(seq2))
	return matrix
}

func generateZeroMatrix(lines, cols int) Matrix {
	matrix := make([][]int, lines+1)

	for i := range matrix {
		matrix[i] = make([]int, cols+1)
	}

	return matrix
}
