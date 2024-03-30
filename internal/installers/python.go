package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Python",
		Desc: "python, pip, and stuff",
		Tags: []Tag{Dev},
		Setup: func() {
			wrapper.Paru("python")
			wrapper.Paru("python-pip") // python package manager
		},
	})
}
