package tictacgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPendingGameState(t *testing.T) {
	assert := assert.New(t)

	pendingBoards := []Board{
		EmptyBoard(),
		Board{
			spaces: []Space{
				X, O, nil,
				nil, nil, nil,
				nil, nil, nil,
			},
		},
	}

	for _, board := range pendingBoards {
		state, winner := board.GameState()

		if state != Pending {
			t.Fatalf(
				"Expected board to be evaluated as pending, got %s. Board: \n%s",
				state,
				board.String(),
			)
		}
		assert.Nil(winner)
	}
}

type GameStateTestData struct {
	Board
	rune
}

func TestVictoryGameState(t *testing.T) {
	testBoardVictories := func(t *testing.T, testData []GameStateTestData) {
		t.Helper()

		assert := assert.New(t)

		for _, tdata := range testData {
			state, winner := tdata.Board.GameState()

			if state != Victory {
				t.Fatalf(
					"Expected board to be evaluated as victory, got %s. Board: \n%s",
					state,
					tdata.Board.String(),
				)
			}
			assert.Equal(tdata.rune, *winner)
		}
	}

	t.Run("victory by row", func(t *testing.T) {
		testBoardVictories(
			t,
			[]GameStateTestData{
				{
					Board{
						spaces: []Space{
							X, X, X,
							nil, O, O,
							nil, nil, nil,
						},
					},
					x,
				},
				{
					Board{
						spaces: []Space{
							X, nil, nil,
							O, O, O,
							X, nil, nil,
						},
					},
					o,
				},
				{
					Board{
						spaces: []Space{
							nil, nil, nil,
							O, nil, O,
							X, X, X,
						},
					},
					x,
				},
			},
		)
	})
}
