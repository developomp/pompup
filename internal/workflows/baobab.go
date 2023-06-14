package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "Baobab",
		Desc: "GNOME disk usage analyzer",
		Setup: func() {
			install.Flatpak("org.gnome.baobab")
		},
	})
}
