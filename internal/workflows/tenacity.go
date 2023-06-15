package workflows

import "github.com/developomp/pompup/internal/install"

func init() {
	register(&Workflow{
		Name: "Tenacity",
		Desc: "Non-evil Adacity fork",
		Setup: func() {
			install.Flatpak("org.tenacityaudio.Tenacity")
		},
	})
}
