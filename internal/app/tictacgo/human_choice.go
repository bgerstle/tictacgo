package tictacgo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func NewHumanPlayer(info PlayerInfo) humanPlayer {
	hp := humanPlayer{
		PlayerInfo: info,
		choiceProvider: ioHumanChoiceProvider{
			in:  bufio.NewReader(os.Stdin),
			out: os.Stdout,
		},
	}
	return hp
}

type ioHumanChoiceProvider struct {
	// Must use the same buffered reader to avoid losing input
	// b/t calls to getChoice
	in  *bufio.Reader
	out io.Writer
}

const PlayerMovePromptf = "Make your move, %c: "

func (cp ioHumanChoiceProvider) getChoice(p PlayerInfo, board Board) (int, error) {
	fmt.Fprintf(cp.out, PlayerMovePromptf, p.Token)
	input, readErr := cp.in.ReadString(byte('\n'))

	if readErr != nil {
		// TODO: have some logger that writes to a file...
		// fmt.Fprintf(os.Stderr, "Player input encountered error %s", readErr.Error())
		fmt.Fprintln(cp.out, "Oops, something went wrong! Let's try again...")
		return -1, readErr
	}
	choice, atoiErr := strconv.Atoi(strings.TrimSpace(input))
	if atoiErr != nil {
		// fmt.Fprintf(os.Stderr, "Failed to convert user input to int: %s", atoiErr.Error())
		fmt.Fprintln(cp.out, "That doesn't look like a number, please enter one of the available spaces.")
		return -1, atoiErr
	}
	if choice < 0 || choice >= board.SpacesLen() {
		fmt.Fprintln(cp.out, "Please choose one of the available spaces.")
		return choice, fmt.Errorf("Choice %d out of bounds", choice)
	}
	if !board.IsSpaceAvailable(choice) {
		fmt.Fprintln(cp.out, fmt.Sprintf("Space '%d' is already taken, please enter one of the spaces on the board.", choice))
		return choice, fmt.Errorf("Choice %d already taken", choice)
	}
	return choice, nil
}
