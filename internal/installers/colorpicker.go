package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Color picker",
		Desc: "picks color from screen",
		Setup: func() {
			wrapper.FlatpakOnce("nl.hjdskes.gcolor3")
		},
	})
}
