package tictacgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPlayer struct {
	mock.Mock
	PlayerInfo
}

func (mp MockPlayer) Info() PlayerInfo {
	mp.Called()
	return mp.PlayerInfo
}

func (mp MockPlayer) ChooseSpace(board Board) int {
	args := mp.Called(board)
	return args.Int(0)
}

func TestPendingGameState(t *testing.T) {
	assert := assert.New(t)

	p1 := MockPlayer{
		PlayerInfo: PlayerInfo{Token: x},
	}
	p2 := MockPlayer{
		PlayerInfo: PlayerInfo{Token: o},
	}

	pendingBoards := []Board{
		EmptyBoard(),
		Board{
			spaces: []Space{
				X, O, nil,
				nil, nil, nil,
				nil, nil, nil,
			},
		},
	}

	for _, board := range pendingBoards {
		g := game{
			Player1: &p1,
			Player2: &p2,
			board:   board,
		}

		state, winner := g.state()

		assert.Equal(Pending, state)
		assert.Nil(winner)
	}
}

func TestVictoryGameState(t *testing.T) {
	assert := assert.New(t)

	p1 := MockPlayer{
		PlayerInfo: PlayerInfo{Token: x},
	}
	p2 := MockPlayer{
		PlayerInfo: PlayerInfo{Token: o},
	}

	victoryTestData := []struct {
		Board
		Player
	}{
		{
			Board{
				spaces: []Space{
					X, X, X,
					nil, O, O,
					nil, nil, nil,
				},
			},
			&p1,
		},
	}

	for _, boardAndPlayer := range victoryTestData {
		g := game{
			Player1: &p1,
			Player2: &p2,
			board:   boardAndPlayer.Board,
		}

		state, winner := g.state()

		if state != Victory {
			t.Fatalf(
				"Expected board to be evaluated as victory, got %s. Board: \n%s",
				state,
				boardAndPlayer.Board.String(),
			)
		}
		assert.Equal(boardAndPlayer.Player, winner)
	}
}
