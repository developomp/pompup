package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "DaVinci Resolve",
		Desc: "Video Editing Tool",
		Tags: []Tag{Gui},
		Setup: func() {
			install.Paru("davinci-resolve")
		},
	})
}
