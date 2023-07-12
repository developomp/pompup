package workflows

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/developomp/pompup/internal/install"
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

			os.WriteFile(
				fmt.Sprintf("%s/.zshrc", getHomeDir()),
				zshConfig,
				DefaultPerm,
			)
		},
	})
}
