package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "gparted",
		Desc: "GUI partition tool",
		Setup: func() {
			wrapper.ParuOnce("gparted")
		},
	})
}
