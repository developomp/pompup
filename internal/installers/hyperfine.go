package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "hyperfine",
		Desc: "benchmarking CLI",
		Setup: func() {
			wrapper.ParuOnce("hyperfine")
		},
	})
}
