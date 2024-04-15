package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Pacman-contrib",
		Desc: "pactree and other stuff",
		Setup: func() {
			wrapper.ParuOnce("pacman-contrib")
		},
	})
}
