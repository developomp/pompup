package workflows

import "github.com/developomp/pompup/internal/install"

func init() {
	register(&Workflow{
		Name: "Inkscape",
		Desc: "Adobe Illustrator but FOSS",
		Setup: func() {
			install.Flatpak("org.inkscape.Inkscape")
		},
	})
}
