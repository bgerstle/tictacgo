package tictacgo

import (
	"bytes"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
)

type ArbitraryAnyBoard struct{ ArbitraryBoard }

func (aab ArbitraryAnyBoard) Generate(rand *rand.Rand, size int) reflect.Value {
	aab.ArbitraryBoard = aab.ArbitraryBoard.Generate(rand, size).Interface().(ArbitraryBoard)

	shuffledSpaces := rand.Perm(aab.Board.SpacesLen())

	numMoves := size % len(shuffledSpaces)

	for i, space := range shuffledSpaces[:numMoves] {
		var currentToken rune
		if i%2 == 0 {
			currentToken = aab.Player1Token
		} else {
			currentToken = aab.Player2Token
		}
		aab.Board = aab.Board.AssignSpace(space, &currentToken)
	}
	return reflect.ValueOf(aab)
}

func TestConsoleGameReporter(t *testing.T) {
	// expectedOutputs := string[]{}

	setup := func() (mockOutput *bytes.Buffer, reporter *ConsoleReporter) {
		t.Helper()
		mockOutput = &bytes.Buffer{}
		reporter = &ConsoleReporter{Out: mockOutput}
		return
	}

	t.Run("Start prints empty board", func(t *testing.T) {
		assert := assert.New(t)

		mockOutput, reporter := setup()

		reporter.ReportGameStart(EmptyBoard())

		actualOutput := mockOutput.String()
		assert.Equal(EmptyBoard().String(), actualOutput)
	})

	t.Run("Progress prints board", func(t *testing.T) {
		assert := assert.New(t)

		assert.NoError(quick.Check(func(aab ArbitraryAnyBoard) bool {
			mockOutput, reporter := setup()

			// last two args aren't used, so passing dummy values
			reporter.ReportGameProgress(aab.Board, aab.Player1Token, 0)

			actualOutput := mockOutput.String()
			return assert.Equal(aab.Board.String(), actualOutput)
		}, nil))
	})

	t.Run("End prints board and state message", func(t *testing.T) {
		assert := assert.New(t)

		assert.NoError(quick.Check(func(afb ArbitraryFullBoard) bool {
			mockOutput, reporter := setup()

			state, winner := afb.GameState()
			reporter.ReportGameEnd(afb.Board, state, winner)

			actualOutput := mockOutput.String()
			return assert.Equal(afb.Board.String()+EndMessageForState(state, winner)+"\n", actualOutput)
		}, nil))
	})
}
