package tictacgo

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIOHumanChoiceProviderIntegration(t *testing.T) {
	assertExpectedMovePrompt := func(t *testing.T, buf *bytes.Buffer, player Player) {
		assert := assert.New(t)
		t.Helper()

		assert.Equal(fmt.Sprintf(PlayerMovePromptf, player.Info().Token), buf.String())
	}

	t.Run("choose available space", func(t *testing.T) {
		assert := assert.New(t)

		board := EmptyBoard()

		for spaceNum := range board.spaces {
			mockOutput := &bytes.Buffer{}
			mockInput := bytes.NewBufferString(fmt.Sprintf("%d\n", spaceNum))

			choiceProvider := ioHumanChoiceProvider{
				in:  mockInput,
				out: mockOutput,
			}

			player := humanPlayer{
				PlayerInfo:     PlayerInfo{Token: 'X'},
				choiceProvider: choiceProvider,
			}

			choice := player.ChooseSpace(board)

			assertExpectedMovePrompt(t, mockOutput, player)

			assert.Equal(spaceNum, choice)
		}
	})
}
