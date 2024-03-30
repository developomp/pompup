package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "Flatseal",
		Desc: "flatpak permission manager",
		Tags: []Tag{Configurator},
		Setup: func() {
			wrapper.Flatpak("com.github.tchx84.Flatseal")
		},
	})
}
