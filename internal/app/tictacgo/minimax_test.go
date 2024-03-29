package tictacgo

import (
	"fmt"
	"math"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
)

func Test_minimaxUtils(t *testing.T) {
	mm := minimax{
		maxPlayer: x,
		minPlayer: o,
	}

	t.Run("opponent returns the other token", func(t *testing.T) {
		assert := assert.New(t)
		assert.Equal(x, mm.opponent(o))
		assert.Equal(o, mm.opponent(x))
	})

	t.Run("score modifier", func(t *testing.T) {
		assert := assert.New(t)
		assert.Equal(1.0, mm.scoreModifier(mm.maxPlayer))
		assert.Equal(-1.0, mm.scoreModifier(mm.minPlayer))
	})

	t.Run("minOrMax", func(t *testing.T) {
		assert := assert.New(t)
		scores := []score{
			score{
				spaceIndex: 0,
				value:      -1,
			},
			score{
				spaceIndex: 1,
				value:      1,
			},
		}
		assert.Equal(1, mm.minOrMax(mm.maxPlayer, scores).spaceIndex)
		assert.Equal(0, mm.minOrMax(mm.minPlayer, scores).spaceIndex)
	})
}

func Test_minimaxTieScoring(t *testing.T) {
	assert := assert.New(t)

	mm := minimax{
		maxPlayer: 'x',
		minPlayer: 'o',
	}

	board := NewTestBoardWithSpaces([]Space{
		X, O, X,
		X, X, O,
		O, nil, O,
	})

	score := mm.score(7, board, 0)

	assert.Equal(0.0, score.value)
}

func Test_minimaxScoring(t *testing.T) {
	assert := assert.New(t)

	mm := minimax{
		maxPlayer: 'x',
		minPlayer: 'o',
	}

	checkPropForPlayer := func(player rune, d1f, d2f float64) bool {
		t.Helper()

		shallow := int(math.Min(d1f, d2f))
		deep := int(math.Max(d1f, d2f))

		shallowScore := mm.scoreVictory(player, shallow)
		deepScore := mm.scoreVictory(player, deep)

		if player == mm.maxPlayer {
			return deepScore < shallowScore
		}
		// minimizing player gets scores that are increasingly negative
		return shallowScore < deepScore
	}

	// for arbitrary depth [0, n)...
	prefersShallowerVictories := func(d1, d2 uint16) bool {
		d1f, d2f := float64(d1), float64(d2)

		trueForMax := checkPropForPlayer(mm.maxPlayer, d1f, d2f)
		trueForMin := checkPropForPlayer(mm.minPlayer, d1f, d2f)

		return trueForMax && trueForMin
	}

	assert.NoError(quick.Check(prefersShallowerVictories, nil))
}

func Test_minimaxBlocksOpponent(t *testing.T) {
	assert := assert.New(t)

	board := NewTestBoardWithSpaces([]Space{
		X, X, nil,
		O, nil, nil,
		nil, nil, nil,
	})
	chosenMove := chooseSpaceForActivePlayer(board)

	assert.Equal(2, chosenMove)
}

type ArbitraryNearVictoryBoard struct {
	ArbitraryVictoryBoard
}

func Test_minimax_arbitraryPicksWinningMove(t *testing.T) {
	assert := assert.New(t)

	picksWinningMove := func(avb ArbitraryVictoryBoard) bool {
		xsSpaces := avb.SpacesAssignedTo(avb.WinningToken)
		osSpaces := avb.SpacesAssignedTo(avb.LosingToken)

		// create a new board by removing X's last winning move
		preVictoryBoard := NewBoard(avb.Board.players).fillWithSpaces([2][]int{
			xsSpaces[0 : len(xsSpaces)-1],
			osSpaces,
		})

		blockingMove := chooseSpaceForActivePlayer(preVictoryBoard)

		return blockingMove == xsSpaces[len(xsSpaces)-1]
	}

	assert.NoError(quick.Check(picksWinningMove, nil))
}

func Test_minimax_arbitraryBlockOpponent(t *testing.T) {
	assert := assert.New(t)

	blocksOpponent := func(avb ArbitraryVictoryBoard) bool {
		xsSpaces := avb.SpacesAssignedTo(avb.WinningToken)
		osSpaces := avb.SpacesAssignedTo(avb.LosingToken)

		// create a new board by removing one each of X's & O's moves
		preVictoryBoard := NewBoard(avb.Board.players).fillWithSpaces([2][]int{
			xsSpaces[1:],
			osSpaces[1:],
		})

		blockingMove := chooseSpaceForActivePlayer(preVictoryBoard)

		// should have picked the move that was taken from X
		return blockingMove == xsSpaces[0]
	}

	assert.NoError(quick.Check(blocksOpponent, nil))
}

func Example_minimax_ChooseSpot() {
	board := NewTestBoardWithSpaces([]Space{
		X, X, nil,
		O, O, nil,
		nil, nil, nil,
	})

	xsBest := chooseSpaceForActivePlayer(board)
	board, _, _ = board.AssignSpace(6)
	osBest := chooseSpaceForActivePlayer(board)

	fmt.Println(fmt.Sprintf("X's winning move in spot %d was chosen", xsBest))
	fmt.Println(fmt.Sprintf("O's winning move in spot %d was chosen", osBest))
	// Output:
	// X's winning move in spot 2 was chosen
	// O's winning move in spot 5 was chosen
}
