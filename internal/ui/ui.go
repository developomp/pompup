// ui contains code related to the visual part of the program.
package ui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

// Select displays a TUI that allows the user to select specific workflows from
// a list. It returns a boolean slice where true values indicate that the
// workflow at the index has been selected. It returns nil if the action was
// cancelled.
func Select() []bool {
	model := NewModel()
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		log.Fatalln(err)
	}

	if model.quit {
		return nil
	}

	return model.selected
}
