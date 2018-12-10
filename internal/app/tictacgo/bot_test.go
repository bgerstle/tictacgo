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
}

func TestBotPlayer_AlwaysTies(t *testing.T) {
	assert := assert.New(t)

	for i := 0; i < 5; i++ {
		botX := BotPlayer{
			PlayerInfo{Token: 'X'},
		}
		botO := BotPlayer{
			PlayerInfo{Token: 'O'},
		}

		g := Game{
			botX,
			botO,
			nil,
			NewBoard([2]PlayerInfo{
				botX.Info(),
				botO.Info(),
			}),
		}

		state, winner := g.Play()

		if state != Tie {
			// bot shouldn't be beatable—even by itself!
			t.Fatalf(
				"Bot should not have been beatable, but result was: %s for %#v",
				state,
				winner)
		}
		assert.Nil(winner)
	}
}
