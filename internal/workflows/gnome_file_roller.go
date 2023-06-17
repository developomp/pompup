package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "GNOME File Roller",
		Desc: "Compression and Decompression",
		Tags: []Tag{Gnome, Gui},
		Setup: func() {
			install.Flatpak("org.gnome.FileRoller")
		},
	})
}
