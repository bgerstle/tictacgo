package tictacgo

type HumanChoiceProvider interface {
	getChoice(p PlayerInfo, board Board) (int, error)
}

type HumanPlayer struct {
	PlayerInfo
	ChoiceProvider HumanChoiceProvider
}

// ChooseSpace will ask the user for their desired space, retrying on error
func (p HumanPlayer) ChooseSpace(board Board) int {
	var choice int
	for {
		c, choiceErr := p.ChoiceProvider.getChoice(p.PlayerInfo, board)
		if choiceErr == nil {
			choice = c
			break
		}
	}
	return choice
}
