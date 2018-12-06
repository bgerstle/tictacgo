package tictacgo

import (
	"fmt"
	"io"
)

type ConsoleReporter struct {
	Out io.Writer
}

func (cr ConsoleReporter) ReportGameStart(b Board) {
	fmt.Fprint(cr.Out, b.String())
}
