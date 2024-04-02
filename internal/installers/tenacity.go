package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Tenacity",
		Desc: "Non-evil Adacity fork",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.FlatpakOnce("org.tenacityaudio.Tenacity")
		},
	})
}
