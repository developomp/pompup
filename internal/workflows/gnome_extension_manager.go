package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "GNOME Extension Manager",
		Desc: "GNOME extensions without browser",
		Tags: []Tag{Gnome, Gui},
		Setup: func() {
			install.Flatpak("com.mattjakeman.ExtensionManager")
		},
	})
}
