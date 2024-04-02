package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Flatseal",
		Desc: "flatpak permission manager",
		Tags: []Tag{Configurator},
		Setup: func() {
			wrapper.FlatpakOnce("com.github.tchx84.Flatseal")
		},
	})
}
