package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "eza",
		Desc: "better ls",
		Setup: func() {
			wrapper.ParuOnce("eza")
		},
	})
}
