package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Monero",
		Desc: "Bitcoin but privacy",
		Setup: func() {
			wrapper.ParuOnce("monero")
			wrapper.ParuOnce("monero-gui")
		},
	})
}
