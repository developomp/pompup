package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Bun",
		Desc: "Bun > Deno > Nodejs",
		Setup: func() {
			wrapper.ParuOnce("bun-bin")
		},
	})
}
