package main

import (
	"fmt"
	"os"

	"github.com/bgerstle/tictacgo/internal/app/tictacgo"
)

func main() {
	tictacgo.WriteWelcomeMessage(os.Stdout)

	board := tictacgo.EmptyBoard()
	fmt.Print(board.String())

	player1 := tictacgo.NewHumanPlayer(tictacgo.PlayerInfo{Token: 'X'})
	player2 := tictacgo.NewHumanPlayer(tictacgo.PlayerInfo{Token: 'O'})

	chooseMove := func(p tictacgo.Player) {
		move := p.ChooseSpace(board)
		fmt.Println(fmt.Sprintf("%c chose space %d", p.Info().Token, move))
	}

	chooseMove(player1)
	chooseMove(player2)
}
