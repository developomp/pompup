package workflows

import (
	"github.com/developomp/pompup/internal/helper"
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "Dconf Editor",
		Desc: "GSettings editor for GNOME",
		Tags: []Tag{Configurator},
		Setup: func() {
			install.Paru("openrazer-meta")
			helper.Run("sudo", "gpasswd", "-a", helper.GetUserName(), "plugdev")
			install.Paru("polychromatic")
		},
	})
}
