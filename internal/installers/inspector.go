package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Inspector",
		Desc: "PC information viewer",
		Setup: func() {
			wrapper.FlatpakOnce("io.github.nokse22.inspector")
		},
	})
}
