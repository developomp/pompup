package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Inkscape",
		Desc: "Adobe Illustrator but FOSS",
		Setup: func() {
			wrapper.FlatpakOnce("org.inkscape.Inkscape")
		},
	})
}
