package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "GNOME Font Manager",
		Desc: "Character viewing and stuff",
		Tags: []Tag{Gnome, Gui},
		Setup: func() {
			wrapper.Flatpak("org.gnome.FontManager")
		},
	})
}
