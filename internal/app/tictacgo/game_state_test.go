package tictacgo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPendingGameState(t *testing.T) {
	t.Run("example pending board", func(t *testing.T) {
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
	})
}

type GameStateTestData struct {
	Board
	rune
}

func TestExampleVictoryGameState(t *testing.T) {
	testBoardVictories := func(t *testing.T, name string, testData []GameStateTestData) {
		t.Helper()

		t.Run(fmt.Sprintf(name), func(t *testing.T) {
			assert := assert.New(t)
			for _, tdata := range testData {
				state, winner := tdata.Board.GameState()

				if state != Victory {
					t.Errorf(
						"Expected board to be evaluated as victory, got %s. Board: \n%s",
						state,
						tdata.Board.String(),
					)
				} else {
					assert.Equal(tdata.rune, *winner)
				}
			}
		})
	}

	testBoardVictories(
		t,
		"row",
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

	testBoardVictories(
		t,
		"column",
		[]GameStateTestData{
			{
				Board{
					spaces: []Space{
						X, O, O,
						X, nil, nil,
						X, nil, nil,
					},
				},
				x,
			},
			{
				Board{
					spaces: []Space{
						nil, O, nil,
						X, O, X,
						nil, O, nil,
					},
				},
				o,
			},
			{
				Board{
					spaces: []Space{
						nil, nil, X,
						nil, nil, X,
						O, O, X,
					},
				},
				x,
			},
		},
	)

	testBoardVictories(
		t,
		"diagonal",
		[]GameStateTestData{
			{
				Board{
					spaces: []Space{
						X, O, O,
						nil, X, nil,
						nil, nil, X,
					},
				},
				x,
			},
			{
				Board{
					spaces: []Space{
						nil, nil, O,
						X, O, X,
						O, nil, nil,
					},
				},
				o,
			},
		},
	)
}

func TestTieGameState(t *testing.T) {
	tiedBoards := []Board{
		Board{
			spaces: []Space{
				X, O, X,
				X, O, O,
				O, X, X,
			},
		},
		Board{
			spaces: []Space{
				X, O, X,
				X, O, X,
				O, X, O,
			},
		},
		Board{
			spaces: []Space{
				O, X, X,
				X, O, O,
				X, O, X,
			},
		},
	}

	for i, board := range tiedBoards {
		t.Run(fmt.Sprintf("example %d", i), func(t *testing.T) {
			assert := assert.New(t)
			state, winner := board.GameState()
			if state != Tie {
				t.Fatalf(
					"Expected board to be evaluated as tie, got %s. Board: \n%s",
					state,
					board.String(),
				)
			}
			assert.Nil(winner)
		})
	}
}

func TestVictoryWithDifferentRefs(t *testing.T) {
	runes := []rune{
		'o', 'x', 'x',
		'x', 'o', 'x',
		'o', 'x', 'o',
	}
	spaces := make([]Space, len(runes))
	for i := range spaces {
		spaces[i] = Space(&runes[i])
	}
	board := Board{spaces}
	state, _ := board.GameState()
	if state != Victory {
		t.Error("Expected board filled with all different pointers to have the expected result")
	}
}
