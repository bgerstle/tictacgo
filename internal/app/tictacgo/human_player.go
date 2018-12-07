package tictacgo

import (
	"bufio"
	"os"
)

type HumanChoiceProvider interface {
	getChoice(p PlayerInfo, board Board) (int, error)
}

type HumanPlayer struct {
	PlayerInfo
	choiceProvider HumanChoiceProvider
}

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
