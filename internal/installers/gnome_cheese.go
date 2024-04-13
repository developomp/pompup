package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "GNOME Cheese",
		Desc: "GNOME photo taking utility",
		Setup: func() {
			wrapper.FlatpakOnce("org.gnome.Cheese")
		},
	})
}
