package tictacgo

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoard(t *testing.T) {
	assert := assert.New(t)

	board := Board{}

	bytes := &bytes.Buffer{}

	writeError := board.Write(bytes)

	assert.Nil(writeError)

	assert.Equal(
		`0|1|2
=+=+=
3|4|5
=+=+=
6|7|8
`,
		bytes.String())
}
