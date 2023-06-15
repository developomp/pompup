package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "Dconf Editor",
		Desc: "GSettings editor for GNOME",
		Setup: func() {
			install.Flatpak("ca.desrt.dconf-editor")
		},
	})
}
