package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Kdenlive",
		Desc: "Adobe Premiere Pro but FOSS",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.Flatpak("org.kde.kdenlive")
		},
	})
}
