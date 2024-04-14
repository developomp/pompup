package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "PikaBackup",
		Desc: "system backup utility",
		Setup: func() {
			wrapper.FlatpakOnce("org.gnome.World.PikaBackup")
		},
	})
}
