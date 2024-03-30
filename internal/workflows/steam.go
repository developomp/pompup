package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "Steam",
		Desc: "Game Downloader & Launcher",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.Flatpak("com.valvesoftware.Steam")
		},
	})
}
