package installers

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
)

//go:embed assets/home/.config/kitty/kitty.conf
var kittyConfig []byte

var kittyInstaller = Installer{
	Name: "Kitty",
	Desc: "terminal emulator",
	Tags: []Tag{Dev, Gui},
	Setup: func() {
		wrapper.ParuOnce("kitty")
		wrapper.ParuOnce("kitty-shell-integration")

		wrapper.WriteFile(wrapper.InHome(".config/kitty/kitty.conf"), kittyConfig)
	},
}

func init() {
	register(&kittyInstaller)
}
