package workflows

import "github.com/developomp/pompup/internal/install"

func init() {
	register(&Workflow{
		Name: "VLC",
		Desc: "video player",
		Tags: []Tag{Gui},
		Setup: func() {
			install.Flatpak("org.videolan.VLC")
		},
	})
}
