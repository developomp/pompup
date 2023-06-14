// ui contains code related to the visual part of the program.
package ui

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/developomp/pompup/internal/workflows"
)

// Select displays a TUI that allows the user to select specific workflows from
// a list. It returns a boolean slice where true values indicate that the
// workflow at the index has been selected. It returns nil if the action was
// cancelled.
func Select() []bool {
	model := NewModel()
	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalln(err)
	}

	if model.quit {
		return nil
	}

	return model.selected
}

// Model is bubbletea's (and elm's) way of saying states
type Model struct {
	quit     bool
	cursor   int
	selected []bool
}

func NewModel() Model {
	return Model{
		cursor:   0,
		selected: make([]bool, len(workflows.Workflows)),
	}
}

// Init is a function that returns an initial command for the bubbletea
// application to run.
func (m Model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

// Update is a function that handles incoming events and updates the model accordingly.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quit = true
			return m, tea.Quit

		case "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down":
			if m.cursor < len(workflows.Workflows)-1 {
				m.cursor++
			}

		case " ":
			// cycle between selected and not selected
			m.selected[m.cursor] = !m.selected[m.cursor]

		case "enter":
			return m, tea.Quit
		}
	}

	return m, nil
}

// View is a function that renders the UI based on the data in the model.
func (m Model) View() (s string) {
	for i, choice := range workflows.Workflows {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, checkbox(m.selected[i], choice.Name))
	}

	return s
}
