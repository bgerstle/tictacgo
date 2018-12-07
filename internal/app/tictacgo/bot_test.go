package tictacgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBotPlayer(t *testing.T) {
	assert := assert.New(t)

	bot := BotPlayer{
		PlayerInfo{Token: 'O'},
	}

	assert.Equal('O', bot.Info().Token)
	assert.Equal(0, bot.ChooseSpace(EmptyBoard()))
}
