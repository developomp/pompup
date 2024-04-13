package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: ".NET",
		Desc: "Microsoft Java",
		Setup: func() {
			wrapper.ParuOnce("dotnet-sdk")
		},
	})
}
