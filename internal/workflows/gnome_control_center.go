package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "GNOME Control Center",
		Desc: "GNOME settings app",
		Tags: []Tag{Gnome, Gui},
		Setup: func() {
			wrapper.Paru("gnome-control-center")
		},
	})
}
