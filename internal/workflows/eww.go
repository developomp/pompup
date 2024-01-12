package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "eww",
		Desc: "Linux Widgets",
		Tags: []Tag{System},
		Setup: func() {
			install.Paru("eww-wayland")
		},
	})
}
