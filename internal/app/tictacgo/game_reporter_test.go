package tictacgo

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameReporterOutput(t *testing.T) {
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

	// board = board.AssignSpace(0, x)
	// expectedOutputs = append(expectedOutputs, board.String())

	// board = board.AssignSpace(1, x)
	// state := Victory
	// winner := x
	// expectedOutputs = append(expectedOutputs, "Player %c won!")
}
