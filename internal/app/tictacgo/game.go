package tictacgo

import "fmt"

// GameReporter is an interface for types that report game updates.
type GameReporter interface {
	ReportGameStart(b Board)
	ReportGameProgress(b Board, lastPlayerToken rune, lastPlayerSpace int)
	ReportGameEnd(finalBoard Board, state GameState, winner Space)
}

// Game is a type that contains the current board, its players, and a reporter.
// This is one of the highest-level objects in the app, as it has to orchestrate
// the logical flow of the game.
type Game struct {
	Player1  Player
	Player2  Player
	Board    Board
	Reporter GameReporter
}

// PlayerForToken returns a reference to the player with a matching token.
func (g Game) PlayerForSpace(s Space) *Player {
	if s == nil {
		return nil
	}
	if g.Player1.Info().Token == *s {
		return &(g.Player1)
	} else if g.Player2.Info().Token == *s {
		return &(g.Player2)
	}
	panic(fmt.Sprintf("Asked for player with unknown token: %s", spaceToString(s, "null")))
}

// Play will start the Tic Tac Toe game, asking each player to choose a spot
// on the board in turn, assigning it on the board, and repeating until
// the game state isn't pending.
func (g *Game) Play() (GameState, *Player) {
	state, winningSpace := g.Board.GameState()

	if state != Pending {
		return state, g.PlayerForSpace(winningSpace)
	}

	if g.Reporter != nil {
		g.Reporter.ReportGameStart(g.Board)
	}

	var winningPlayer *Player
	state, winningPlayer = g.takeTurn()

	for {
		state, winningPlayer = g.takeTurn()
		if state != Pending {
			break
		}
	}

	return state, winningPlayer
}

func (g *Game) takeTurn() (GameState, *Player) {
	currentToken := g.Board.ActivePlayerToken()
	currentPlayer := g.PlayerForSpace(&currentToken)

	// ask current player to choose their space
	space := (*currentPlayer).ChooseSpace(g.Board)

	// assign it on the board
	var (
		state        GameState
		winningSpace Space
	)
	g.Board, state, winningSpace = g.Board.AssignSpace(space)

	// report progress
	if g.Reporter != nil {
		g.Reporter.ReportGameProgress(g.Board, currentToken, space)
		if state != Pending {
			g.Reporter.ReportGameEnd(g.Board, state, winningSpace)
		}
	}

	return state, g.PlayerForSpace(winningSpace)
}
