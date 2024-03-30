package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "Dconf Editor",
		Desc: "GSettings editor for GNOME",
		Tags: []Tag{Configurator},
		Setup: func() {
			wrapper.Paru("openrazer-meta")
			wrapper.Run("sudo", "gpasswd", "-a", wrapper.GetUserName(), "plugdev")
			wrapper.Paru("polychromatic")
		},
	})
}
