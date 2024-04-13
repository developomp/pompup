package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Dconf Editor",
		Desc: "GSettings editor for GNOME",
		Setup: func() {
			wrapper.FlatpakOnce("ca.desrt.dconf-editor")
		},
	})
}
