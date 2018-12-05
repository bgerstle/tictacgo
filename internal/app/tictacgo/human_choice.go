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
			in:  os.Stdin,
			out: os.Stdout,
		},
	}
	return hp
}

type ioHumanChoiceProvider struct {
	in  io.Reader
	out io.Writer
}

const PlayerMovePromptf = "Make your move, %c: "

func (cp ioHumanChoiceProvider) getChoice(p PlayerInfo, board Board) (int, error) {
	reader := bufio.NewReader(cp.in)
	fmt.Fprintf(cp.out, PlayerMovePromptf, p.Token)
	input, readErr := reader.ReadString(byte('\n'))
	if readErr != nil {
		fmt.Fprintf(os.Stderr, "Player input encountered error %s", readErr.Error())
		fmt.Fprintln(cp.out, "Oops! Let's try again...")
		return -1, readErr
	}
	choice, atoiErr := strconv.Atoi(strings.TrimSpace(input))
	if atoiErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to convert user input to int: %s", readErr.Error())
		fmt.Fprintln(cp.out, "Choices must be one of the available spaces on the board. Let's try again.")
		return -1, atoiErr
	}
	return choice, nil
}
