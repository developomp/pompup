package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: ".NET",
		Desc: "Microsoft Java",
		Tags: []Tag{Dev, Cli},
		Setup: func() {
			wrapper.Paru("dotnet-sdk")
		},
	})
}
