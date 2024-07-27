package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Python",
		Desc: "python, pip, and stuff",
		Setup: func() {
			wrapper.ParuOnce("python")
			wrapper.ParuOnce("python-pip") // python package manager
			wrapper.ParuOnce("pyenv")
		},
	})
}
