package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "GNOME Disk Utility",
		Desc: "Disk partitioning and stuff",
		Tags: []Tag{Gnome, Gui},
		Setup: func() {
			wrapper.ParuOnce("gnome-disk-utility")
		},
	})
}
