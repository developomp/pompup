package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "Python",
		Desc: "python, pip, and stuff",
		Tags: []Tag{Dev},
		Setup: func() {
			wrapper.Paru("python")
			wrapper.Paru("python-pip") // python package manager
		},
	})
}
