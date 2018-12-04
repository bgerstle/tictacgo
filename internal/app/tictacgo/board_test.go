package tictacgo

import (
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
}
