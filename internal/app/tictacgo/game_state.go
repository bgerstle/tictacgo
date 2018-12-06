package tictacgo

type GameState string

const (
	Pending = GameState("pending")
	Victory = GameState("victory")
	Tie     = GameState("tie")
)

func isWinningVector(spaces []Space) bool {
	return spaces[0] != Space(nil) &&
		spaces[0] == spaces[1] &&
		spaces[1] == spaces[2]
}

func (board Board) GameState() (GameState, Space) {
	for _, row := range board.rows() {
		if isWinningVector(row) {
			return Victory, row[0]
		}
	}
	return Pending, nil
}
