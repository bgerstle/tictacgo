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

// AvailableSpaces returns a list of indexes for empty spaces on the board.
func (b Board) AvailableSpaces() []int {
	availableSpaces := []int{}
	for i, space := range b.spaces {
		if space == nil {
			availableSpaces = append(availableSpaces, i)
		}
	}
	return availableSpaces
}

// SpacesLen returns the total number of spaces on the board (i.e. 9 for a 3x3).
func (b Board) SpacesLen() int {
	return len(b.spaces)
}

// IsSpaceAvailable returns whether or not the space at the specified index has a token.
func (b Board) IsSpaceAvailable(index int) bool {
	return b.spaces[index] == nil
}

// AssignSpace returns a new board with the chosen space assigned.
// A new board is returned in order to keep boards (publicly) immutable.
func (b Board) AssignSpace(index int, value Space) Board {
	if !b.IsSpaceAvailable(index) {
		panic(fmt.Sprintf("Space %d is already taken", index))
	}

	newBoard := EmptyBoard()
	copy(newBoard.spaces, b.spaces)
	newBoard.spaces[index] = value
	return newBoard
}

func (b Board) spaceVectorsForIndexVectors(indexVectors [][]int) [][]Space {
	spaceVectors := [][]Space{}
	for _, indexes := range indexVectors {
		spaceSet := []Space{}
		for _, index := range indexes {
			spaceSet = append(spaceSet, b.spaces[index])
		}
		spaceVectors = append(spaceVectors, spaceSet)
	}
	return spaceVectors
}

// Get the spaces on the board partitioned into rows
func (b Board) rows() [][]Space {
	return b.spaceVectorsForIndexVectors(b.rowIndexVectors())
}

func (b Board) rowIndexVectors() [][]int {
	rowIndexVectors := make([][]int, 3)
	for i := range b.spaces {
		rowNumber := i / 3
		row := rowIndexVectors[rowNumber]
		rowIndexVectors[rowNumber] = append(row, i)
	}
	return rowIndexVectors

}

func (b Board) columns() [][]Space {
	return b.spaceVectorsForIndexVectors(b.columnIndexVectors())
}

func (b Board) columnIndexVectors() [][]int {
	columnIndexVectors := make([][]int, 3)
	for i := range b.spaces {
		columnNumber := i % 3
		column := columnIndexVectors[columnNumber]
		columnIndexVectors[columnNumber] = append(column, i)
	}
	return columnIndexVectors
}

func (b Board) diagonals() [][]Space {
	return b.spaceVectorsForIndexVectors(b.diagonalIndexVectors())
}

func (b Board) diagonalIndexVectors() [][]int {
	return [][]int{
		[]int{
			0, 4, 8,
		},
		[]int{
			2, 4, 6,
		},
	}

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
