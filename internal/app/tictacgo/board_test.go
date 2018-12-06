package tictacgo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var x = 'X'
var o = 'O'
var X = Space(&x)
var O = Space(&o)

func TestBoard(t *testing.T) {
	t.Run("Initializes as empty, with all spaces available", func(t *testing.T) {
		assert := assert.New(t)

		board := EmptyBoard()

		availableSpaces := make([]int, len(board.spaces))
		for i, space := range board.spaces {
			availableSpaces[i] = i
			assert.True(board.IsSpaceAvailable(i))
			assert.Nil(space)
		}

		assert.Equal(availableSpaces, board.AvailableSpaces())
	})
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
}

func TestAssignSpace(t *testing.T) {
	t.Run("Returns all but the assigned spaces", func(t *testing.T) {
		testTakingAvailableSpace(t, EmptyBoard())
	})
}

func testTakingAvailableSpace(t *testing.T, board Board) {
	t.Helper()
	assert := assert.New(t)

	preAssignAvailableSpaces := board.AvailableSpaces()
	spaceToAssign := preAssignAvailableSpaces[0]
	expectedSpacesAfterAssign := preAssignAvailableSpaces[1:]

	newBoard := board.AssignSpace(spaceToAssign, X)

	assert.Equal(X, newBoard.spaces[spaceToAssign])
	assert.Equal(expectedSpacesAfterAssign, newBoard.AvailableSpaces())
	assert.Equal(preAssignAvailableSpaces, board.AvailableSpaces(), "original board should remain the same")
	assert.False(newBoard.IsSpaceAvailable(spaceToAssign))

	if len(newBoard.AvailableSpaces()) != 0 {
		testTakingAvailableSpace(t, newBoard)
	}
}

func TestBoardPrinting(t *testing.T) {
	t.Run("Prints expected output when empty", func(t *testing.T) {
		assert := assert.New(t)

		board := EmptyBoard()

		assert.Equal(
			` 0 | 1 | 2 
===+===+===
 3 | 4 | 5 
===+===+===
 6 | 7 | 8 
`,
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
