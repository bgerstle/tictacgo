package tictacgo

import (
	"fmt"
	"io"
)

type ConsoleReporter struct {
	Out io.Writer
}

func EndMessageForState(state GameState, winner Space) string {
	switch state {
	case Tie:
		return "No winners or losers here, it's a tie!"
	case Victory:
		return fmt.Sprintf("Player '%c' won!", *winner)
	default:
		panic(fmt.Sprintf("Only expected victory or tie, got: %s", state))
	}
}

func (cr ConsoleReporter) ReportGameStart(b Board) {
	fmt.Fprint(cr.Out, b.String())
}

func moveAnnouncement(playerToken rune, space int) string {
	return fmt.Sprintf("Player '%c' chose space %d.", playerToken, space)
}

func (cr ConsoleReporter) ReportGameProgress(b Board, playerToken rune, space int) {
	fmt.Fprintln(cr.Out, moveAnnouncement(playerToken, space))
	fmt.Fprint(cr.Out, b.String())
}

func (cr ConsoleReporter) ReportGameEnd(b Board, state GameState, winner Space) {
	fmt.Fprintln(cr.Out, EndMessageForState(state, winner))
}
