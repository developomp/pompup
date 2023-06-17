package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "GNOME Font Manager",
		Desc: "Character viewing and stuff",
		Tags: []Tag{Gnome, Gui},
		Setup: func() {
			install.Flatpak("org.gnome.FontManager")
		},
	})
}
