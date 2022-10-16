package solver

import (
	"fmt"

	"github.com/franciscobonand/sequence-alignment/matrix"
	"github.com/franciscobonand/sequence-alignment/score"
)

// AlignmentSolver defines a solver's structure for the alignment problem
type AlignmentSolver struct {
	m          matrix.Matrix
	s          score.Evaluation
	seq1, seq2 string
}

// NewAlignmentSolver creates a new solver for the alignment problem
func NewAlignmentSolver(s score.Evaluation, seq1, seq2 string) AlignmentSolver {
	m := matrix.New(len(seq1), len(seq2))
	return AlignmentSolver{m: m, s: s, seq1: seq1, seq2: seq2}
}

// Solve solves the alignment problem
func (as AlignmentSolver) Solve() error {
	for x := 1; x < len(as.m); x++ {
		for y := 1; y < len(as.m[x]); y++ {
			cell, err := as.calculateCell(x, y)
			if err != nil {
				return fmt.Errorf("error on cell %d %d: %v", x, y, err)
			}
			as.m[x][y] = *cell
		}
	}
	return nil
}

// PrintResultMatrix prints the result matrix
func (as AlignmentSolver) PrintResultMatrix() {
	for i, line := range as.m {
		if i == 0 {
			fmt.Print("    ")
			for _, s := range as.seq1 {
				fmt.Printf("%c ", s)
			}
			fmt.Print("\n  ")
			for _, cell := range line {
				fmt.Printf("%d ", cell.Value)
			}
			fmt.Println()
			continue
		}
		fmt.Printf("%s ", string(as.seq2[i-1]))
		for _, cell := range line {
			fmt.Printf("%d ", cell.Value)
		}
		fmt.Println()
	}
}

// PrintResult prints the result of the alignment
func (as AlignmentSolver) PrintResult() {
	var seq1, seq2 string
	currCell := as.m[len(as.m)-1][len(as.m[0])-1]

	for {
		if currCell.From == matrix.Diagonal {
			seq1 = string(as.seq1[currCell.Y-1]) + seq1
			seq2 = string(as.seq2[currCell.X-1]) + seq2
			currCell = as.m[currCell.X-1][currCell.Y-1]
		} else if currCell.From == matrix.Left {
			seq1 = string(as.seq1[currCell.Y-1]) + seq1
			seq2 = "-" + seq2
			currCell = as.m[currCell.X][currCell.Y-1]
		} else if currCell.From == matrix.Top {
			seq1 = "-" + seq1
			seq2 = string(as.seq2[currCell.X-1]) + seq2
			currCell = as.m[currCell.X-1][currCell.Y]
		} else {
			break
		}
	}

	fmt.Println(seq1)
	fmt.Println(seq2)
}

// PrintResultMatrixDebug prints the result matrix with each matrix cell's full information
func (as AlignmentSolver) PrintResultMatrixDebug() {
	for i, line := range as.m {
		if i == 0 {
			fmt.Print("     ")
			for _, s := range as.seq1 {
				fmt.Printf("%c ", s)
			}
			fmt.Println("\n ", line)
			continue
		}
		fmt.Println(string(as.seq2[i-1]), line)
	}
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

	// TODO: Make it work for different size sequences
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
