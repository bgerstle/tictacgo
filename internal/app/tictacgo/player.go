package tictacgo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Player struct {
	Token rune
}

const PlayerMovePromptf = "Make your move, %c: "

func (p Player) ChooseSpace(out io.Writer, in io.Reader, board Board) int {
	reader := bufio.NewReader(in)
	fmt.Fprintf(out, PlayerMovePromptf, p.Token)
	input, readErr := reader.ReadString(byte('\n'))
	if readErr != nil {
		fmt.Fprintf(os.Stderr, "Player input encountered error %s", readErr.Error())
		fmt.Fprintln(out, "Oops! Let's try again...")
		return p.ChooseSpace(out, in, board)
	}
	choice, atoiErr := strconv.Atoi(strings.TrimSpace(input))
	if atoiErr != nil {
		fmt.Fprintln(out, "Choices must be one of the available spaces on the board. Let's try again.")
		return p.ChooseSpace(out, in, board)
	}
	return choice
}
