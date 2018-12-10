package main

import (
	"bufio"
	"os"

	"github.com/bgerstle/tictacgo/internal/app/tictacgo"
)

func main() {
	tictacgo.WriteWelcomeMessage(os.Stdout)

	bufStdin := bufio.NewReader(os.Stdin)

	player1 := tictacgo.NewHumanPlayer('X', bufStdin)

	player2 := tictacgo.BotPlayer{
		PlayerInfo: tictacgo.PlayerInfo{Token: 'O'},
	}

	game := tictacgo.NewGame(
		player1,
		player2,
		tictacgo.ConsoleReporter{
			Out: os.Stdout,
		},
	)

	game.Play()
}
