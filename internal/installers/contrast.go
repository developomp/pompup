package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Contrast",
		Desc: "WCAG contrast requirements checker",
		Setup: func() {
			wrapper.FlatpakOnce("org.gnome.design.Contrast")
		},
	})
}
