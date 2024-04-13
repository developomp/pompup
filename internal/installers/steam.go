package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Steam",
		Desc: "Game Downloader & Launcher",
		Setup: func() {
			wrapper.FlatpakOnce("com.valvesoftware.Steam")
		},
	})
}
