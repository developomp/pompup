package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Dconf Editor",
		Desc: "GSettings editor for GNOME",
		Tags: []Tag{Configurator},
		Setup: func() {
			wrapper.ParuOnce("openrazer-meta")
			wrapper.Run("sudo", "gpasswd", "-a", wrapper.GetUserName(), "plugdev")
			wrapper.ParuOnce("polychromatic")
		},
	})
}
