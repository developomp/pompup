package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Eye Of Gnome (eog)",
		Desc: "GNOME image viewing utility",
		Tags: []Tag{Gnome, Gui},
		Setup: func() {
			wrapper.FlatpakOnce("org.gnome.eog")
		},
	})
}
