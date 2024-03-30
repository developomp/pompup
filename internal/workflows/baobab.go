package workflows

import (
	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Workflow{
		Name: "Baobab",
		Desc: "GNOME disk usage analyzer",
		Tags: []Tag{Gnome, Gui},
		Setup: func() {
			wrapper.Flatpak("org.gnome.baobab")
		},
	})
}
