package tictacgo

import (
	"fmt"
	"io"
)

// ConsoleReporter is a GameReporter that will write game event reports to its Out field.
type ConsoleReporter struct {
	Out io.Writer
}

// EndMessageForState will return a string based on the result of the game and, if appropriate, its winner.
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

// ReportGameStart will print the board.
func (cr ConsoleReporter) ReportGameStart(b Board) {
	fmt.Fprint(cr.Out, b.String())
}

// returns a string announcing the player's move
func moveAnnouncement(playerToken rune, space int) string {
	return fmt.Sprintf("Player '%c' chose space %d.", playerToken, space)
}

// ReportGameProgress will print the move the specified player just took and the new board.
func (cr ConsoleReporter) ReportGameProgress(b Board, playerToken rune, space int) {
	fmt.Fprintln(cr.Out, moveAnnouncement(playerToken, space))
	fmt.Fprint(cr.Out, b.String())
}

// ReportGameEnd will print the end message according to the state and winner.
func (cr ConsoleReporter) ReportGameEnd(b Board, state GameState, winner Space) {
	fmt.Fprintln(cr.Out, EndMessageForState(state, winner))
}
