package tictacgo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minmaxUtils(t *testing.T) {
	mm := minimax{
		maxPlayer: x,
		minPlayer: o,
	}

	t.Run("opponent returns the other token", func(t *testing.T) {
		assert := assert.New(t)
		assert.Equal(x, mm.opponent(o))
		assert.Equal(o, mm.opponent(x))
	})

	t.Run("score modifier", func(t *testing.T) {
		assert := assert.New(t)
		assert.Equal(1.0, mm.scoreModifier(mm.maxPlayer))
		assert.Equal(-1.0, mm.scoreModifier(mm.minPlayer))
	})

	t.Run("minOrMax", func(t *testing.T) {
		assert := assert.New(t)
		scores := []score{
			score{
				spaceIndex: 0,
				value:      -1,
			},
			score{
				spaceIndex: 1,
				value:      1,
			},
		}
		assert.Equal(1, mm.minOrMax(mm.maxPlayer, scores).spaceIndex)
		assert.Equal(0, mm.minOrMax(mm.minPlayer, scores).spaceIndex)
	})
}

func Example_minimax_ChooseSpot() {
	board := Board{
		spaces: []Space{
			X, X, nil,
			O, O, nil,
			nil, nil, nil,
		},
	}

	mm := minimax{
		maxPlayer: x,
		minPlayer: o,
	}

	xsBest := mm.chooseSpot(x, board)
	osBest := mm.chooseSpot(o, board)

	fmt.Println(fmt.Sprintf("X's winning move in spot %d was chosen", xsBest))
	fmt.Println(fmt.Sprintf("O's winning move in spot %d was chosen", osBest))
	// Output:
	// X's winning move in spot 2 was chosen
	// O's winning move in spot 5 was chosen
}
