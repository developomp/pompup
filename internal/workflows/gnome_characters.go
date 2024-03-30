package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "GNOME characters",
		Desc: "Special character browser",
		Tags: []Tag{Gnome, Gui},
		Setup: func() {
			wrapper.Flatpak("org.gnome.Characters")
		},
	})
}
