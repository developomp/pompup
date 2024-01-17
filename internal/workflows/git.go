package workflows

import (
	_ "embed"

	"github.com/developomp/pompup/internal/helper"
	"github.com/developomp/pompup/internal/install"
)

//go:embed assets/home/.gitconfig
var _gitconfig []byte

func init() {
	register(&Workflow{
		Name: "git",
		Desc: "git gud",
		Tags: []Tag{Dev, Cli},
		Setup: func() {
			install.Paru("git")
			install.Paru("git-lfs")

			helper.WriteFile(helper.InHome(".gitconfig"), _gitconfig)
		},
	})
}
