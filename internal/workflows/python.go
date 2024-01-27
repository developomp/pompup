package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "Python",
		Desc: "python, pip, and stuff",
		Tags: []Tag{Dev},
		Setup: func() {
			install.Paru("python")
			install.Paru("python-pip") // python package manager
		},
	})
}
