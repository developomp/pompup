package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "GNOME File Roller",
		Desc: "Compression and Decompression",
		Tags: []Tag{Gnome, Gui},
		Setup: func() {
			wrapper.Flatpak("org.gnome.FileRoller")
		},
	})
}
