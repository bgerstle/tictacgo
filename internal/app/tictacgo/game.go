package tictacgo

type Game struct {
	Player1 Player
	Player2 Player
	Board   Board
}

func (g Game) Play() (GameState, Player) {
	return Pending, nil
}
