package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Mullvad Browser",
		Desc: "Tor browser without tor",
		Setup: func() {
			wrapper.FlatpakOnce("net.mullvad.MullvadBrowser")
		},
	})
}
