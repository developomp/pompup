package installers

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
)

//go:embed assets/home/.config/OpenTabletDriver/settings.json
var _otdSettings []byte

func init() {
	register(&Installer{
		Name: "OpenTabletDriver",
		Desc: "for osu",
		Tags: []Tag{Dev, Cli},
		Setup: func() {
			if !wrapper.IsArchPkgInstalled("opentabletdriver") {
				// https://opentabletdriver.net/Wiki/Install/Linux#arch
				wrapper.Paru("opentabletdriver")

				wrapper.Run("sudo", "mkinitcpio", "-P")
				wrapper.Run("sudo", "rmmod", "wacom")
				wrapper.Run("sudo", "rmmod", "hid_uclogic")
			}

			wrapper.WriteFile(wrapper.InHome(".config/OpenTabletDriver/settings.json"), _otdSettings)
		},
	})
}
