package main

import (
	"os"

	"github.com/bgerstle/tictacgo/internal/app/tictacgo"
)

func main() {
	tictacgo.WriteWelcomeMessage(os.Stdout)
	board := tictacgo.Board{}
	writeErr := board.Write(os.Stdout)
	if writeErr != nil {
		panic(writeErr)
	}
}
