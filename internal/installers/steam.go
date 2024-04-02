package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Steam",
		Desc: "Game Downloader & Launcher",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.FlatpakOnce("com.valvesoftware.Steam")
		},
	})
}
