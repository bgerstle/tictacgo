package tictacgo

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Player struct {
	Token rune
}

func (p Player) ChooseSpace(out io.Writer, in io.Reader, board Board) int {
	fmt.Fprintf(out, "Make your move, %c: ", p.Token)
	reader := bufio.NewReader(in)
	input, readErr := reader.ReadString(byte('\n'))
	if readErr != nil {
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
