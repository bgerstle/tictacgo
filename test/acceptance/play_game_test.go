package tictacgo_test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/bgerstle/tictacgo/internal/app/tictacgo"

	"github.com/stretchr/testify/assert"
)

func TestExamplePvPGame(t *testing.T) {
	assert := assert.New(t)

	testHarness := StartTicTacGo(t)

	testHarness.ReadInitialOutput()

	/*
		play the game:
			X, X, X,
			O, nil, nil,
			O, nil, nil,
	*/
	for _, space := range []int{0, 4, 8} {
		fmt.Fprintln(testHarness.In, fmt.Sprintf("%d", space))
	}

	programOutput, readErr := ioutil.ReadAll(testHarness.Out)
	assert.NoError(readErr)

	outputStr := string(programOutput)
	outputLines := strings.Split(outputStr, "\n")
	lastLine := outputLines[len(outputLines)-2]

	expectedWinningToken := 'X'
	expectedLastLine := tictacgo.EndMessageForState(tictacgo.Victory, &expectedWinningToken)
	assert.Equal(expectedLastLine, lastLine)

	testHarness.Cmd.Wait()
}
