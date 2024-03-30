package workflows

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
)

//go:embed assets/home/.gitconfig
var _gitconfig []byte

func init() {
	register(&Workflow{
		Name: "git",
		Desc: "git gud",
		Tags: []Tag{Dev, Cli},
		Setup: func() {
			wrapper.Paru("git")
			wrapper.Paru("git-lfs")

			wrapper.WriteFile(wrapper.InHome(".gitconfig"), _gitconfig)
		},
	})
}
