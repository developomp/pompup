package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "Matrix",
		Desc: "Matrix communication GUI",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.Flatpak("im.riot.Riot")
		},
	})
}
