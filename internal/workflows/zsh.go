package workflows

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/developomp/pompup/internal/install"
	"github.com/pterm/pterm"
)

//go:embed assets/home/.zshrc
var zshConfig []byte

func init() {
	register(&Workflow{
		Name: "zsh",
		Desc: "Like bash but better",
		Setup: func() {
			install.Paru("zsh")
			installConfig()
		},
	})
}

func installConfig() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		pterm.Fatal.Println(err)
		os.Exit(1)
	}

	os.WriteFile(fmt.Sprintf("%s/.zshrc", dirname), zshConfig, 0644)
}
