package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "Brave Browser",
		Desc: "Least worst browser",
		Tags: []Tag{Gui},

		Setup: func() {
			install.Flatpak("com.brave.Browser")
		},
		Reminders: []string{
			"Enable Brave sync",
		},
	})
}
