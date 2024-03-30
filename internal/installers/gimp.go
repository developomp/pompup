package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "GIMP",
		Desc: "Photoshop but FOSS",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.Flatpak("org.gimp.GIMP")
		},
	})
}
