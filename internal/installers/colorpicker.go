package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Color picker",
		Desc: "picks color from screen",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.FlatpakOnce("nl.hjdskes.gcolor3")
		},
	})
}
