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
		assert.Equal([]Space{X, nil, O}, rows[1])
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
		assert.Equal(" 0 | X | X ", boardStringLines[0])
		assert.Equal(" O | O | 5 ", boardStringLines[1])
		assert.Equal(" X | 7 | O ", boardStringLines[2])
	})
}
