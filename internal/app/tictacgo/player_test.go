package tictacgo

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer(t *testing.T) {
	assert := assert.New(t)

	t.Run("Choose space 0", func(t *testing.T) {
		mockOutput := &bytes.Buffer{}
		mockInput := bytes.NewBufferString("0\n")

		board := EmptyBoard()

		player := Player{Token: 'X'}

		choice := player.ChooseSpace(mockOutput, mockInput, board)

		assert.Equal("Make your move, X: ", mockOutput.String())
		assert.Equal(0, choice)
	})

	t.Run("Choose space 4", func(t *testing.T) {
		mockOutput := &bytes.Buffer{}
		mockInput := bytes.NewBufferString("4\n")

		board := EmptyBoard()

		player := Player{Token: 'X'}

		choice := player.ChooseSpace(mockOutput, mockInput, board)

		assert.Equal("Make your move, X: ", mockOutput.String())
		assert.Equal(4, choice)
	})
}
