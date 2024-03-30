package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "GIMP",
		Desc: "Photoshop but FOSS",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.Flatpak("org.gimp.GIMP")
		},
	})
}
