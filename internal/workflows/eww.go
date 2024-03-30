package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "eww",
		Desc: "Linux Widgets",
		Tags: []Tag{System},
		Setup: func() {
			wrapper.Paru("eww-wayland")
		},
	})
}
