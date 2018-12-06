package tictacgo

import "fmt"

type Game struct {
	Player1 Player
	Player2 Player
	Board   Board
}

func (g Game) PlayerForToken(t rune) *Player {
	if g.Player1.Info().Token == t {
		return &(g.Player1)
	} else if g.Player2.Info().Token == t {
		return &(g.Player2)
	}
	panic(fmt.Sprintf("Asked for player with unknown token: %c", t))
}

func (g *Game) Play() (GameState, *Player) {
	p1MoveCount := len(g.Board.SpacesAssignedTo(g.Player1.Info().Token))
	p2MoveCount := len(g.Board.SpacesAssignedTo(g.Player2.Info().Token))

	var (
		evenPlayer *Player
		oddPlayer  *Player
	)
	if p1MoveCount >= p2MoveCount {
		evenPlayer = &g.Player1
		oddPlayer = &g.Player2
	} else {
		evenPlayer = &g.Player2
		oddPlayer = &g.Player1
	}

	var (
		state        GameState
		winningSpace Space
	)
	for i := range g.Board.AvailableSpaces() {
		// get current player for this turn
		var currentPlayer *Player
		if i%2 == 0 {
			currentPlayer = evenPlayer
		} else {
			currentPlayer = oddPlayer
		}
		currentToken := (*currentPlayer).Info().Token

		// ask them to choose their space
		space := (*currentPlayer).ChooseSpace(g.Board)

		// assign it on the board
		newBoard := g.Board.AssignSpace(space, &currentToken)

		// apply new board to the game
		g.Board = newBoard

		// check if game is over. if so, end the game
		state, winningSpace = g.Board.GameState()
		if state != Pending {
			break
		}
	}

	var winningPlayer *Player
	if winningSpace != nil {
		winningPlayer = g.PlayerForToken(*winningSpace)
	}

	return state, winningPlayer
}
