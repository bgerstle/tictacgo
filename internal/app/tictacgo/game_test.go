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

type MockGameReporter struct {
	mock.Mock
}

func (mgr MockGameReporter) ReportGameStart(b Board) {
	mgr.Called(b)
}

func (mgr MockGameReporter) ReportGameProgress(b Board, lastPlayerToken rune, lastPlayerSpace int) {
	mgr.Called(b, lastPlayerToken, lastPlayerSpace)
}

func (mgr MockGameReporter) ReportGameEnd(finalBoard Board, state GameState, winner Space) {
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

	g := Game{
		Player1: &mockPlayer1,
		Player2: &mockPlayer2,
		Board:   EmptyBoard(),
	}

	playerForX := g.PlayerForToken(x)
	require.NotNil(playerForX)
	assert.Equal(&mockPlayer1, *playerForX)

	playerForO := g.PlayerForToken(o)
	require.NotNil(playerForO)
	assert.Equal(&mockPlayer2, *playerForO)
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
	mockReporter := MockGameReporter{}

	finalBoard := Board{
		spaces: []Space{
			X, X, X,
			O, nil, nil,
			O, nil, nil,
		},
	}

	g := Game{
		Player1:  &mockPlayer1,
		Player2:  &mockPlayer2,
		Board:    EmptyBoard(),
		Reporter: &mockReporter,
	}

	mockReporter.On("ReportGameStart", g.Board).Return().Once()

	expectedBoard := g.Board
	expectedP1Moves := finalBoard.SpacesAssignedTo(mockPlayer1.Token)
	p1WriteIndex := 0
	expectedP2Moves := finalBoard.SpacesAssignedTo(mockPlayer2.Token)
	p2WriteIndex := 0
	for i := 0; i < len(expectedP1Moves)+len(expectedP2Moves); i++ {
		var (
			currentPlayer *MockPlayer
			space         int
		)
		if i%2 == 0 {
			currentPlayer = &mockPlayer1
			space = expectedP1Moves[p1WriteIndex]
			p1WriteIndex++
		} else {
			currentPlayer = &mockPlayer2
			space = expectedP2Moves[p2WriteIndex]
			p2WriteIndex++
		}

		token := (*currentPlayer).Token
		(*currentPlayer).On("ChooseSpace", expectedBoard).Return(space).Once()
		expectedBoard = expectedBoard.AssignSpace(space, &token)
		mockReporter.On("ReportGameProgress", expectedBoard, token, space).Return().Once()
	}

	mockReporter.On("ReportGameEnd", expectedBoard, Victory, X)

	state, winner := g.Play()

	if state != Victory {
		t.Errorf("Expected game to be played to completion, but final state was: %s", state)
	}

	require.NotNil(winner)
	assert.Equal(&g.Player1, winner)
	assert.Equal(finalBoard, g.Board)
	mockReporter.AssertExpectations(t)
}
