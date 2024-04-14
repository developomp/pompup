package installers

import (
	_ "embed"
	"strings"

	"github.com/developomp/pompup/internal/wrapper"
)

//go:embed assets/dconf/gnome-nautilus.conf
var _gnomeNautilusDconf string

//go:embed assets/home/.config/gtk-3.0/bookmarks
var _gnomeNautilusBookmarks string

func init() {
	register(&Installer{
		Name: "GNOME Files",
		Desc: "nautilus",
		Setup: func() {
			wrapper.ParuOnce("nautilus")
			wrapper.TryDconf(_gnomeNautilusDconf)

			_gnomeNautilusBookmarks = strings.ReplaceAll(_gnomeNautilusBookmarks, "$HOME", wrapper.GetHomeDir())
			wrapper.WriteFile(wrapper.InHome(".config/gtk-3.0/bookmarks"), []byte(_gnomeNautilusBookmarks))

			// allow nautilus to open directories in terminal
			wrapper.ParuOnce("nautilus-open-any-terminal")
			wrapper.Run("gsettings", "set", "com.github.stunkymonkey.nautilus-open-any-terminal", "terminal", "kitty")
			wrapper.Run("gsettings", "set", "com.github.stunkymonkey.nautilus-open-any-terminal", "keybindings", "''")
			wrapper.Run("gsettings", "set", "com.github.stunkymonkey.nautilus-open-any-terminal", "new-tab", "true")
		},
	})
}
