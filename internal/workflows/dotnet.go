package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: ".NET",
		Desc: "Microsoft Java",
		Tags: []Tag{Dev, Cli},
		Setup: func() {
			wrapper.Paru("dotnet-sdk")
		},
	})
}
