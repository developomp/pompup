package installers

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
)

//go:embed assets/home/.config/kitty/kitty.conf
var kittyConfig []byte

func init() {
	register(&Installer{
		Name: "Kitty",
		Desc: "terminal emulator",
		Setup: func() {
			wrapper.ParuOnce("kitty")
			wrapper.ParuOnce("kitty-shell-integration")

			wrapper.WriteFile(wrapper.InHome(".config/kitty/kitty.conf"), kittyConfig)
		},
	})
}
