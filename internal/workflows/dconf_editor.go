package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "Dconf Editor",
		Desc: "GSettings editor for GNOME",
		Tags: []Tag{Gnome, Configurator},
		Setup: func() {
			install.Flatpak("ca.desrt.dconf-editor")
		},
	})
}
