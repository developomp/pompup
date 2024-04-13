package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "GNOME Software",
		Desc: "GUI Flatpak Installer",
		Setup: func() {
			wrapper.ParuOnce("gnome-software")
		},
	})
}
