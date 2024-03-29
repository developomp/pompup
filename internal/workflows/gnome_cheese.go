package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "GNOME Cheese",
		Desc: "GNOME photo taking utility",
		Tags: []Tag{Gnome, Gui},
		Setup: func() {
			install.Flatpak("org.gnome.Cheese")
		},
	})
}
