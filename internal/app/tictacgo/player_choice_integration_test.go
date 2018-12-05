package tictacgo

import (
	"bufio"
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

	t.Run("returns choice specified in input", func(t *testing.T) {
		assert := assert.New(t)

		board := EmptyBoard()

		for spaceNum := range board.AvailableSpaces() {
			mockOutput := &bytes.Buffer{}
			mockInput := bufio.NewReader(bytes.NewBufferString(fmt.Sprintf("%d\n", spaceNum)))

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

	t.Run("returns second, valid choice after invalid choice", func(t *testing.T) {
		assert := assert.New(t)

		board := EmptyBoard()

		mockOutput := &bytes.Buffer{}

		availableSpace := 0

		// input a character (invalid) then an available space index (0)
		mockInput := bytes.NewBufferString(fmt.Sprintf("a\n%d\n", availableSpace))

		choiceProvider := ioHumanChoiceProvider{
			in:  bufio.NewReader(mockInput),
			out: mockOutput,
		}

		player := humanPlayer{
			PlayerInfo:     PlayerInfo{Token: 'X'},
			choiceProvider: choiceProvider,
		}

		choice := player.ChooseSpace(board)

		assert.Equal(availableSpace, choice)
	})
}
