package tictacgo_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/bgerstle/tictacgo/internal/app/tictacgo"

	"github.com/stretchr/testify/assert"
)

func TestExamplePvPGame(t *testing.T) {
	assert := assert.New(t)

	testHarness := StartTicTacGo(t)

	/*
		play the game:
			X, X, X,
			O, nil, nil,
			O, nil, nil,
	*/
	for _, space := range []int{0, 3, 1, 6, 2} {
		fmt.Fprintf(testHarness.In, "%d\n", space)
	}

	programOutput, readErr := ioutil.ReadAll(testHarness.Out)
	assert.NoError(readErr)

	lastLine := programOutput[len(programOutput)-1]

	expectedWinningToken := 'X'
	assert.Equal(tictacgo.EndMessageForState(tictacgo.Victory, &expectedWinningToken), lastLine)

	testHarness.Cmd.Wait()
}
