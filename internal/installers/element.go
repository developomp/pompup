package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Matrix",
		Desc: "Matrix communication GUI",
		Setup: func() {
			wrapper.FlatpakOnce("im.riot.Riot")
		},
	})
}
