package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "GNOME Tweaks",
		Desc: "Complementary GNOME settings app",
		Tags: []Tag{Gnome, Configurator},
		Setup: func() {
			install.Paru("gnome-tweaks")
		},
	})
}
