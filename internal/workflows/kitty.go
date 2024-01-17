package workflows

import (
	_ "embed"

	"github.com/developomp/pompup/internal/helper"
	"github.com/developomp/pompup/internal/install"
)

//go:embed assets/home/.config/kitty/kitty.conf
var kittyConfig []byte

func init() {
	register(&Workflow{
		Name: "Kitty",
		Desc: "terminal emulator",
		Tags: []Tag{Dev, Gui},
		Setup: func() {
			install.Paru("kitty")
			install.Paru("kitty-shell-integration")

			helper.WriteFile(helper.InHome(".config/kitty/kitty.conf"), kittyConfig)
		},
	})
}
