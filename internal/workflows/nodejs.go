package workflows

import (
	"os/exec"

	"github.com/developomp/pompup/internal/install"
	"github.com/pterm/pterm"
)

func init() {
	register(&Workflow{
		Name: "Node.JS",
		Desc: "Node.JS and related CLI tools",
		Tags: []Tag{Cli, Dev},
		Setup: func() {
			install.Paru("nodejs-lts-hydrogen") // Node.JS v18
			install.Paru("npm")

			pterm.Info.Println("Installing pnpm")
			exec.Command("npm", "install", "--global", "pnpm").Run()

			pterm.Info.Println("Installing yarn")
			exec.Command("npm", "install", "--global", "yarn").Run()
		},
	})
}
