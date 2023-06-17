package workflows

import "github.com/developomp/pompup/internal/install"

func init() {
	register(&Workflow{
		Name: "GIMP",
		Desc: "Photoshop but FOSS",
		Tags: []Tag{Gui},
		Setup: func() {
			install.Flatpak("org.gimp.GIMP")
		},
	})
}
