package tictacgo_test

import (
	"io"
	"io/ioutil"
	"os/exec"
	"sort"
	"strings"
	"testing"

	"github.com/bgerstle/tictacgo/internal/app/tictacgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ReadAllLines(reader io.Reader) ([]string, error) {
	allOut, allErr := ioutil.ReadAll(reader)
	if allErr != nil {
		return nil, allErr
	}
	allLines := strings.Split(string(allOut), "\n")
	allLinesLen := len(allLines)
	if allLinesLen > 1 && allLines[allLinesLen-1] == "" {
		allLines = allLines[:allLinesLen-1]
	}
	return allLines, nil
}

func StartTicTacGo(t *testing.T) (cmd *exec.Cmd, out io.Reader, in io.Writer) {
	t.Helper()

	require := require.New(t)

	ttgCmd := exec.Command("go", "run", "../../cmd/tictacgo/main.go")

	stdout, stdoutErr := ttgCmd.StdoutPipe()
	require.Nil(stdoutErr)

	stderr, stderrErr := ttgCmd.StderrPipe()
	require.Nil(stderrErr)

	stdin, stdinErr := ttgCmd.StdinPipe()
	require.Nil(stdinErr)

	combinedOut := io.MultiReader(stdout, stderr)

	startErr := ttgCmd.Start()
	require.Nil(startErr)

	return ttgCmd, combinedOut, stdin
}

func TestInitialOutput(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	ttgCmd, combinedOut, _ := StartTicTacGo(t)

	combinedOutLines, readAllErr := ReadAllLines(combinedOut)
	require.Nil(readAllErr)

	require.Condition(func() (success bool) {
		return len(combinedOutLines) > 1
	}, "Expected output to have more than 1 line")

	assert.Equal(tictacgo.WelcomeMessage, combinedOutLines[0])

	actualBoardLines := combinedOutLines[1:]

	expectedBoardLines := strings.Split(tictacgo.EmptyBoard().String(), "\n")

	firstOutputBoardLineIndex :=
		sort.SearchStrings(actualBoardLines, expectedBoardLines[0])

	require.NotEqual(
		len(actualBoardLines)-1, firstOutputBoardLineIndex,
		"Couldn't find first line of board in output")

	lastOutputBoardLineIndex := firstOutputBoardLineIndex + len(expectedBoardLines) - 1

	actualRestOfBoard :=
		actualBoardLines[firstOutputBoardLineIndex+1 : lastOutputBoardLineIndex+1]

	assert.Equal(
		expectedBoardLines[1:],
		actualRestOfBoard)

	promptLineIndex := lastOutputBoardLineIndex + 1

	require.Condition(
		func() (success bool) {
			return promptLineIndex < len(combinedOutLines)
		},
		"Expected to find prompt for user input, but ran out of output lines")

	prompt := combinedOutLines[promptLineIndex]

	assert.Equal("Enter X's move: ", prompt)

	ttgCmd.Wait()
}
