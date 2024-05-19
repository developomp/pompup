package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Zig",
		Desc: "somewhere between rust, go, and c",
		Setup: func() {
			wrapper.ParuOnce("zig")
			wrapper.ParuOnce("zls")
		},
	})
}
