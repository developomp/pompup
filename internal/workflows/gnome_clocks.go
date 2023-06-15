package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "GNOME Clocks",
		Desc: "GNOME Time management utility",
		Tags: []Tag{Gnome, App},
		Setup: func() {
			install.Flatpak("org.gnome.clocks")
		},
	})
}
