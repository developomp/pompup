package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "man DB",
		Desc: "Arch Linux manual",
		Setup: func() {
			wrapper.ParuOnce("man-db")
		},
	})
}
