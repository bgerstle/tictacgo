package tictacgo

type game struct {
	Player1 Player
	Player2 Player
	board   Board
}

type GameState string

const (
	Pending = GameState("pending")
	Victory = GameState("victory")
	Tie     = GameState("tie")
)

func (g game) state() (state GameState, winner Player) {
	return Pending, nil
}
