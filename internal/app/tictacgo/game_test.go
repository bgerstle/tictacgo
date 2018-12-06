package tictacgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockPlayer struct {
	mock.Mock
	PlayerInfo
}

func (mp MockPlayer) Info() PlayerInfo {
	return mp.PlayerInfo
}

func (mp MockPlayer) ChooseSpace(b Board) int {
	args := mp.Called(b)
	return args.Int(0)
}

func TestPlayChoosesMovesThenEnds(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	mockPlayer1 := MockPlayer{
		mock.Mock{},
		PlayerInfo{Token: x},
	}
	mockPlayer2 := MockPlayer{
		mock.Mock{},
		PlayerInfo{Token: o},
	}

	finalBoard := Board{
		spaces: []Space{
			X, X, X,
			O, O, X,
			O, X, O,
		},
	}

	for i, space := range finalBoard.spaces {
		var mockForToken *MockPlayer
		if *space == mockPlayer1.Token {
			mockForToken = &mockPlayer1
		} else {
			mockForToken = &mockPlayer2
		}
		mockForToken.On("ChooseSpace", mock.AnythingOfType("Board")).Return(i).Once()
	}

	g := Game{
		Player1: &mockPlayer1,
		Player2: &mockPlayer2,
		Board:   EmptyBoard(),
	}

	state, winner := g.Play()

	if state != Pending {
		t.Errorf("Expected game to be played to completion, but final state was: %s", state)
	}

	require.NotNil(winner)
	assert.Equal(mockPlayer1, winner)
	assert.Equal(finalBoard, g.Board)
}
