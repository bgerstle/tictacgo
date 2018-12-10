package tictacgo_test

import (
	"bufio"
	"io"
	"os"
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

	teedStdout := io.TeeReader(stdout.(io.Reader), os.Stdout)

	stderr, stderrErr := ttgCmd.StderrPipe()
	require.Nil(stderrErr)

	teedStderr := io.TeeReader(stderr.(io.Reader), os.Stdout)

	stdin, stdinErr := ttgCmd.StdinPipe()
	require.Nil(stdinErr)

	startErr := ttgCmd.Start()
	require.Nil(startErr)

	return AppTestHarness{
		Cmd: ttgCmd,
		Out: bufio.NewReader(teedStdout),
		Err: bufio.NewReader(teedStderr),
		In:  stdin,
		t:   t,
	}
}

func (testHarness AppTestHarness) ReadBoard() []string {
	testHarness.t.Helper()
	require := require.New(testHarness.t)

	expectedBoardLines := strings.Split(tictacgo.NewBoard([2]tictacgo.PlayerInfo{
		tictacgo.PlayerInfo{Token: 'X'},
		tictacgo.PlayerInfo{Token: 'O'},
	}).String(), "\n")
	// drop trailing empty line
	expectedBoardLines = expectedBoardLines[0 : len(expectedBoardLines)-1]
	expectedBoardLinesLen := len(expectedBoardLines)

	boardLines := []string{}
	for i := 0; i < expectedBoardLinesLen; i++ {
		line, lineErr := testHarness.Out.ReadString('\n')
		require.NoError(lineErr)

		line = strings.TrimRight(line, "\n")

		boardLines = append(boardLines, line)
	}

	return boardLines
}

func (testHarness AppTestHarness) ReadInitialOutput() (initialOutputLines []string) {
	testHarness.t.Helper()
	require := require.New(testHarness.t)

	welcomeLine, welcomeErr := testHarness.Out.ReadString('\n')
	require.NoError(welcomeErr)
	welcomeLine = strings.TrimRight(welcomeLine, "\n")
	require.Equal(tictacgo.WelcomeMessage, welcomeLine)

	boardLines := testHarness.ReadBoard()

	return append([]string{welcomeLine}, boardLines...)
}
