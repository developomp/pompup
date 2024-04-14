package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Collision",
		Desc: "File hash checker",
		Setup: func() {
			wrapper.FlatpakOnce("dev.geopjr.Collision")
		},
	})
}
