package workflows

import "github.com/developomp/pompup/internal/install"

func init() {
	register(&Workflow{
		Name: "gparted",
		Desc: "GUI partition tool",
		Tags: []Tag{Gui},
		Setup: func() {
			install.Paru("gparted")
		},
	})
}
