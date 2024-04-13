package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Tenacity",
		Desc: "Non-evil Adacity fork",
		Setup: func() {
			wrapper.FlatpakOnce("org.tenacityaudio.Tenacity")
		},
	})
}
