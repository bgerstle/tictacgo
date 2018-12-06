package tictacgo

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
)

type ArbitraryPendingBoard struct {
	Board
	Player1Token rune
	Player2Token rune
}

const lastSingleByteUTF8CodePoint = 0x7F

func (b ArbitraryPendingBoard) Generate(rand *rand.Rand, size int) reflect.Value {
	b.Board = EmptyBoard()
	b.Player1Token = rune(rand.Int31n(lastSingleByteUTF8CodePoint))
	b.Player2Token = rune(rand.Int31n(lastSingleByteUTF8CodePoint))
	numMoves := rand.Int31n(5)
	shuffledSpaces := rand.Perm(b.Board.SpacesLen())
	for i, space := range shuffledSpaces[:numMoves+1] {
		var currentToken rune
		if i%2 == 0 {
			currentToken = b.Player1Token
		} else {
			currentToken = b.Player2Token
		}
		b.Board = b.Board.AssignSpace(space, &currentToken)
	}
	return reflect.ValueOf(b)
}

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

	t.Run("boards with 4 or less moves", func(t *testing.T) {
		assert := assert.New(t)

		qcErr := quick.Check(func(arbitraryPendingBoard ArbitraryPendingBoard) bool {
			state, winner := arbitraryPendingBoard.Board.GameState()
			return state == Pending && winner == nil
		}, nil)
		assert.Nil(qcErr)
	})
}

type GameStateTestData struct {
	Board
	rune
}

func TestVictoryGameState(t *testing.T) {
	testBoardVictories := func(t *testing.T, name string, testData []GameStateTestData) {
		t.Helper()

		t.Run("test victory by "+name, func(t *testing.T) {
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
	assert := assert.New(t)

	tiedBoards := []Board{
		Board{
			spaces: []Space{
				X, O, X,
				X, O, O,
				O, X, X,
			},
		},
	}

	for _, board := range tiedBoards {
		state, winner := board.GameState()
		if state != Tie {
			t.Fatalf(
				"Expected board to be evaluated as tie, got %s. Board: \n%s",
				state,
				board.String(),
			)
		}
		assert.Nil(winner)
	}
}
