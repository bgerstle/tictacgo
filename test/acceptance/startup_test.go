package tictacgo_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/bgerstle/tictacgo/internal/app/tictacgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitialOutput(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	ttgCmd, combinedOut, _ := StartTicTacGo(t)

	expectedBoardLines := strings.Split(tictacgo.EmptyBoard().String(), "\n")
	// drop trailing empty line
	expectedBoardLines = expectedBoardLines[0 : len(expectedBoardLines)-1]

	combinedOutLines := ReadInitialOutput(t, combinedOut)

	assert.Equal(tictacgo.WelcomeMessage, combinedOutLines[0])

	actualBoardLines := combinedOutLines[1 : len(expectedBoardLines)+1]

	require.Equal(expectedBoardLines, actualBoardLines)

	killErr := ttgCmd.Process.Kill()

	require.Nil(killErr)
}

func TestEnterMove(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	ttgCmd, combinedOut, in := StartTicTacGo(t)

	ReadInitialOutput(t, combinedOut)

	expectedPrompt := fmt.Sprintf(tictacgo.PlayerMovePromptf, 'X')

	reader := NewUnbufferedReader(combinedOut)

	actualPromptBytes, readError := reader.ReadBytes(len([]byte(expectedPrompt)))

	require.Nil(readError)

	actualPrompt := string(actualPromptBytes)

	assert.Equal(expectedPrompt, actualPrompt)

	fmt.Fprintln(in, "4")

	tempResponse, tempResponseErr := reader.ReadLine()

	require.Nil(tempResponseErr)

	assert.Equal("X chose space 4", tempResponse)

	ttgCmd.Wait()
}
