package tictacgo

// BotPlayer is a Player that picks its own moves.
type BotPlayer struct {
	PlayerInfo
}

// ChooseSpace will use fancy algorithms to pick an available space on the board.
func (bp BotPlayer) ChooseSpace(b Board) int {
	return b.AvailableSpaces()[0]
}
