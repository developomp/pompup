package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "Inkscape",
		Desc: "Adobe Illustrator but FOSS",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.Flatpak("org.inkscape.Inkscape")
		},
	})
}
