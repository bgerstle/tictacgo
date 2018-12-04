package tictacgo

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWelcome(t *testing.T) {
	assert := assert.New(t)

	bytes := &bytes.Buffer{}

	writeWelcomeErr := WriteWelcomeMessage(bytes)

	assert.Nil(writeWelcomeErr)

	assert.Equal(WelcomeMessage+"\n", bytes.String())
}
