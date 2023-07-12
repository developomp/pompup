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
		Tags: []Tag{System},
		Setup: func() {
			install.Paru("zsh")
			installConfig()
		},
	})
}

func installConfig() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		pterm.Fatal.Println(err)
	}

	configPath := fmt.Sprintf("%s/.zshrc", homeDir)

	os.WriteFile(configPath, zshConfig, DefaultPerm)
}
