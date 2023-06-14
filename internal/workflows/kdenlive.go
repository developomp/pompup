package workflows

import "github.com/developomp/pompup/internal/install"

func init() {
	register(&Workflow{
		Name: "Kdenlive",
		Desc: "Adobe Premiere Pro but FOSS",
		Setup: func() {
			install.Flatpak("org.kde.kdenlive")
		},
	})
}
