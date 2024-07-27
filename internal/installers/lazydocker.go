package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Installer{
		Name: "lazydocker",
		Desc: "docker git tui",
		Setup: func() {
			wrapper.ParuOnce("lazydocker-bin")
		},
	})
}
