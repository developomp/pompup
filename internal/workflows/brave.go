package workflows

import (
	"github.com/developomp/pompup/internal/install"
	"github.com/pterm/pterm"
)

func init() {
	register(&Workflow{
		Name: "Brave Browser",
		Desc: "Least worst browser",
		Tags: []Tag{Gui},
		PostSetup: func() {
			pterm.Info.Println("Enable Brave sync")
		},
		Setup: func() {
			install.Flatpak("com.brave.Browser")
		},
	})
}
