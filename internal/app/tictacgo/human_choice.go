package tictacgo

import (
	"bufio"
	"errors"
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

const InputErrorMessage = "Oops, something went wrong! Let's try again..."

const InputNotNumberErrorMessage = "That doesn't look like a number. Please enter one of the available spaces."

func ErrorMessageForOutOfBoundsSpace(space int) string {
	return fmt.Sprintf("Choice %d out of bounds", space)
}

func ErrorMessageForTakenSpace(space int) string {
	return fmt.Sprintf("Space '%d' is already taken. Please enter one of the spaces on the board.", space)
}

// Internal type used for getting choices from stdin/stdout, and writing error messages
func (cp ioHumanChoiceProvider) getChoice(p PlayerInfo, board Board) (int, error) {
	fmt.Fprintf(cp.Out, PlayerMovePromptf, p.Token)
	input, readErr := cp.In.ReadString(byte('\n'))

	if readErr != nil {
		// TODO: have some logger that writes to a file...
		// fmt.Fprintf(os.Stderr, "Player input encountered error %s", readErr.Error())
		fmt.Fprintln(cp.Out)
		fmt.Fprintln(cp.Out, InputErrorMessage)
		return -1, readErr
	}
	choice, atoiErr := strconv.Atoi(strings.TrimSpace(input))
	if atoiErr != nil {
		// fmt.Fprintf(os.Stderr, "Failed to convert user input to int: %s", atoiErr.Error())
		fmt.Fprintln(cp.Out, InputNotNumberErrorMessage)
		return -1, atoiErr
	}
	if choice < 0 || choice >= board.SpacesLen() {
		msg := ErrorMessageForOutOfBoundsSpace(choice)
		fmt.Fprintln(cp.Out, msg)
		return choice, errors.New(msg)
	}
	if !board.IsSpaceAvailable(choice) {
		msg := ErrorMessageForTakenSpace(choice)
		fmt.Fprintln(cp.Out, msg)
		return choice, fmt.Errorf(msg)
	}
	return choice, nil
}
