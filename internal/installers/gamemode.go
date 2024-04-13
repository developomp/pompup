package installers

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
)

//go:embed assets/etc/gamemode.ini
var _gamemodeConfig string

func init() {
	register(&Installer{
		Name: "gamemode",
		Desc: "gamemoderun",
		Setup: func() {
			wrapper.ParuOnce("gamemode")
			wrapper.AddGroup("gamemode")
			wrapper.SudoWriteFile("/etc/gamemode.ini", _gamemodeConfig)
		},
	})
}
