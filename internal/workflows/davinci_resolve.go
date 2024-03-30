package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "DaVinci Resolve",
		Desc: "Video Editing Tool",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.Paru("davinci-resolve")
		},
	})
}
