package tictacgo

type humanChoiceProvider interface {
	getChoice(p PlayerInfo, board Board) (int, error)
}

type humanPlayer struct {
	PlayerInfo
	choiceProvider humanChoiceProvider
}

// Info will return the info for this player
func (p humanPlayer) Info() PlayerInfo {
	return p.PlayerInfo
}

// ChooseSpace will ask the user for their desired space, retrying on error
func (p humanPlayer) ChooseSpace(board Board) int {
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
