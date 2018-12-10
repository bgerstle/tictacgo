package tictacgo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type tokenSpacesTuple struct {
	token  rune
	spaces []int
}

func popInt(xs *[]int) int {
	x := (*xs)[0]
	popped := (*xs)[1:]
	xs = &popped
	return x
}

func interleaveInts(s1, s2 []int) []int {
	ss := make([]int, len(s1)+len(s2))
	for i := range ss {
		var space int
		if i%2 == 0 {
			space = s1[0]
			s1 = s1[1:]
		} else {
			space = s2[0]
			s2 = s2[1:]
		}
		ss[i] = space
	}
	return ss
}

func (b Board) fillWithSpaces(spaces [2][]int) Board {
	interleaved := interleaveInts(spaces[0], spaces[1])

	if len(interleaved) > len(b.AvailableSpaces()) {
		panic(fmt.Sprintf(
			"Can't fill board with %d spaces when there are only %d available",
			len(interleaved),
			len(b.AvailableSpaces())))
	}

	for _, space := range interleaved {
		b, _, _ = b.AssignSpace(space)
	}

	return b
}

func NewTestBoardWithSpaces(spaces []Space) Board {
	b := NewEmptyTestBoard()

	activeSpaces := spacesAssignedTo(x, spaces)
	nextSpaces := spacesAssignedTo(o, spaces)
	return b.fillWithSpaces([2][]int{
		activeSpaces,
		nextSpaces,
	})
}

func TestPendingGameState(t *testing.T) {
	t.Run("example pending board", func(t *testing.T) {
		assert := assert.New(t)

		pendingBoards := []Board{
			NewEmptyTestBoard(),
			NewTestBoardWithSpaces([]Space{
				X, O, nil,
				nil, nil, nil,
				nil, nil, nil,
			}),
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
				NewTestBoardWithSpaces([]Space{
					X, X, X,
					nil, O, O,
					nil, nil, nil,
				}),
				x,
			},
			{
				NewTestBoardWithSpaces([]Space{
					X, nil, X,
					O, O, O,
					X, nil, nil,
				}),
				o,
			},
			{
				NewTestBoardWithSpaces([]Space{
					nil, nil, nil,
					O, nil, O,
					X, X, X,
				}),
				x,
			},
		},
	)

	testBoardVictories(
		t,
		"column",
		[]GameStateTestData{
			{
				NewTestBoardWithSpaces([]Space{
					X, O, O,
					X, nil, nil,
					X, nil, nil,
				}),
				x,
			},
			{
				NewTestBoardWithSpaces([]Space{
					nil, O, nil,
					X, O, X,
					nil, O, X,
				}),
				o,
			},
			{
				NewTestBoardWithSpaces([]Space{
					nil, nil, X,
					nil, nil, X,
					O, O, X,
				}),
				x,
			},
		},
	)

	testBoardVictories(
		t,
		"diagonal",
		[]GameStateTestData{
			{
				NewTestBoardWithSpaces([]Space{
					X, O, O,
					nil, X, nil,
					nil, nil, X,
				}),
				x,
			},
			{
				NewTestBoardWithSpaces([]Space{
					X, nil, O,
					X, O, X,
					O, nil, nil,
				}),
				o,
			},
		},
	)
}

func TestTieGameState(t *testing.T) {
	tiedBoards := []Board{
		NewTestBoardWithSpaces([]Space{
			X, O, X,
			X, O, O,
			O, X, X,
		}),
		NewTestBoardWithSpaces([]Space{
			X, O, X,
			X, O, X,
			O, X, O,
		}),
		NewTestBoardWithSpaces([]Space{
			O, X, X,
			X, O, O,
			X, O, X,
		}),
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
		'O', 'X', 'O',
		'O', 'O', 'X',
		'X', 'X', 'X',
	}
	spaces := make([]Space, len(runes))
	for i := range spaces {
		spaces[i] = Space(&runes[i])
	}
	board := NewTestBoardWithSpaces(spaces)
	state, _ := board.GameState()
	if state != Victory {
		t.Error("Expected board filled with all different pointers to have the expected result")
	}
}
