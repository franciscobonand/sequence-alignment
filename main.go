package main

import (
	"log"

	"github.com/franciscobonand/sequence-alignment/score"
	"github.com/franciscobonand/sequence-alignment/solver"
)

func main() {
	// s := score.NewSimpleScore(1, 0, 0)
	s := score.NewNeedlemanWunschScore(0)
	as := solver.NewAlignmentSolver(s, "DRQTAQAAGTTTIT", "DRNTAQLLGTDTT")
	err := as.Solve()
	if err != nil {
		log.Fatalf("error solving alignment: %v", err)
	}
	as.PrintResultMatrix()
	as.PrintResult()
}
