package tictacgo

import (
	"strconv"
	"strings"
)

type Space *rune

type Board struct {
	spaces []Space
}

func spaceToString(space Space, fallback string) string {
	if space == nil {
		return fallback
	}
	return string(*space)
}

func EmptyBoard() Board {
	b := Board{}
	b.spaces = make([]Space, 9)
	return b
}

func (b Board) rows() [][]Space {
	rows := make([][]Space, 3)
	for i, space := range b.spaces {
		rowNumber := i / 3
		row := rows[rowNumber]
		rows[rowNumber] = append(row, space)
	}
	return rows
}

const rowSeparator = "===+===+==="
const spaceSeparator = "|"

func surroundWithSpaces(s string) string {
	return " " + s + " "
}

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
