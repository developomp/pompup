package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Inkscape",
		Desc: "Adobe Illustrator but FOSS",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.FlatpakOnce("org.inkscape.Inkscape")
		},
	})
}
