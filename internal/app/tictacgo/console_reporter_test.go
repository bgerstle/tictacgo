package tictacgo

import (
	"bytes"
	"fmt"
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

	for _, space := range shuffledSpaces[:numMoves] {
		var state GameState
		aab.Board, state, _ = aab.Board.AssignSpace(space)
		if state != Pending {
			break
		}
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

		reporter.ReportGameStart(NewEmptyTestBoard())

		actualOutput := mockOutput.String()
		assert.Equal(NewEmptyTestBoard().String(), actualOutput)
	})

	t.Run("Progress prints move, then board", func(t *testing.T) {
		assert := assert.New(t)

		assert.NoError(quick.Check(func(aab ArbitraryAnyBoard) bool {
			mockOutput, reporter := setup()

			// last two args aren't used, so passing dummy values
			reporter.ReportGameProgress(aab.Board, aab.Board.ActivePlayerToken(), 0)

			actualOutput := mockOutput.String()
			return assert.Equal(
				moveAnnouncement(aab.Board.ActivePlayerToken(), 0)+"\n"+aab.Board.String(),
				actualOutput,
			)
		}, nil))
	})

	t.Run("End prints board and state message", func(t *testing.T) {
		assert := assert.New(t)

		assert.NoError(quick.Check(func(avb ArbitraryVictoryBoard) bool {
			mockOutput, reporter := setup()

			state, winner := avb.GameState()
			reporter.ReportGameEnd(avb.Board, state, winner)

			actualOutput := mockOutput.String()
			return assert.Equal(EndMessageForState(state, winner)+"\n", actualOutput)
		}, nil))
	})

	t.Run("Prints tie message for a tie", func(t *testing.T) {
		assert := assert.New(t)

		mockOutput, reporter := setup()

		board := NewTestBoardWithSpaces([]Space{
			X, O, X,
			O, X, X,
			O, X, O,
		})
		state, winner := board.GameState()

		reporter.ReportGameEnd(board, state, winner)

		actualOutput := mockOutput.String()

		assert.Equal(EndMessageForState(state, winner)+"\n", actualOutput)
	})
}

func ExampleEndMessageForState() {
	fmt.Println(EndMessageForState(Tie, Space(nil)))
	fmt.Println(EndMessageForState(Victory, X))
	// Output:
	// No winners or losers here, it's a tie!
	// Player 'X' won!
}
