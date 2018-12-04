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

func TestWelcomeMessage(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	ttgCmd := exec.Command("go", "run", "../../cmd/tictacgo/main.go")

	ttgStdout, stdoutErr := ttgCmd.StdoutPipe()
	require.Nil(stdoutErr)

	ttgStderr, stderrErr := ttgCmd.StderrPipe()
	require.Nil(stderrErr)

	combinedPipe := io.MultiReader(ttgStdout, ttgStderr)

	startErr := ttgCmd.Start()
	require.Nil(startErr)

	combinedOutLines, readAllErr := ReadAllLines(combinedPipe)
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

	lastOutputBoardLineIndex := firstOutputBoardLineIndex + len(expectedBoardLines)

	actualRestOfBoard :=
		actualBoardLines[firstOutputBoardLineIndex+1 : lastOutputBoardLineIndex]

	assert.Equal(
		expectedBoardLines[1:],
		actualRestOfBoard)

	ttgCmd.Wait()
}
