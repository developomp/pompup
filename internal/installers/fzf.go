package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "fzf",
		Desc: "fuzzy finder cli",
		Setup: func() {
			wrapper.ParuOnce("fzf")
		},
	})
}
