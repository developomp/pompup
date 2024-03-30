package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "GNOME Tweaks",
		Desc: "Complementary GNOME settings app",
		Tags: []Tag{Gnome, Configurator},
		Setup: func() {
			wrapper.Paru("gnome-tweaks")
		},
	})
}
