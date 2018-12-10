package tictacgo

import (
	"strconv"
	"strings"
)

const rowSeparator = "===+===+==="
const spaceSeparator = "|"

func surroundWithSpaces(s string) string {
	return " " + s + " "
}

// Serialize the board into a string, suitable for writing to the console.
func (b Board) String() string {
	spaceStrs := make([]string, len(b.spaces))
	for i, space := range b.spaces {
		spaceStrs[i] = surroundWithSpaces(spaceToString(space, strconv.Itoa(i)))
	}

	rows := make([]string, 3)
	for i, token := range spaceStrs {
		rowNumber := i / 3
		row := rows[rowNumber]
		if row != "" {
			row += spaceSeparator
		}
		rows[rowNumber] = row + token
	}

	builder := strings.Builder{}

	for i, row := range rows {
		if i != 0 {
			builder.WriteString(rowSeparator)
			builder.WriteString("\n")
		}
		builder.WriteString(row)
		builder.WriteString("\n")
	}
	return builder.String()
}
