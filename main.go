package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/franciscobonand/sequence-alignment/score"
	"github.com/franciscobonand/sequence-alignment/solver"
)

func main() {
	args := os.Args[1:]
	if len(args) > 2 {
		var s score.Evaluation
		var seqs []string
		if args[0] == "simple" && len(args) > 5 {
			scoreValues, err := stringToIntegerSlice(args[1], args[2], args[3])
			if err != nil {
				raiseError()
			}
			s = score.NewSimpleScore(scoreValues[0], scoreValues[1], scoreValues[2])
			seqs = args[4:]
		} else if args[0] == "blosum" && len(args) > 3 {
			gap, err := stringToIntegerSlice(args[1])
			if err != nil {
				raiseError()
			}
			s = score.NewBlosum62Score(gap[0])
			seqs = args[2:]
		} else {
			raiseError()
		}

		if len(seqs) == 2 {
			as := solver.NewAlignmentSolver(s, seqs[0], seqs[1])
			err := as.Solve()
			if err != nil {
				log.Fatal(err)
			}
			as.PrintResultMatrix()
			as.PrintResult()
		} else {
			mas := solver.NewMultipleAlignmentSolver(s, seqs...)
			err := mas.Solve()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func stringToIntegerSlice(s ...string) ([]int, error) {
	var intSlice []int
	for _, v := range s {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		intSlice = append(intSlice, n)
	}
	return intSlice, nil
}

func raiseError() {
	fmt.Println("Must use one of the following options:")
	fmt.Println("Usage: go run . simple <match> <mismatch> <gap> <sequence1> <sequence2>")
	fmt.Println("Usage: go run . blosum <gap> <sequence1> <sequence2>")
	log.Fatal("Wrong arguments")
}
