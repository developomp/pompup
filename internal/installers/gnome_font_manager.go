package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "GNOME Font Manager",
		Desc: "Character viewing and stuff",
		Setup: func() {
			wrapper.FlatpakOnce("org.gnome.FontManager")
		},
	})
}
