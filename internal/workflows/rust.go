package workflows

import (
	"os/exec"

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

			pterm.Info.Println("Installing stable rust toolchain")
			exec.Command("rustup", "install", "stable").Run()
		},
	})
}
