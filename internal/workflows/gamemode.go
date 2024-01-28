package workflows

import (
	_ "embed"

	"github.com/developomp/pompup/internal/helper"
	"github.com/developomp/pompup/internal/install"
	"github.com/pterm/pterm"
)

//go:embed assets/etc/gamemode.ini
var _gamemodeConfig string

func init() {
	register(&Workflow{
		Name:  "gamemode",
		Desc:  "gamemoderun",
		Tags:  []Tag{System},
		Setup: setupGamemode,
	})
}

func setupGamemode() {
	install.Paru("gamemode")

	helper.BashRun("sudo usermod -a -G gamemode $USER")

	err := helper.SudoWriteFile("/etc/gamemode.ini", _gamemodeConfig)
	if err != nil {
		pterm.Fatal.Println("Failed to write to /etc/gamemode.ini")
	}
}
