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
		PlayerInfo: PlayerInfo{Token: 'X'},
	}
	p2 := MockPlayer{
		PlayerInfo: PlayerInfo{Token: 'O'},
	}
	g := game{
		Player1: p1,
		Player2: p2,
		board:   EmptyBoard(),
	}

	state, winner := g.state()

	assert.Equal(Pending, state)
	assert.Nil(winner)
}
