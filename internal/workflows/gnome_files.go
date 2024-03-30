package workflows

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
)

//go:embed assets/dconf/gnome-nautilus.conf
var _gnomeNautilusDconf string

func init() {
	register(&Workflow{
		Name: "GNOME Files",
		Desc: "nautilus",
		Tags: []Tag{Gnome, Gui},
		Setup: func() {
			wrapper.Paru("nautilus")
			wrapper.Dconf(_gnomeNautilusDconf)

			wrapper.Paru("nautilus-open-any-terminal") // allow nautilus to open directory in terminal
			if !wrapper.IsBinInstalled("kitty") {
				setupKitty()
			}

			wrapper.Run("gsettings", "set", "com.github.stunkymonkey.nautilus-open-any-terminal", "terminal", "kitty")
			wrapper.Run("gsettings", "set", "com.github.stunkymonkey.nautilus-open-any-terminal", "keybindings", "''")
			wrapper.Run("gsettings", "set", "com.github.stunkymonkey.nautilus-open-any-terminal", "new-tab", "true")
		},
	})
}
