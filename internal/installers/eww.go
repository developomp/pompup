package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "eww",
		Desc: "Linux Widgets",
		Tags: []Tag{System},
		Setup: func() {
			wrapper.ParuOnce("eww")
		},
	})
}
