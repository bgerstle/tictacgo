package tictacgo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockchoiceProvider struct {
	mock.Mock
}

func (mock *MockchoiceProvider) getChoice(p PlayerInfo, b Board) (int, error) {
	args := mock.Called(p, b)
	return args.Int(0), args.Error(1)
}

func TestHumanPlayer(t *testing.T) {
	t.Run("Returns value when no error occurred", func(t *testing.T) {
		assert := assert.New(t)

		board := NewEmptyTestBoard()

		for expectedChoice := range board.spaces {
			mockCP := MockchoiceProvider{}

			HumanPlayer := HumanPlayer{
				PlayerInfo:     PlayerInfo{Token: x},
				choiceProvider: &mockCP,
			}

			assert.Equal(x, HumanPlayer.Info().Token)

			mockCP.On("getChoice", HumanPlayer.PlayerInfo, board).Return(expectedChoice, nil)

			actualChoice := HumanPlayer.ChooseSpace(board)

			assert.Equal(expectedChoice, actualChoice)
		}
	})

	t.Run("Retries on error", func(t *testing.T) {
		assert := assert.New(t)

		board := NewEmptyTestBoard()

		mockCP := MockchoiceProvider{}

		HumanPlayer := HumanPlayer{
			PlayerInfo:     PlayerInfo{Token: 'X'},
			choiceProvider: &mockCP,
		}

		expectedChoice := 1

		mockCP.On("getChoice", HumanPlayer.PlayerInfo, board).Return(-1, errors.New("test")).Once()
		mockCP.On("getChoice", HumanPlayer.PlayerInfo, board).Return(expectedChoice, nil)

		actualChoice := HumanPlayer.ChooseSpace(board)

		assert.Equal(expectedChoice, actualChoice)
	})
}
