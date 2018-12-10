package tictacgo

import (
	"fmt"
)

// Space on a board, can either contain a run (e.g. 'X') or `nil` if the space is empty.
type Space *rune

func spaceToString(space Space, fallback string) string {
	if space == nil {
		return fallback
	}
	return string(*space)
}

// Board represents the board in a game of Tic Tac Toe.
type Board struct {
	// The spaces on the board. Not meant to be manipulated directly.
	spaces  []Space
	players [2]PlayerInfo
	turn    int
}

func NewBoard(players [2]PlayerInfo) Board {
	if players[0].Token == players[1].Token {
		panic("Can't create a board with players who share the same token.")
	}
	b := Board{}
	b.spaces = make([]Space, 9)
	b.players = players
	b.turn = 0
	return b
}

func (b Board) ActivePlayerToken() rune {
	index := b.turn % len(b.players)
	return b.players[index].Token
}

func (b Board) NextPlayerToken() rune {
	index := (b.turn + 1) % len(b.players)
	return b.players[index].Token
}

func spacesAssignedTo(t rune, spaces []Space) []int {
	tSpaces := []int{}
	for i, space := range spaces {
		if space != nil && *space == t {
			tSpaces = append(tSpaces, i)
		}
	}
	return tSpaces
}

// SpacesAssignedTo returns the indexes of spaces assigned to the given token.
func (b Board) SpacesAssignedTo(t rune) []int {
	return spacesAssignedTo(t, b.spaces)
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
func (b Board) AssignSpace(index int) (Board, GameState, Space) {
	currentState, _ := b.GameState()
	if currentState != Pending {
		panic(fmt.Sprintf("Can't assign spaces to a board in state '%s'", currentState))
	}
	if !b.IsSpaceAvailable(index) {
		panic(fmt.Sprintf("Space %d is already taken", index))
	}

	currentToken := b.ActivePlayerToken()

	newBoard := NewBoard(b.players)
	copy(newBoard.spaces, b.spaces)
	newBoard.spaces[index] = &currentToken
	newBoard.turn = b.turn + 1

	state, winner := newBoard.GameState()

	return newBoard, state, winner
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
