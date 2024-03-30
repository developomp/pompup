package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "Brave Browser",
		Desc: "Least worst browser",
		Tags: []Tag{Gui},

		Setup: func() {
			wrapper.Flatpak("com.brave.Browser")
		},
		Reminders: []string{
			"Enable Brave sync",
		},
	})
}
