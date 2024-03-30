package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

func init() {
	register(&Installer{
		Name: "Node.JS",
		Desc: "Node.JS and related CLI tools",
		Tags: []Tag{Cli, Dev},
		Setup: func() {
			wrapper.Paru("nvm")
			wrapper.Run("nvm install --lts")

			pterm.Debug.Println("Installing pnpm")
			wrapper.Run("npm", "install", "--global", "pnpm")

			pterm.Debug.Println("Installing yarn")
			wrapper.Run("npm", "install", "--global", "yarn")
		},
	})
}
