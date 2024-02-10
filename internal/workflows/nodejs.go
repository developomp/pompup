package workflows

import (
	"github.com/developomp/pompup/internal/helper"
	"github.com/developomp/pompup/internal/install"
	"github.com/pterm/pterm"
)

func init() {
	register(&Workflow{
		Name: "Node.JS",
		Desc: "Node.JS and related CLI tools",
		Tags: []Tag{Cli, Dev},
		Setup: func() {
			install.Paru("nvm")
			helper.Run("nvm install --lts")

			pterm.Debug.Println("Installing pnpm")
			helper.Run("npm", "install", "--global", "pnpm")

			pterm.Debug.Println("Installing yarn")
			helper.Run("npm", "install", "--global", "yarn")
		},
	})
}
