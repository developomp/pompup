package workflows

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/developomp/pompup/internal/install"
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
		log.Fatalln(err)
	}

	os.WriteFile(fmt.Sprintf("%s/.zshrc", dirname), zshConfig, 0644)
}
