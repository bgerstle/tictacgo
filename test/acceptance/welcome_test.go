package tictacgo_test

import (
	"io"
	"io/ioutil"
	"os/exec"
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

	require.Equal(1, len(combinedOutLines))
	assert.Equal(tictacgo.WelcomeMessage, combinedOutLines[0])

	ttgCmd.Wait()
}
