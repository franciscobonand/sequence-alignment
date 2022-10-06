package solver

import "github.com/franciscobonand/sequence-alignment/matrix"

func max(diag, left, top int) *matrix.Cell {
	if diag >= left && diag >= top {
		return &matrix.Cell{Value: diag, From: matrix.Diagonal}
	}
	if left >= top {
		return &matrix.Cell{Value: left, From: matrix.Left}
	}
	return &matrix.Cell{Value: top, From: matrix.Top}
}
