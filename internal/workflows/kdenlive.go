package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "Kdenlive",
		Desc: "Adobe Premiere Pro but FOSS",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.Flatpak("org.kde.kdenlive")
		},
	})
}
