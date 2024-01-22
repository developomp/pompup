package workflows

import (
	_ "embed"

	"github.com/developomp/pompup/internal/install"
)

//go:embed assets/dconf/gnome-nautilus.conf
var _gnomeNautilusDconf string

func init() {
	register(&Workflow{
		Name: "GNOME Files",
		Desc: "nautilus",
		Tags: []Tag{Gnome, Gui},
		Setup: func() {
			install.Paru("nautilus")
			install.Dconf(_gnomeNautilusDconf)
		},
	})
}
