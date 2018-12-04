package tictacgo

import (
	"fmt"
	"io"
	"strings"
)

type Board struct{}

const rowSeparator = "=+=+="
const spaceSeparator = "|"

func (b Board) Write(out io.Writer) error {
	lines := []string{
		strings.Join([]string{"0", "1", "2"}, spaceSeparator),
		rowSeparator,
		strings.Join([]string{"3", "4", "5"}, spaceSeparator),
		rowSeparator,
		strings.Join([]string{"6", "7", "8"}, spaceSeparator),
	}
	for _, line := range lines {
		_, error := fmt.Fprintln(out, line)
		if error != nil {
			return error
		}
	}
	return nil
}
