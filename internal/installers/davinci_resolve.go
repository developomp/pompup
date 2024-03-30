package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "DaVinci Resolve",
		Desc: "Video Editing Tool",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.Paru("davinci-resolve")
		},
	})
}
