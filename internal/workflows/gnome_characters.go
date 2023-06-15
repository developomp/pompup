package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "GNOME characters",
		Desc: "Special character browser",
		Tags: []Tag{Gnome, App},
		Setup: func() {
			install.Flatpak("org.gnome.Characters")
		},
	})
}
