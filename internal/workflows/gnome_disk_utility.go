package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "GNOME Disk Utility",
		Desc: "Disk partitioning and stuff",
		Tags: []Tag{Gnome, Gui},
		Setup: func() {
			install.Paru("gnome-disk-utility")
		},
	})
}
