package main

import (
	"github.com/developomp/pompup/cmd"
	"github.com/pterm/pterm"
)

func main() {
	pterm.EnableDebugMessages()

	cmd.Execute()
}
