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

func (mp *MockPlayer) ChooseSpace(b Board) int {
	return mp.Called(b).Int(0)
}

type MockGameReporter struct {
	mock.Mock
}

func (mgr *MockGameReporter) ReportGameStart(b Board) {
	mgr.Called(b)
}

func (mgr *MockGameReporter) ReportGameProgress(b Board, lastPlayerToken rune, lastPlayerSpace int) {
	mgr.Called(b, lastPlayerToken, lastPlayerSpace)
}

func (mgr *MockGameReporter) ReportGameEnd(finalBoard Board, state GameState, winner Space) {
	mgr.Called(finalBoard, state, winner)
}

func TestGame_PlayerForToken(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	mockPlayer1 := MockPlayer{
		mock.Mock{},
		PlayerInfo{Token: x},
	}
	mockPlayer2 := MockPlayer{
		mock.Mock{},
		PlayerInfo{Token: o},
	}

	g := NewGame(&mockPlayer1, &mockPlayer2, nil)

	playerForX := g.PlayerForSpace(X)
	require.NotNil(playerForX)
	assert.Equal(&mockPlayer1, *playerForX)

	playerForO := g.PlayerForSpace(O)
	require.NotNil(playerForO)
	assert.Equal(&mockPlayer2, *playerForO)
}

func TestPlayChoosesMovesThenEnds(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	finalBoard := NewTestBoardWithSpaces([]Space{
		X, X, X,
		O, nil, nil,
		O, nil, nil,
	})

	mockPlayer1 := MockPlayer{
		mock.Mock{},
		PlayerInfo{Token: x},
	}
	mockPlayer2 := MockPlayer{
		mock.Mock{},
		PlayerInfo{Token: o},
	}
	mockReporter := MockGameReporter{}

	g := NewGame(&mockPlayer1, &mockPlayer2, &mockReporter)

	mockReporter.On("ReportGameStart", g.board).Return().Once()

	expectedSpaces := interleaveInts(finalBoard.SpacesAssignedTo(x), finalBoard.SpacesAssignedTo(o))

	expectedBoard := g.board
	for _, space := range expectedSpaces {
		var currentPlayer *MockPlayer
		if expectedBoard.ActivePlayerToken() == mockPlayer1.Token {
			currentPlayer = &mockPlayer1
		} else {
			currentPlayer = &mockPlayer2
		}
		(*currentPlayer).On("ChooseSpace", expectedBoard).Return(space).Once()
		expectedBoard, _, _ = expectedBoard.AssignSpace(space)
		mockReporter.On("ReportGameProgress", expectedBoard, currentPlayer.Token, space).Return().Once()
	}

	mockReporter.On("ReportGameEnd", finalBoard, Victory, X)

	state, winner := g.Play()

	if state != Victory {
		t.Errorf("Expected game to be played to completion, but final state was: %s", state)
	}

	require.NotNil(winner)
	assert.Equal(&g.Player1, winner)
	assert.Equal(finalBoard, g.board)
	mockReporter.AssertExpectations(t)
	mockPlayer1.AssertExpectations(t)
	mockPlayer2.AssertExpectations(t)
}
