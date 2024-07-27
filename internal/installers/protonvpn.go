package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "ProtonVPN",
		Desc: "ProtonVPN",
		Setup: func() {
			// https://wiki.archlinux.org/title/ProtonVPN

			wrapper.ParuOnce("openvpn")
			wrapper.ParuOnce("wireguard-tools")
			wrapper.ParuOnce("protonvpn")
		},
	})
}
