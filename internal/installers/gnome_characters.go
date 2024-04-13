package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "GNOME characters",
		Desc: "Special character browser",
		Setup: func() {
			wrapper.FlatpakOnce("org.gnome.Characters")
		},
	})
}
