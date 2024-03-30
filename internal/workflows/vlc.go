package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "VLC",
		Desc: "video player",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.Flatpak("org.videolan.VLC")
		},
	})
}
