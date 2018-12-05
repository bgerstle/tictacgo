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

	playerX := tictacgo.Player{Token: 'X'}
	xsMove := playerX.ChooseSpace(os.Stdout, os.Stdin, board)
	fmt.Printf("X chose space %d", xsMove)
}
