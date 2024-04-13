package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "GNOME Calculator",
		Desc: "GNOME Calculator",
		Setup: func() {
			wrapper.FlatpakOnce("org.gnome.Calculator")
		},
	})
}
