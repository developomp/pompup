package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "GNOME Software",
		Desc: "GUI Flatpak Installer",
		Tags: []Tag{Gnome, Gui},
		Setup: func() {
			wrapper.Paru("gnome-software")
		},
	})
}
