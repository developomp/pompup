package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "Dconf Editor",
		Desc: "GSettings editor for GNOME",
		Tags: []Tag{Gnome, Configurator},
		Setup: func() {
			wrapper.Flatpak("ca.desrt.dconf-editor")
		},
	})
}
