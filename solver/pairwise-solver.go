package solver

import (
	"fmt"

	"github.com/franciscobonand/sequence-alignment/matrix"
	"github.com/franciscobonand/sequence-alignment/score"
)

// AlignmentSolver defines a solver's structure for the pairwise alignment problem
type AlignmentSolver struct {
	m          matrix.Matrix
	s          score.Evaluation
	seq1, seq2 string
}

// NewAlignmentSolver creates a new solver for the pairwise alignment problem
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
			fmt.Printf("%d", cell.Value)
			if cell.From == matrix.Diagonal {
				fmt.Print("\\ ")
			} else if cell.From == matrix.Left {
				fmt.Print("_ ")
			} else if cell.From == matrix.Top {
				fmt.Print("| ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

// PrintResult prints the result of the pairwise alignment
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

// ReturnResult returns the result of the pairwise alignment
func (as AlignmentSolver) ReturnResult() (string, string) {
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

	return seq1, seq2
}
