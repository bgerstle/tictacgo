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
	for _, column := range board.columns() {
		if isWinningVector(column) {
			return Victory, column[0]
		}
	}
	for _, diagonal := range board.diagonals() {
		if isWinningVector(diagonal) {
			return Victory, diagonal[0]
		}
	}
	if len(board.AvailableSpaces()) == 0 {
		return Tie, nil
	}
	return Pending, nil
}
