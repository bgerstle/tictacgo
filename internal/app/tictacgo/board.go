package tictacgo

import (
	"fmt"
	"strconv"
	"strings"
)

// Space on a board, can either contain a run (e.g. 'X') or `nil` if the space is empty.
type Space *rune

// Board represents the board in a game of Tic Tac Toe.
type Board struct {
	// The spaces on the board. Not meant to be manipulated directly.
	spaces []Space
}

func spaceToString(space Space, fallback string) string {
	if space == nil {
		return fallback
	}
	return string(*space)
}

// EmptyBoard creates a new Board with empty spaces.
func EmptyBoard() Board {
	b := Board{}
	b.spaces = make([]Space, 9)
	return b
}

func (b Board) AvailableSpaces() (spaces []int) {
	for i, space := range b.spaces {
		if space == nil {
			spaces = append(spaces, i)
		}
	}
	return
}

func (b Board) SpacesCount() int {
	return len(b.spaces)
}

func (b Board) AssignSpace(index int, value Space) {
	maxIndex := b.SpacesCount() - 1
	if index < 0 || index > maxIndex {
		panic(fmt.Sprintf("Space index must be in bounds (0 < index < %d), got %d", maxIndex, index))
	}
	b.spaces[index] = value
}

// Get the spaces on the board partitioned into rows
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
