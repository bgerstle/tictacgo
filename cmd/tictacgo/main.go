package main

import (
	"fmt"
	"os"

	"github.com/bgerstle/tictacgo/internal/app/tictacgo"
)

func main() {
	tictacgo.WriteWelcomeMessage(os.Stdout)
	board := tictacgo.NewBoard()
	fmt.Print(board.String())
}
