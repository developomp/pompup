package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Polychromatic",
		Desc: "openrazer frontend",
		Tags: []Tag{Configurator},
		Setup: func() {
			// install openrazer
			wrapper.ParuOnce("openrazer-daemon")
			wrapper.ParuOnce("openrazer-driver-dkms")
			wrapper.AddGroup("plugdev")

			// install polychromatic
			wrapper.ParuOnce("polychromatic")
		},
	})
}
