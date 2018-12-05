package tictacgo

type humanChoiceProvider interface {
	getChoice(p PlayerInfo, board Board) (int, error)
}

type humanPlayer struct {
	PlayerInfo
	choiceProvider humanChoiceProvider
}

func (hp humanPlayer) Info() PlayerInfo {
	return hp.PlayerInfo
}

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
