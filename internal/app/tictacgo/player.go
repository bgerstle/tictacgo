package tictacgo

type PlayerInfo struct {
	Token rune
}

type Player interface {
	Info() PlayerInfo
	ChooseSpace(board Board) int
}
