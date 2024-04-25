package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "just",
		Desc: "just a command runner",
		Setup: func() {
			wrapper.ParuOnce("just")
		},
	})
}
