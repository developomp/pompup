package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: ".NET",
		Desc: "Microsoft Java",
		Tags: []Tag{Dev, Cli},
		Setup: func() {
			install.Paru("dotnet-sdk")
		},
	})
}
