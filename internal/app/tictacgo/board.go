package tictacgo

import (
	"strings"
)

type PlayerInfo struct {
	token rune
}

func (pi PlayerInfo) TokenStr() string {
	return string(pi.token)
}

type Space struct {
	token rune
}

type Board struct {
	spaces []*Space
}

func NewBoard() Board {
	b := Board{}
	b.spaces = make([]*Space, 9)
	for i := range b.spaces {
		b.spaces[i] = nil
	}
	return b
}

const rowSeparator = "===+===+==="
const spaceSeparator = "|"

func surroundWithSpaces(s string) string {
	return " " + s + " "
}

func (b Board) String() string {
	lines := []string{
		strings.Join([]string{
			surroundWithSpaces("0"),
			surroundWithSpaces("1"),
			surroundWithSpaces("2"),
		}, spaceSeparator),
		rowSeparator,
		strings.Join([]string{
			surroundWithSpaces("3"),
			surroundWithSpaces("4"),
			surroundWithSpaces("5"),
		}, spaceSeparator),
		rowSeparator,
		strings.Join([]string{
			surroundWithSpaces("6"),
			surroundWithSpaces("7"),
			surroundWithSpaces("8"),
		}, spaceSeparator),
		"",
	}
	return strings.Join(lines, "\n")
}
