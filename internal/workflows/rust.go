package workflows

import (
	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

func init() {
	register(&Workflow{
		Name: "Rust",
		Desc: "Rustup and stuff",
		Tags: []Tag{Cli, Dev},
		Setup: func() {
			wrapper.Paru("rustup")

			pterm.Debug.Println("Installing stable rust toolchain")
			wrapper.Run("rustup", "install", "stable")
		},
	})
}
