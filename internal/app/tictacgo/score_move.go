package tictacgo

import (
	"math"
	"sort"
)

type ScoringFunc func(sp Space, currentToken rune, winnerToken rune, depth int) float64

const MaxScore = math.MaxFloat64

type Score struct {
	SpaceIndex int
	Value      float64
}

func (s Score) LessThan(otherScore Score) bool {
	return s.Value < otherScore.Value
}

type ScoreSlice []Score

func (ss *ScoreSlice) Sort() {
	sort.Slice(*ss, func(i, j int) bool {
		return (*ss)[i].LessThan((*ss)[j])
	})
}

func (b Board) ScoresFor(playerToken rune) (scores ScoreSlice) {
	for i := range b.AvailableSpaces() {
		scores = append(scores, Score{
			SpaceIndex: i,
			Value:      0.0,
		})
	}
	return
}
