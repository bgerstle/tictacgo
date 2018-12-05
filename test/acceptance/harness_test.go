package tictacgo_test

import (
	"bufio"
	"io"
	"os/exec"
	"strings"
	"testing"

	"github.com/bgerstle/tictacgo/internal/app/tictacgo"
	"github.com/stretchr/testify/require"
)

type AppTestHarness struct {
	Cmd *exec.Cmd
	Out *bufio.Reader
	Err *bufio.Reader
	In  io.Writer
	t   *testing.T
}

func StartTicTacGo(t *testing.T) AppTestHarness {
	t.Helper()

	require := require.New(t)

	ttgCmd := exec.Command("go", "run", "../../cmd/tictacgo/main.go")

	stdout, stdoutErr := ttgCmd.StdoutPipe()
	require.Nil(stdoutErr)

	stderr, stderrErr := ttgCmd.StderrPipe()
	require.Nil(stderrErr)

	stdin, stdinErr := ttgCmd.StdinPipe()
	require.Nil(stdinErr)

	startErr := ttgCmd.Start()
	require.Nil(startErr)

	return AppTestHarness{
		Cmd: ttgCmd,
		Out: bufio.NewReader(stdout),
		Err: bufio.NewReader(stderr),
		In:  stdin,
		t:   t,
	}
}

func (testHarness AppTestHarness) ReadInitialOutput() (initialOutputLines []string) {
	require := require.New(testHarness.t)
	testHarness.t.Helper()

	expectedBoardLines := strings.Split(tictacgo.EmptyBoard().String(), "\n")
	// drop trailing empty line
	expectedBoardLines = expectedBoardLines[0 : len(expectedBoardLines)-1]
	expectedBoardLinesLen := len(expectedBoardLines)
	lastExpectedBoardLine := expectedBoardLines[expectedBoardLinesLen-1]

	initialOutputLines = []string{}
	for {
		line, lineErr := testHarness.Out.ReadString('\n')

		line = strings.TrimRight(line, "\n")

		require.Nil(lineErr)

		initialOutputLines = append(initialOutputLines, line)
		if line == lastExpectedBoardLine {
			break
		}
	}

	require.Len(
		initialOutputLines,
		expectedBoardLinesLen+1,
		"Expected output to contain welcome message plus empty board")

	return
}
