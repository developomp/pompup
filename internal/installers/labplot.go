package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Labplot",
		Desc: "data visualization and analysis tool",
		Setup: func() {
			wrapper.FlatpakOnce("org.kde.labplot2")
		},
	})
}
