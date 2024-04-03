package installers

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

//go:embed assets/etc/gamemode.ini
var _gamemodeConfig string

func init() {
	register(&Installer{
		Name: "gamemode",
		Desc: "gamemoderun",
		Tags: []Tag{System},
		Setup: func() {
			wrapper.ParuOnce("gamemode")

			err := wrapper.AddGroup("gamemode")
			if err != nil {
				pterm.Fatal.Println("Failed to add user to gamemode group:", err)
			}

			wrapper.SudoWriteFile("/etc/gamemode.ini", _gamemodeConfig)
		},
	})
}
