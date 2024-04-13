package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Golang",
		Desc: "Golang Tools",
		Setup: func() {
			wrapper.ParuOnce("go")
			wrapper.ParuOnce("go-tools")
		}})
}
