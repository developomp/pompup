package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "GNOME Cheese",
		Desc: "GNOME photo taking utility",
		Tags: []Tag{Gnome, Gui},
		Setup: func() {
			wrapper.Flatpak("org.gnome.Cheese")
		},
	})
}
