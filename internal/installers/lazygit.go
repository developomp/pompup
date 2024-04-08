package installers

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Installer{
		Name: "lazygit",
		Desc: "lazy git tui",
		Tags: []Tag{Dev, Cli},
		Setup: func() {
			wrapper.ParuOnce("lazygit")
		},
	})
}
