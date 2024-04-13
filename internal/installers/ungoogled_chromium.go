package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Ungoogled Chromium",
		Desc: "Chromium but ungoogled",
		Setup: func() {
			wrapper.FlatpakOnce("com.github.Eloston.UngoogledChromium")
		},
	})
}
