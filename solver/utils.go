package solver

import (
	"fmt"

	"github.com/franciscobonand/sequence-alignment/matrix"
)

func max(diag, left, top int) *matrix.Cell {
	if diag >= left && diag >= top {
		return &matrix.Cell{Value: diag, From: matrix.Diagonal}
	}
	if left >= top {
		return &matrix.Cell{Value: left, From: matrix.Left}
	}
	return &matrix.Cell{Value: top, From: matrix.Top}
}

func (as AlignmentSolver) calculateCell(x, y int) (*matrix.Cell, error) {
	diagonal, err := as.getDiagonalValue(x, y)
	if err != nil {
		return nil, err
	}
	left, err := as.getLeftValue(x, y)
	if err != nil {
		return nil, err
	}
	top, err := as.getTopValue(x, y)
	if err != nil {
		return nil, err
	}

	cell := max(diagonal, left, top)
	cell.X = x
	cell.Y = y
	return cell, nil
}

func (as AlignmentSolver) getDiagonalValue(x, y int) (int, error) {
	if x < 1 || y < 1 || x >= len(as.m) || y >= len(as.m[0]) {
		return 0, fmt.Errorf("index x=%d or y=%d out of range", x, y)
	}

	if as.seq1[y-1] == as.seq2[x-1] {
		return as.m[x-1][y-1].Value + as.s.Match(string(as.seq1[y-1]), string(as.seq2[x-1])), nil
	}

	return as.m[x-1][y-1].Value + as.s.Mismatch(string(as.seq1[y-1]), string(as.seq2[x-1])), nil
}

func (as AlignmentSolver) getLeftValue(x, y int) (int, error) {
	if y < 1 || y >= len(as.m[0]) {
		return 0, fmt.Errorf("index x=%d out of range", x)
	}

	return as.m[x][y-1].Value + as.s.Gap(), nil
}

func (as AlignmentSolver) getTopValue(x, y int) (int, error) {
	if x < 1 || x >= len(as.m) {
		return 0, fmt.Errorf("index y=%d out of range", y)
	}

	return as.m[x-1][y].Value + as.s.Gap(), nil
}
