package workflows

import "github.com/developomp/pompup/internal/install"

func init() {
	register(&Workflow{
		Name: "gsmartcontrol",
		Desc: "GUI for S.M.A.R.T",
		Tags: []Tag{System},
		Setup: func() {
			install.Paru("gsmartcontrol")
		},
	})
}
