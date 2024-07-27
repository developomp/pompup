package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Installer{
		Name: "Bluetooth",
		Desc: "bluetooth",
		Setup: func() {
			// https://wiki.archlinux.org/title/Bluetooth

			wrapper.ParuOnce("bluez")
			wrapper.SystemctlEnable("bluetooth", wrapper.EnableNow)
		},
	})
}
