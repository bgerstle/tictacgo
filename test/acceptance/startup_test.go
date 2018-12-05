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

	testHarness := StartTicTacGo(t)

	expectedBoardLines := strings.Split(tictacgo.EmptyBoard().String(), "\n")
	// drop trailing empty line
	expectedBoardLines = expectedBoardLines[0 : len(expectedBoardLines)-1]

	combinedOutLines := testHarness.ReadInitialOutput()

	assert.Equal(tictacgo.WelcomeMessage, combinedOutLines[0])

	actualBoardLines := combinedOutLines[1 : len(expectedBoardLines)+1]

	require.Equal(expectedBoardLines, actualBoardLines)

	killErr := testHarness.Cmd.Process.Kill()

	require.Nil(killErr)
}

func TestEnterMove(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	testHarness := StartTicTacGo(t)

	testHarness.ReadInitialOutput()

	expectedPrompt := fmt.Sprintf(tictacgo.PlayerMovePromptf, 'X')

	expectedNumBytes := len([]byte(expectedPrompt))

	promptBuf := make([]byte, expectedNumBytes)
	readBytes, readError := testHarness.Out.Read(promptBuf)

	require.Nil(readError)
	require.Equal(expectedNumBytes, readBytes)

	actualPrompt := string(promptBuf)

	assert.Equal(expectedPrompt, actualPrompt)

	fmt.Fprintln(testHarness.In, "4")

	tempResponse, tempResponseErr := testHarness.Out.ReadString('\n')

	require.Nil(tempResponseErr)

	assert.Equal("X chose space 4\n", tempResponse)

	testHarness.Cmd.Process.Kill()
}
