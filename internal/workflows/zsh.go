package workflows

import (
	_ "embed"

	"github.com/developomp/pompup/internal/install"
	"github.com/pterm/pterm"
)

//go:embed assets/home/.zshrc
var zshConfig []byte

func init() {
	register(&Workflow{
		Name: "zsh",
		Desc: "Like bash but better",
		Tags: []Tag{System},
		Setup: func() {
			install.Paru("zsh")

			err := writeFile(inHome(".zshrc"), zshConfig)
			if err != nil {
				pterm.Fatal.Println("Failed to restore zsh config file:", err)
			}
		},
	})
}
