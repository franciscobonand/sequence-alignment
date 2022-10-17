package solver

import (
	"fmt"

	"github.com/franciscobonand/sequence-alignment/score"
)

// MultipleAlignmentSolver defines a solver's structure for multiple sequences alignment problem
type MultipleAlignmentSolver struct {
	s    score.Evaluation
	seqs []string
}

// NewMultipleAlignmentSolver creates a new solver for multiple sequences alignment problem
func NewMultipleAlignmentSolver(s score.Evaluation, seqs ...string) MultipleAlignmentSolver {
	return MultipleAlignmentSolver{s: s, seqs: seqs}
}

func (mas MultipleAlignmentSolver) Solve() error {
	var aligned []string
	var nextAlignedSeq string

	orderedSeqs := getMostSimilarSequences(mas.seqs)
	as := NewAlignmentSolver(mas.s, orderedSeqs[0], orderedSeqs[1])
	err := as.Solve()
	if err != nil {
		return err
	}
	alignedSeq1, alignedSeq2 := as.ReturnResult()
	aligned = append(aligned, alignedSeq1, alignedSeq2)

	for i, seq := range orderedSeqs {
		if i == 0 || i == 1 {
			continue
		}
		as = NewAlignmentSolver(mas.s, orderedSeqs[i-1], seq)
		err = as.Solve()
		if err != nil {
			return err
		}
		_, nextAlignedSeq = as.ReturnResult()
		aligned = append(aligned, nextAlignedSeq)
	}

	printResult(aligned)
	return nil
}

func getMostSimilarSequences(seqs []string) []string {
	scores := make(map[string]int)

	for i := 0; i < len(seqs); i++ {
		for j := i + 1; j < len(seqs); j++ {
			if len(seqs[i]) > len(seqs[j]) {
				score := sequenceSimilarityScore(seqs[j], seqs[i])
				scores[seqs[i]] += score
				scores[seqs[j]] += score
			} else {
				score := sequenceSimilarityScore(seqs[i], seqs[j])
				scores[seqs[i]] += score
				scores[seqs[j]] += score
			}
		}
	}

	return getOrderedSequences(scores)
}

func sequenceSimilarityScore(s1, s2 string) int {
	score := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			score++
		}
	}
	return score
}

func getOrderedSequences(scores map[string]int) []string {
	var biggestScore int
	var biggestScoreSeq string
	var orderedSeqs []string
	for i := 0; i < len(scores); i++ {
		for seq, score := range scores {
			if score > biggestScore {
				biggestScore = score
				biggestScoreSeq = seq
			}
		}
		orderedSeqs = append(orderedSeqs, biggestScoreSeq)
		scores[biggestScoreSeq] = -1
		biggestScore = 0
	}
	return orderedSeqs
}

// printResult prints the result of the alignment
func printResult(seq []string) {
	for _, s := range seq {
		fmt.Println(s)
	}
}
