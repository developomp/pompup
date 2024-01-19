package workflows

import (
	"github.com/developomp/pompup/internal/helper"
	"github.com/developomp/pompup/internal/install"
	"github.com/pterm/pterm"
)

func init() {
	register(&Workflow{
		Name: "Rust",
		Desc: "Rustup and stuff",
		Tags: []Tag{Cli, Dev},
		Setup: func() {
			install.Paru("rustup")

			pterm.Debug.Println("Installing stable rust toolchain")
			helper.Run("rustup", "install", "stable")
		},
	})
}
