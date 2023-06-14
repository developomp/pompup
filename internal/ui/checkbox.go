package ui

import (
	"fmt"
)

func checkbox(checked bool, text string) string {
	checkMark := " "
	if checked {
		checkMark = "X"
	}

	return fmt.Sprintf("[%s] %s", checkMark, text)
}
