package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "Signal",
		Desc: "private messaging",
		Tags: []Tag{Gui},
		Setup: func() {
			install.Flatpak("org.signal.Signal")
		},
	})
}
