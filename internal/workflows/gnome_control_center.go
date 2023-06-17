package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "GNOME Control Center",
		Desc: "GNOME settings app",
		Tags: []Tag{Gnome, Gui},
		Setup: func() {
			install.Paru("gnome-control-center")
		},
	})
}
