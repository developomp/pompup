package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Tree",
		Desc: "Show directory structure",
		Setup: func() {
			wrapper.ParuOnce("tree")
		},
	})
}
