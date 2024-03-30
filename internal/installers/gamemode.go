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
		Name:  "gamemode",
		Desc:  "gamemoderun",
		Tags:  []Tag{System},
		Setup: setupGamemode,
	})
}

func setupGamemode() {
	wrapper.Paru("gamemode")

	wrapper.BashRun("sudo usermod -a -G gamemode $USER")

	err := wrapper.SudoWriteFile("/etc/gamemode.ini", _gamemodeConfig)
	if err != nil {
		pterm.Fatal.Println("Failed to write to /etc/gamemode.ini")
	}
}
