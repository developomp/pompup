package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "btop",
		Desc: "top but better",
		Tags: []Tag{Cli},
		Setup: func() {
			wrapper.Paru("btop")
		},
	})
}
