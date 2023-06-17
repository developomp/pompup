package workflows

import "github.com/developomp/pompup/internal/install"

func init() {
	register(&Workflow{
		Name: "Inkscape",
		Desc: "Adobe Illustrator but FOSS",
		Tags: []Tag{Gui},
		Setup: func() {
			install.Flatpak("org.inkscape.Inkscape")
		},
	})
}
