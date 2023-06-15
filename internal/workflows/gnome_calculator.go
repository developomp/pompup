package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "GNOME Calculator",
		Desc: "GNOME Calculator",
		Tags: []Tag{Gnome, App},
		Setup: func() {
			install.Flatpak("org.gnome.Calculator")
		},
	})
}
