package workflows

import "github.com/developomp/pompup/internal/install"

func init() {
	register(&Workflow{
		Name: "btop",
		Desc: "top but better",
		Tags: []Tag{Cli},
		Setup: func() {
			install.Paru("btop")
		},
	})
}
