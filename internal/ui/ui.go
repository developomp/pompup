// ui contains code related to the visual part of the program.
package ui

import (
	"strconv"
	"strings"

	"atomicgo.dev/keyboard/keys"
	"github.com/developomp/pompup/internal/workflows"
	"github.com/pterm/pterm"
)

// Select displays a TUI that allows the user to select specific workflows from
// a list. It returns a boolean slice where true values indicate that the
// workflow at the index has been selected. It returns nil if the action was
// cancelled.
func Select() []bool {
	var options []string
	result := make([]bool, len(workflows.Workflows))

	for i, workflow := range workflows.Workflows {
		options = append(
			options,
			pterm.Sprintf(
				"%v. %v - %v %v",
				i,
				pterm.FgWhite.Sprint(workflow.Name),
				pterm.FgGray.Sprint(workflow.Desc),
				pterm.FgLightWhite.Sprint(workflow.Tags),
			),
		)
	}

	printer := pterm.DefaultInteractiveMultiselect.WithOptions(options)
	printer.MaxHeight = 10
	printer.DefaultText = `
  Format: X. Name - Description [Tags]

Please select your options`
	printer.Checkmark = &pterm.Checkmark{
		Checked:   pterm.Green("+"),
		Unchecked: " ",
	}
	printer.KeyConfirm = keys.Enter
	printer.KeySelect = keys.Space
	printer.Filter = false

	selectedOptions, _ := printer.Show()
	for _, selected := range selectedOptions {
		i, err := strconv.Atoi(selected[:strings.IndexByte(selected, '.')])
		if err != nil {
			pterm.Fatal.Println("Bad programming")
		}

		result[i] = true
	}

	return result
}
