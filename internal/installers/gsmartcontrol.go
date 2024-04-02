package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "gsmartcontrol",
		Desc: "GUI for S.M.A.R.T",
		Tags: []Tag{System},
		Setup: func() {
			wrapper.ParuOnce("gsmartcontrol")
		},
	})
}
