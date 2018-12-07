package tictacgo

type BotPlayer struct {
	PlayerInfo
}

func (bp BotPlayer) ChooseSpace(b Board) int {
	return b.AvailableSpaces()[0]
}
