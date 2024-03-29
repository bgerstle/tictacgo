package tictacgo

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHumanChoiceProviderIntegration(t *testing.T) {
	t.Run("returns choice specified in input", func(t *testing.T) {
		assert := assert.New(t)

		board := NewEmptyTestBoard()

		for spaceNum := range board.AvailableSpaces() {
			mockOutput := &bytes.Buffer{}
			mockInput := bufio.NewReader(bytes.NewBufferString(fmt.Sprintf("%d\n", spaceNum)))

			choiceProvider := ioHumanChoiceProvider{
				In:  mockInput,
				Out: mockOutput,
			}

			player := HumanPlayer{
				PlayerInfo:     PlayerInfo{Token: 'X'},
				choiceProvider: choiceProvider,
			}

			choice := player.ChooseSpace(board)

			assert.Equal(spaceNum, choice)
		}
	})

	t.Run("returns second, valid choice after invalid choice", func(t *testing.T) {
		assert := assert.New(t)

		board := NewEmptyTestBoard()

		mockOutput := &bytes.Buffer{}

		availableSpace := 0

		// input a character (invalid) then an available space index (0)
		mockInput := bytes.NewBufferString(fmt.Sprintf("a\n%d\n", availableSpace))

		choiceProvider := ioHumanChoiceProvider{
			In:  bufio.NewReader(mockInput),
			Out: mockOutput,
		}

		player := HumanPlayer{
			PlayerInfo:     PlayerInfo{Token: 'X'},
			choiceProvider: choiceProvider,
		}

		choice := player.ChooseSpace(board)

		assert.Equal(availableSpace, choice)
	})
}
