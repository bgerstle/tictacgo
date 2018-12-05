package tictacgo

// PlayerInfo describes a player, including their associated token (e.g. 'X', 'O')
type PlayerInfo struct {
	Token rune
}

// Player is a type that has PlayerInfo and can choose a space on the board.
type Player interface {
	Info() PlayerInfo
	ChooseSpace(board Board) int
}
