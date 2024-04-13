package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Kdenlive",
		Desc: "Adobe Premiere Pro but FOSS",
		Setup: func() {
			wrapper.FlatpakOnce("org.kde.kdenlive")
		},
	})
}
