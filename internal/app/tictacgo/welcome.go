package tictacgo

import (
	"fmt"
	"io"
)

// WelcomeMessage contains the welcome message output by the app at startup.
const WelcomeMessage = "Welcome To Tic Tac Go!"

// WriteWelcomeMessage outputs the WelcomeMessage using the given writer.
func WriteWelcomeMessage(writer io.Writer) error {
	_, err := fmt.Fprintln(writer, WelcomeMessage)
	return err
}
