package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "Eye Of Gnome (eog)",
		Desc: "GNOME image viewing utility",
		Tags: []Tag{Gnome, App},
		Setup: func() {
			install.Flatpak("org.gnome.eog")
		},
	})
}
