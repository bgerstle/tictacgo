package tictacgo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoard(t *testing.T) {
	assert := assert.New(t)

	t.Run("Initializes as empty", func(t *testing.T) {
		board := EmptyBoard()

		for _, space := range board.spaces {
			assert.Nil(space)
		}
	})

	t.Run("Prints expected output when empty", func(t *testing.T) {
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

	t.Run("Returns expected rows", func(t *testing.T) {
		x := 'X'
		o := 'O'
		X := Space(&x)
		O := Space(&o)
		board := Board{
			spaces: []Space{
				nil, X, X,
				O, O, nil,
				X, nil, O,
			},
		}
		rows := board.rows()
		assert.Equal([]Space{nil, X, X}, rows[0])
		assert.Equal([]Space{O, O, nil}, rows[1])
		assert.Equal([]Space{X, nil, O}, rows[2])
	})

	t.Run("Prints tokens in the spaces they occupy", func(t *testing.T) {
		x := 'X'
		o := 'O'
		X := Space(&x)
		O := Space(&o)
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
	assert := assert.New(t)

	t.Run("Space with token", func(t *testing.T) {
		x := 'X'
		X := Space(&x)

		assert.Equal(string(x), spaceToString(X, "1"))
	})

	t.Run("Space without token", func(t *testing.T) {
		fallback := "0"
		X := Space(nil)

		assert.Equal(fallback, spaceToString(X, fallback))
	})
}
