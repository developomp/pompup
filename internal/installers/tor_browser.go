package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Tor Browser",
		Desc: "whistleblower's favorite browser",
		Setup: func() {
			wrapper.FlatpakOnce("org.torproject.torbrowser-launcher")
		},
	})
}
