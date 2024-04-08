package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Ungoogled Chromium",
		Desc: "Chromium but ungoogled",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.FlatpakOnce("com.github.Eloston.UngoogledChromium")
		},
	})
}