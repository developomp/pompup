package installers

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
)

//go:embed assets/home/.gitconfig
var _gitconfig []byte

func init() {
	register(&Installer{
		Name: "git",
		Desc: "git gud",
		Tags: []Tag{Dev, Cli},
		Setup: func() {
			if wrapper.IsArchPkgInstalled("pacman", "git") {
				return
			}

			wrapper.Paru("git")
			wrapper.Paru("git-lfs")

			wrapper.WriteFile(wrapper.InHome(".gitconfig"), _gitconfig)
		},
	})
}
