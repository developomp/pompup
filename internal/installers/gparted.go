package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "gparted",
		Desc: "GUI partition tool",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.Paru("gparted")
		},
	})
}
