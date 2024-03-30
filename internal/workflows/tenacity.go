package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "Tenacity",
		Desc: "Non-evil Adacity fork",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.Flatpak("org.tenacityaudio.Tenacity")
		},
	})
}
