package tictacgo

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer(t *testing.T) {
	assertExpectedMovePrompt := func(t *testing.T, buf *bytes.Buffer, player Player) {
		assert := assert.New(t)
		t.Helper()

		assert.Equal(fmt.Sprintf(PlayerMovePromptf, player.Token), buf.String())
	}

	t.Run("X chooses 0", func(t *testing.T) {
		assert := assert.New(t)

		mockOutput := &bytes.Buffer{}
		mockInput := bytes.NewBufferString("0\n")

		board := EmptyBoard()

		player := Player{Token: 'X'}

		choice := player.ChooseSpace(mockOutput, mockInput, board)

		assertExpectedMovePrompt(t, mockOutput, player)

		assert.Equal(0, choice)
	})

	t.Run("O chooses 4", func(t *testing.T) {
		assert := assert.New(t)

		mockOutput := &bytes.Buffer{}
		mockInput := bytes.NewBufferString("4\n")

		board := EmptyBoard()

		player := Player{Token: 'O'}

		choice := player.ChooseSpace(mockOutput, mockInput, board)

		assertExpectedMovePrompt(t, mockOutput, player)

		assert.Equal(4, choice)
	})
}
