package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "bitwarden",
		Desc: "Password manager",
		Tags: []Tag{Gui, Cli},
		Setup: func() {
			install.Paru("bitwarden")
			install.Paru("bitwarden-cli")
		},
	})
}
