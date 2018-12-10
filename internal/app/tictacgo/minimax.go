package tictacgo

import (
	"fmt"
	"math"
	"sort"
)

// entry point to minimax algorithm.
func chooseSpotForActivePlayer(board Board) int {
	mm := minimax{
		maxPlayer: board.ActivePlayerToken(),
		minPlayer: board.NextPlayerToken(),
	}
	return mm.chooseSpotAtDepth(board, 0)
}

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

// Score a victory for the specified winner at the given depth.
// Earlier, shallower victories are given more weight than victories later in the game, i.e. deeper
func (mm minimax) scoreVictory(winner rune, depth int) float64 {
	return mm.scoreModifier(winner) * maxScore / float64(depth)
}

func (mm minimax) score(space int, board Board, depth int) score {
	playerToken := board.ActivePlayerToken()

	nextBoard, state, winner := board.AssignSpace(space)

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

	scores := mm.scoreSpaces(nextBoard, depth+1)
	return score{
		spaceIndex: space,
		value:      mm.minOrMax(playerToken, scores).value,
	}
}

func (mm minimax) scoreSpaces(board Board, depth int) (scores []score) {
	spaces := board.AvailableSpaces()
	scores = make([]score, len(spaces))
	for i, space := range spaces {
		scores[i] = mm.score(space, board, depth)
	}
	return
}

func (mm minimax) chooseSpotAtDepth(board Board, depth int) int {
	playerToken := board.ActivePlayerToken()
	scores := mm.scoreSpaces(board, depth)
	return mm.minOrMax(playerToken, scores).spaceIndex
}
