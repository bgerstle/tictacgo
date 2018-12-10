package tictacgo

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

var x = 'X'
var o = 'O'
var X = Space(&x)
var O = Space(&o)

// Utility function for creating an empty test board w/ players X & O
func NewEmptyTestBoard() Board {
	return NewBoard([2]PlayerInfo{
		PlayerInfo{x},
		PlayerInfo{o},
	})
}

type ArbitrarySpaceIndexes []int

func (asi ArbitrarySpaceIndexes) Generate(rand *rand.Rand, size int) reflect.Value {
	length := 9 * size / math.MaxInt64
	spaceIndexes := rand.Perm(9)[:length]
	return reflect.ValueOf(spaceIndexes)
}

func TestBoardSpaces(t *testing.T) {
	t.Run("Initializes as empty, with all spaces available", func(t *testing.T) {
		assert := assert.New(t)

		board := NewEmptyTestBoard()

		availableSpaces := make([]int, len(board.spaces))
		for i, space := range board.spaces {
			availableSpaces[i] = i
			assert.True(board.IsSpaceAvailable(i))
			assert.Nil(space)
		}

		assert.Equal(availableSpaces, board.AvailableSpaces())
	})

	t.Run("Full board has no available spaces", func(t *testing.T) {
		assert := assert.New(t)

		fullBoard := Board{
			spaces: []Space{
				X, O, X,
				X, X, O,
				O, O, X,
			},
		}

		assert.Len(fullBoard.AvailableSpaces(), 0)
	})

	t.Run("Returns number of spaces for each token", func(t *testing.T) {
		assert := assert.New(t)

		qcErr := quick.Check(func(spaceIndexes ArbitrarySpaceIndexes) bool {
			board := NewEmptyTestBoard()
			xs := []int{}
			os := []int{}
			for _, space := range spaceIndexes {
				if board.ActivePlayerToken() == x {
					xs = append(xs, space)
				} else {
					os = append(os, space)
				}
				var state GameState
				board, state, _ = board.AssignSpace(space)
				if state != Pending {
					// don't keep going if state ends
					break
				}
			}

			gotExpectedXs := assert.ElementsMatch(xs, board.SpacesAssignedTo(x))
			gotExpectedOs := assert.ElementsMatch(os, board.SpacesAssignedTo(o))

			takenSpaces := append(xs, os...)
			availableSpaces := []int{}
			for i := 0; i < board.SpacesLen(); i++ {
				isTaken := false
				for _, takenSpace := range takenSpaces {
					if takenSpace == i {
						isTaken = true
						break
					}
				}
				if !isTaken {
					availableSpaces = append(availableSpaces, i)
				}
			}
			gotExpectedAvailableSpaces := assert.ElementsMatch(availableSpaces, board.AvailableSpaces())

			return gotExpectedXs && gotExpectedOs && gotExpectedAvailableSpaces
		}, nil)

		assert.NoError(qcErr)
	})
}

func TestBoard_AssigningSpacesAdvancesTurn(t *testing.T) {
	require := require.New(t)

	board := NewEmptyTestBoard()
	require.Equal(0, board.turn)
	require.Equal(x, board.ActivePlayerToken())

	var state GameState
	board, state, _ = board.AssignSpace(0)
	require.Equal(1, board.turn)
	require.Equal(o, board.ActivePlayerToken())
	if state != Pending {
		t.Fatalf("State should be pending, got %s", state)
	}

	board, state, _ = board.AssignSpace(1)
	require.Equal(2, board.turn)
	require.Equal(x, board.ActivePlayerToken())
	if state != Pending {
		t.Fatalf("State should be pending, got %s", state)
	}
}

func ExampleBoard_SpacesAssignedTo() {
	board, _, _ := NewEmptyTestBoard().AssignSpace(0)
	board, _, _ = board.AssignSpace(1)
	board, _, _ = board.AssignSpace(2)
	xsSpaces := board.SpacesAssignedTo(x)
	osSpaces := board.SpacesAssignedTo(o)
	fmt.Println(fmt.Sprintf("X has %d spaces: %d and %d", len(xsSpaces), xsSpaces[0], xsSpaces[1]))
	fmt.Println(fmt.Sprintf("O has %d space: %d", len(osSpaces), osSpaces[0]))
	// Output:
	// X has 2 spaces: 0 and 2
	// O has 1 space: 1
}

func TestBoardVectors(t *testing.T) {
	board := Board{
		spaces: []Space{
			nil, X, X,
			O, O, nil,
			X, nil, O,
		},
	}

	t.Run("Returns expected rows", func(t *testing.T) {
		assert := assert.New(t)

		rows := board.rows()
		assert.Equal([]Space{nil, X, X}, rows[0])
		assert.Equal([]Space{O, O, nil}, rows[1])
		assert.Equal([]Space{X, nil, O}, rows[2])
	})

	t.Run("Returns expected columns", func(t *testing.T) {
		assert := assert.New(t)

		columns := board.columns()
		assert.Equal([]Space{nil, O, X}, columns[0])
		assert.Equal([]Space{X, O, nil}, columns[1])
		assert.Equal([]Space{X, nil, O}, columns[2])
	})

	t.Run("Returns expected diagonals", func(t *testing.T) {
		assert := assert.New(t)

		diagonals := board.diagonals()
		assert.Equal([]Space{nil, O, O}, diagonals[0])
		assert.Equal([]Space{X, O, X}, diagonals[1])
	})
}

func TestBoardPrinting(t *testing.T) {
	t.Run("Prints expected output when empty", func(t *testing.T) {
		assert := assert.New(t)

		board := NewEmptyTestBoard()

		assert.Equal(
			strings.Join([]string{
				" 0 | 1 | 2 ",
				rowSeparator,
				" 3 | 4 | 5 ",
				rowSeparator,
				" 6 | 7 | 8 ",
				"",
			}, "\n"),
			board.String())
	})

	t.Run("Prints tokens in the spaces they occupy", func(t *testing.T) {
		assert := assert.New(t)

		board := Board{
			spaces: []Space{
				nil, X, X,
				O, O, nil,
				X, nil, O,
			},
		}
		boardStringLines := strings.Split(board.String(), "\n")
		assert.Equal(
			[]string{
				" 0 | X | X ",
				rowSeparator,
				" O | O | 5 ",
				rowSeparator,
				" X | 7 | O ",
				"",
			},
			boardStringLines)
	})

}

func TestSpaceToString(t *testing.T) {
	t.Run("Space with token", func(t *testing.T) {
		assert := assert.New(t)

		assert.Equal(string(x), spaceToString(X, "1"))
	})

	t.Run("Space without token", func(t *testing.T) {
		assert := assert.New(t)

		fallback := "0"

		assert.Equal(fallback, spaceToString(Space(nil), fallback))
	})
}
