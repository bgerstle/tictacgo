package tictacgo

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type ioHumanChoiceProvider struct {
	// Must use the same buffered reader to avoid losing input
	// b/t calls to getChoice
	In  *bufio.Reader
	Out io.Writer
}

// PlayerMovePromptf is a format string accepting a player's token.
const PlayerMovePromptf = "Make your move, %c: "

// Internal type used for getting choices from stdin/stdout, and writing error messages
func (cp ioHumanChoiceProvider) getChoice(p PlayerInfo, board Board) (int, error) {
	fmt.Fprintf(cp.Out, PlayerMovePromptf, p.Token)
	input, readErr := cp.In.ReadString(byte('\n'))

	if readErr != nil {
		// TODO: have some logger that writes to a file...
		// fmt.Fprintf(os.Stderr, "Player input encountered error %s", readErr.Error())
		fmt.Fprintln(cp.Out, "Oops, something went wrong! Let's try again...")
		return -1, readErr
	}
	choice, atoiErr := strconv.Atoi(strings.TrimSpace(input))
	if atoiErr != nil {
		// fmt.Fprintf(os.Stderr, "Failed to convert user input to int: %s", atoiErr.Error())
		fmt.Fprintln(cp.Out, "That doesn't look like a number, please enter one of the available spaces.")
		return -1, atoiErr
	}
	if choice < 0 || choice >= board.SpacesLen() {
		fmt.Fprintln(cp.Out, "Please choose one of the available spaces.")
		return choice, fmt.Errorf("Choice %d out of bounds", choice)
	}
	if !board.IsSpaceAvailable(choice) {
		fmt.Fprintln(cp.Out, fmt.Sprintf("Space '%d' is already taken, please enter one of the spaces on the board.", choice))
		return choice, fmt.Errorf("Choice %d already taken", choice)
	}
	return choice, nil
}
