package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "btop",
		Desc: "top but better",
		Tags: []Tag{Cli},
		Setup: func() {
			wrapper.Paru("btop")
		},
	})
}
