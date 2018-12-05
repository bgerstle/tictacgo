package tictacgo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockChoiceProvider struct {
	mock.Mock
}

func (mock MockChoiceProvider) getChoice(p PlayerInfo, b Board) (int, error) {
	args := mock.Called(p, b)
	return args.Int(0), args.Error(1)
}

func TestHumanPlayer(t *testing.T) {
	t.Run("Returns value when no error occurred", func(t *testing.T) {
		assert := assert.New(t)

		board := EmptyBoard()

		for expectedChoice := range board.spaces {
			mockCP := MockChoiceProvider{}

			humanPlayer := humanPlayer{
				PlayerInfo:     PlayerInfo{Token: 'X'},
				choiceProvider: &mockCP,
			}

			mockCP.On("getChoice", humanPlayer.PlayerInfo, board).Return(expectedChoice, nil)

			actualChoice := humanPlayer.ChooseSpace(board)

			assert.Equal(expectedChoice, actualChoice)
		}
	})

	t.Run("Retries on error", func(t *testing.T) {
		assert := assert.New(t)

		board := EmptyBoard()

		mockCP := MockChoiceProvider{}

		humanPlayer := humanPlayer{
			PlayerInfo:     PlayerInfo{Token: 'X'},
			choiceProvider: &mockCP,
		}

		expectedChoice := 1

		mockCP.On("getChoice", humanPlayer.PlayerInfo, board).Return(-1, errors.New("test")).Once()
		mockCP.On("getChoice", humanPlayer.PlayerInfo, board).Return(expectedChoice, nil)

		actualChoice := humanPlayer.ChooseSpace(board)

		assert.Equal(expectedChoice, actualChoice)
	})
}