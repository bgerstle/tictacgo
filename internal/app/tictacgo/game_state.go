package tictacgo

// GameState describes the current state of the game. See constants defined below.
type GameState string

const (
	// Pending means the game is still going on.
	Pending = GameState("pending")

	// Victory indicates that one of the players has won.
	Victory = GameState("victory")

	// Tie happens when the board has been filled without either player winning.
	Tie = GameState("tie")
)

func isWinningVector(spaces []Space) bool {
	return spaces[0] != nil &&
		spaces[1] != nil &&
		spaces[2] != nil &&
		*spaces[0] == *spaces[1] &&
		*spaces[1] == *spaces[2]
}

// GameState retrieves the GameState from the specified board, along with a winner if there is one.
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
