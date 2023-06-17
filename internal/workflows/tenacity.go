package workflows

import "github.com/developomp/pompup/internal/install"

func init() {
	register(&Workflow{
		Name: "Tenacity",
		Desc: "Non-evil Adacity fork",
		Tags: []Tag{Gui},
		Setup: func() {
			install.Flatpak("org.tenacityaudio.Tenacity")
		},
	})
}
