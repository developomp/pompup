package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "GNOME Control Center",
		Desc: "GNOME settings app",
		Tags: []Tag{Gnome, Gui},
		Setup: func() {
			wrapper.ParuOnce("gnome-control-center")
		},
	})
}
