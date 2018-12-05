package tictacgo_test

import (
	"io"
	"os/exec"
	"strings"
	"testing"

	"github.com/bgerstle/tictacgo/internal/app/tictacgo"
	"github.com/stretchr/testify/require"
)

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

func ReadInitialOutput(t *testing.T, out io.Reader) (initialOutputLines []string) {
	require := require.New(t)
	t.Helper()

	reader := NewUnbufferedReader(out)

	expectedBoardLines := strings.Split(tictacgo.EmptyBoard().String(), "\n")
	// drop trailing empty line
	expectedBoardLines = expectedBoardLines[0 : len(expectedBoardLines)-1]
	expectedBoardLinesLen := len(expectedBoardLines)
	lastExpectedBoardLine := expectedBoardLines[expectedBoardLinesLen-1]

	initialOutputLines, error := reader.ReadLinesUntil(
		lastExpectedBoardLine,
		expectedBoardLinesLen+1,
	)
	require.Nil(error)
	require.Len(
		initialOutputLines,
		expectedBoardLinesLen+1,
		"Expected output to contain welcome message plus empty board")
	return
}
