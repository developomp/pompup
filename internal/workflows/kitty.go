package workflows

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
)

//go:embed assets/home/.config/kitty/kitty.conf
var kittyConfig []byte

func init() {
	register(&Workflow{
		Name:  "Kitty",
		Desc:  "terminal emulator",
		Tags:  []Tag{Dev, Gui},
		Setup: setupKitty,
	})
}

func setupKitty() {
	wrapper.Paru("kitty")
	wrapper.Paru("kitty-shell-integration")

	wrapper.WriteFile(wrapper.InHome(".config/kitty/kitty.conf"), kittyConfig)
}
