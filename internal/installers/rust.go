package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Installer{
		Name:     "Rust",
		Desc:     "Rustup and stuff",
		Priority: -1,
		Setup: func() {
			wrapper.ParuOnce("rustup")

			wrapper.Run("rustup", "install", "stable")
			wrapper.Run("rustup", "default", "stable")
		},
	})
}
