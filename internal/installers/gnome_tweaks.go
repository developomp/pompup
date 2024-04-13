package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "GNOME Tweaks",
		Desc: "Complementary GNOME settings app",
		Setup: func() {
			wrapper.ParuOnce("gnome-tweaks")
		},
	})
}
