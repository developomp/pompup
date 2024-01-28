package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "Matrix",
		Desc: "Matrix communication GUI",
		Tags: []Tag{Gui},
		Setup: func() {
			install.Flatpak("im.riot.Riot")
		},
	})
}
