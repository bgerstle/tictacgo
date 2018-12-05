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

	playerX := tictacgo.NewHumanPlayer(tictacgo.PlayerInfo{Token: 'X'})
	xsMove := playerX.ChooseSpace(board)

	fmt.Printf("X chose space %d", xsMove)
	fmt.Println()
}
