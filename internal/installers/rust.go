package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Installer{
		Name: "Rust",
		Desc: "Rustup and stuff",
		Tags: []Tag{Cli, Dev},
		Setup: func() {
			wrapper.ParuOnce("rustup")

			wrapper.Run("rustup", "install", "stable")
			wrapper.Run("rustup", "default", "stable")
		},
	})
}
