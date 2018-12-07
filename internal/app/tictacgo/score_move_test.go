package tictacgo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreMove(t *testing.T) {
	t.Run("number of scores equal to number of available spaces", func(t *testing.T) {
		assert := assert.New(t)
		board := EmptyBoard()
		assert.Len(board.ScoresFor(x), len(board.AvailableSpaces()))
	})
}

func ExampleScoreMove() {
	board := Board{
		spaces: []Space{
			X, X, nil,
			O, O, nil,
			nil, nil, nil,
		},
	}

	scores := board.ScoresFor(x)
	scores.Sort()
	maxScore := scores[len(scores)-1]

	fmt.Printf("X's winning move in spot %d has the highest score", maxScore.SpaceIndex)

	// Output:
	// X's winning move in spot 2 has the highest score
}
