package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Firefox",
		Desc: "off-brand chrome",
		Setup: func() {
			wrapper.FlatpakOnce("org.mozilla.firefox")
		},
	})
}
