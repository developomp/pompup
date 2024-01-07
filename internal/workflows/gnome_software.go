package workflows

import "github.com/developomp/pompup/internal/install"

func init() {
	register(&Workflow{
		Name: "GNOME Software",
		Desc: "GUI Flatpak Installer",
		Tags: []Tag{Gnome, Gui},
		Setup: func() {
			install.Paru("gnome-software")
		},
	})
}
