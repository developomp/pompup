package workflows

import (
	_ "embed"

	"github.com/developomp/pompup/internal/check"
	"github.com/developomp/pompup/internal/helper"
	"github.com/developomp/pompup/internal/install"
)

//go:embed assets/dconf/gnome-nautilus.conf
var _gnomeNautilusDconf string

func init() {
	register(&Workflow{
		Name: "GNOME Files",
		Desc: "nautilus",
		Tags: []Tag{Gnome, Gui},
		Setup: func() {
			install.Paru("nautilus")
			install.Dconf(_gnomeNautilusDconf)

			install.Paru("nautilus-open-any-terminal") // allow nautilus to open directory in terminal
			if !check.IsBinInstalled("kitty") {
				setupKitty()
			}

			helper.Run("gsettings", "set", "com.github.stunkymonkey.nautilus-open-any-terminal", "terminal", "kitty")
			helper.Run("gsettings", "set", "com.github.stunkymonkey.nautilus-open-any-terminal", "keybindings", "''")
			helper.Run("gsettings", "set", "com.github.stunkymonkey.nautilus-open-any-terminal", "new-tab", "true")
		},
	})
}
