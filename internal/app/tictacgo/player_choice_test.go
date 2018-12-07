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

	expectedPrompt := fmt.Sprintf(PlayerMovePromptf, player.Token)

	// Read until : in prompt, not newline, otherwise we could accidentally pick up the
	// error message
	actualPrompt, err := reader.ReadString(':')

	require.Nil(err)

	assert.Condition(func() bool {
		return strings.HasPrefix(expectedPrompt, actualPrompt)
	})
}

func TestHumanChoiceProviderValidInputs(t *testing.T) {
	assert := assert.New(t)

	board := EmptyBoard()

	for spaceNum := range board.AvailableSpaces() {
		mockOutput := &bytes.Buffer{}
		mockInput := bytes.NewBufferString(fmt.Sprintf("%d\n", spaceNum))

		choiceProvider := ioHumanChoiceProvider{
			In:  bufio.NewReader(mockInput),
			Out: mockOutput,
		}

		playerInfo := PlayerInfo{Token: 'X'}

		choice, err := choiceProvider.getChoice(playerInfo, board)

		assert.Nil(err)

		assertExpectedMovePrompt(t, bufio.NewReader(mockOutput), playerInfo)

		assert.Equal(spaceNum, choice)
	}
}

func TestHumanChoiceProviderInvalidInputs(t *testing.T) {
	invalidInputExamples := []string{
		"",
		"a\n",
		"abc\n",
		"10000\n",
		"-1\n",
	}
	for i, invalidInput := range invalidInputExamples {
		t.Run(fmt.Sprintf("invalid input %d", i), func(t *testing.T) {
			assert := assert.New(t)

			board := EmptyBoard()

			mockOutput := &bytes.Buffer{}
			mockInput := bytes.Buffer{}

			fmt.Fprint(&mockInput, invalidInput)

			choiceProvider := ioHumanChoiceProvider{
				In:  bufio.NewReader(&mockInput),
				Out: mockOutput,
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
			In:  bufio.NewReader(mockInput),
			Out: mockOutput,
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
