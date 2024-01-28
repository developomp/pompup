package workflows

import (
	_ "embed"

	"github.com/developomp/pompup/internal/helper"
	"github.com/developomp/pompup/internal/install"
)

//go:embed assets/home/.config/OpenTabletDriver/settings.json
var _otdSettings []byte

func init() {
	register(&Workflow{
		Name: "OpenTabletDriver",
		Desc: "for osu",
		Tags: []Tag{Dev, Cli},
		Setup: func() {
			// https://opentabletdriver.net/Wiki/Install/Linux#arch

			install.Paru("opentabletdriver")

			helper.Run("sudo", "mkinitcpio", "-P")
			helper.Run("sudo", "rmmod", "wacom")
			helper.Run("sudo", "rmmod", "hid_uclogic")

			helper.WriteFile(helper.InHome(".config/OpenTabletDriver/settings.json"), _otdSettings)
		},
	})
}
