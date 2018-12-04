package tictacgo

import (
	"strings"
)

type Board struct{}

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
