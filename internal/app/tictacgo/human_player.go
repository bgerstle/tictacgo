package tictacgo

import (
	"bufio"
	"os"
)

type humanChoiceProvider interface {
	getChoice(p PlayerInfo, board Board) (int, error)
}

// HumanPlayer is a Player that picks its moves by reading user input.
type HumanPlayer struct {
	PlayerInfo
	choiceProvider humanChoiceProvider
}

// NewHumanPlayer is a factory method that creates a new HumanPlayer
// with the given buffered input. The input must be buffered so as to
// prevent other readers from "stealing" input meant for this HumanPlayer.
func NewHumanPlayer(token rune, stdin *bufio.Reader) HumanPlayer {
	return HumanPlayer{
		PlayerInfo: PlayerInfo{Token: token},
		choiceProvider: ioHumanChoiceProvider{
			In:  stdin,
			Out: os.Stdout,
		},
	}
}

// ChooseSpace will ask the user for their desired space, retrying on error
func (p HumanPlayer) ChooseSpace(board Board) int {
	var choice int
	for {
		c, choiceErr := p.choiceProvider.getChoice(p.PlayerInfo, board)
		if choiceErr == nil {
			choice = c
			break
		}
	}
	return choice
}
