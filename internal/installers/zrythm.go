package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Zrythm",
		Desc: "DAW",
		Setup: func() {
			wrapper.FlatpakOnce("org.zrythm.Zrythm")
		},
	})
}
