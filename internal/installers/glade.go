package installers

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Installer{
		Name: "Glade",
		Desc: "UI design tool for gtk applications",
		Tags: []Tag{Dev, Gui},
		Setup: func() {
			wrapper.FlatpakOnce("org.gnome.Glade")
		},
	})
}
