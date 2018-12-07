package main

import (
	"bufio"
	"os"

	"github.com/bgerstle/tictacgo/internal/app/tictacgo"
)

func main() {
	tictacgo.WriteWelcomeMessage(os.Stdout)

	bufStdin := bufio.NewReader(os.Stdin)

	player1 := tictacgo.HumanPlayer{
		PlayerInfo: tictacgo.PlayerInfo{Token: 'X'},
		ChoiceProvider: tictacgo.IOHumanChoiceProvider{
			In:  bufStdin,
			Out: os.Stdout,
		},
	}

	player2 := tictacgo.HumanPlayer{
		PlayerInfo: tictacgo.PlayerInfo{Token: 'O'},
		ChoiceProvider: tictacgo.IOHumanChoiceProvider{
			In:  bufStdin,
			Out: os.Stdout,
		},
	}

	game := tictacgo.Game{
		Player1: player1,
		Player2: player2,
		Board:   tictacgo.EmptyBoard(),
		Reporter: tictacgo.ConsoleReporter{
			Out: os.Stdout,
		},
	}

	game.Play()
}
