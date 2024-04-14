package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "atuin",
		Desc: "better shell history",
		Setup: func() {
			wrapper.ParuOnce("atuin")
		},
	})
}
