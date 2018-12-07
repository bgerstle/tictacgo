package tictacgo

// PlayerInfo describes a player, including their associated token (e.g. 'X', 'O')
type PlayerInfo struct {
	Token rune
}

// Info will return the receiver. Prevents Player types that embed PlayerInfo from having to implement it.
func (pi PlayerInfo) Info() PlayerInfo {
	return pi
}

// Player is a type that has PlayerInfo and can choose a space on the board.
type Player interface {
	Info() PlayerInfo
	ChooseSpace(board Board) int
}
