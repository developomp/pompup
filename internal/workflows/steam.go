package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "Steam",
		Desc: "Game Downloader & Launcher",
		Tags: []Tag{Gui},
		Setup: func() {
			install.Flatpak("com.valvesoftware.Steam")
		},
	})
}
