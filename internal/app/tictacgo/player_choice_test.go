package tictacgo

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertExpectedMovePrompt(t *testing.T, reader *bufio.Reader, player PlayerInfo) {
	require := require.New(t)
	assert := assert.New(t)
	t.Helper()

	prompt, err := reader.ReadString('\n')

	require.Nil(err)

	assert.Equal(fmt.Sprintf(PlayerMovePromptf, player.Token), strings.TrimRight(prompt, "\n"))
}

func TestIOHumanChoiceProvider(t *testing.T) {
	t.Run("valid inputs return the corresponding int", func(t *testing.T) {
		assert := assert.New(t)

		board := EmptyBoard()

		for spaceNum := range board.AvailableSpaces() {
			mockOutput := &bytes.Buffer{}
			mockInput := bytes.NewBufferString(fmt.Sprintf("%d\n", spaceNum))

			choiceProvider := ioHumanChoiceProvider{
				in:  bufio.NewReader(mockInput),
				out: mockOutput,
			}

			playerInfo := PlayerInfo{Token: 'X'}

			choice, err := choiceProvider.getChoice(playerInfo, board)

			assert.Nil(err)

			assertExpectedMovePrompt(t, bufio.NewReader(mockOutput), playerInfo)

			assert.Equal(spaceNum, choice)
		}
	})

	invalidInputExamples := []string{
		"",
		"a\n",
		"abc\n",
		"10000\n",
		"-1\n",
	}
	for _, invalidInput := range invalidInputExamples {
		t.Run("invalid inputs return an error", func(t *testing.T) {
			assert := assert.New(t)

			board := EmptyBoard()

			mockOutput := &bytes.Buffer{}
			mockInput := bytes.Buffer{}

			fmt.Fprint(&mockInput, invalidInput)

			choiceProvider := ioHumanChoiceProvider{
				in:  bufio.NewReader(&mockInput),
				out: mockOutput,
			}

			playerInfo := PlayerInfo{Token: 'X'}

			_, err := choiceProvider.getChoice(playerInfo, board)

			assert.NotNil(err)

			reader := bufio.NewReader(mockOutput)

			assertExpectedMovePrompt(t, reader, playerInfo)

			errorOutput, errorOutputReadErr := reader.ReadString('\n')

			assert.Nil(errorOutputReadErr)

			assert.Condition(func() (success bool) {
				return len(errorOutput) > 0
			})
		})
	}

	t.Run("taken spaces return an error", func(t *testing.T) {
		assert := assert.New(t)

		playerInfo := PlayerInfo{Token: 'X'}

		board := EmptyBoard().AssignSpace(0, &playerInfo.Token)

		mockOutput := &bytes.Buffer{}
		mockInput := bytes.NewBufferString("0\n")

		choiceProvider := ioHumanChoiceProvider{
			in:  bufio.NewReader(mockInput),
			out: mockOutput,
		}

		_, err := choiceProvider.getChoice(playerInfo, board)

		assert.NotNil(err)

		reader := bufio.NewReader(mockOutput)

		assertExpectedMovePrompt(t, reader, playerInfo)

		errorOutput, errorOutputReadErr := reader.ReadString('\n')

		assert.Nil(errorOutputReadErr)

		assert.Condition(func() (success bool) {
			return len(errorOutput) > 0
		})
	})
}
