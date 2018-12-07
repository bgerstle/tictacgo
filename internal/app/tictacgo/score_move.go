package tictacgo

import (
	"fmt"
	"math"
	"sort"
)

const maxScore = math.MaxFloat64

type score struct {
	spaceIndex int
	value      float64
}

func (s score) lessThan(otherScore score) bool {
	return s.value < otherScore.value
}

type minimax struct {
	maxPlayer rune
	minPlayer rune
}

// entry point to minimax algorithm.
func (mm minimax) chooseSpot(playerToken rune, board Board) int {
	return mm.chooseSpotAtDepth(playerToken, board, 0)
}

func (mm minimax) opponent(player rune) rune {
	if player == mm.maxPlayer {
		return mm.minPlayer
	}
	return mm.maxPlayer
}

func (mm minimax) scoreModifier(player rune) float64 {
	if player == mm.maxPlayer {
		return 1
	}
	return -1
}

func (mm minimax) minOrMax(player rune, ss []score) score {
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].lessThan(ss[j])
	})
	if player == mm.maxPlayer {
		return ss[len(ss)-1]
	}
	return ss[0]
}

func (mm minimax) scoreVictory(winner rune, depth int) float64 {
	return mm.scoreModifier(winner) * maxScore / float64(depth)
}

func (mm minimax) score(space int, board Board, playerToken rune, depth int) score {
	result := board.AssignSpace(space, &playerToken)
	state, winner := result.GameState()
	switch state {
	case Victory:
		return score{
			spaceIndex: space,
			value:      mm.scoreVictory(*winner, depth),
		}
	case Tie:
		return score{
			spaceIndex: space,
			value:      0,
		}
	}

	if state != Pending {
		panic(fmt.Sprintf("Unexpected state: %s", state))
	}

	spaces := board.AvailableSpaces()
	if len(spaces) == 0 {
		panic("Should have been a tie")
	}

	nextPlayer := mm.opponent(playerToken)
	scores := mm.scoreSpaces(nextPlayer, result, depth+1)
	return score{
		spaceIndex: space,
		value:      mm.minOrMax(playerToken, scores).value,
	}
}

func (mm minimax) scoreSpaces(player rune, board Board, depth int) (scores []score) {
	spaces := board.AvailableSpaces()
	scores = make([]score, len(spaces))
	for i, space := range spaces {
		scores[i] = mm.score(space, board, player, depth)
	}
	return
}

func (mm minimax) chooseSpotAtDepth(playerToken rune, board Board, depth int) int {
	scores := mm.scoreSpaces(playerToken, board, depth)
	return mm.minOrMax(playerToken, scores).spaceIndex
}
