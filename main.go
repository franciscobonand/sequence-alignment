package main

import (
	"log"

	"github.com/franciscobonand/sequence-alignment/score"
	"github.com/franciscobonand/sequence-alignment/solver"
)

func main() {
	s := score.NewSimpleScore()
	as := solver.NewAlignmentSolver(s, "ATCGTAC", "ATGTTAT")
	err := as.Solve("ATCGTAC", "ATGTTAT")
	if err != nil {
		log.Fatalf("error solving alignment: %v", err)
	}
	as.PrintResultMatrix()
	as.PrintResult()
}
