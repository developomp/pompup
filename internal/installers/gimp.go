package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "GIMP",
		Desc: "Photoshop but FOSS",
		Setup: func() {
			wrapper.FlatpakOnce("org.gimp.GIMP")
		},
	})
}
