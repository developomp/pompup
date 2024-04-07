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
			wrapper.ParuOnce("nvm")

			// nvm requires bash interpretation
			// install and use latest node version
			err := wrapper.BashRun("source ~/.bashrc && nvm install node --lts && nvm use node")
			if err != nil {
				pterm.Fatal.Println("Failed to install latest Node.js version:", err)
			}

			if !wrapper.IsBinInstalled("pnpm") {
				pterm.Debug.Println("Installing pnpm")
				wrapper.Run("npm", "install", "--global", "pnpm")
			}

			if !wrapper.IsBinInstalled("yarn") {
				pterm.Debug.Println("Installing yarn")
				wrapper.Run("npm", "install", "--global", "yarn")
			}
		},
	})
}
