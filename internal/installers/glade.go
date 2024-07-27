package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Installer{
		Name: "Glade",
		Desc: "UI design tool for gtk applications",
		Setup: func() {
			wrapper.FlatpakOnce("org.gnome.Glade")
		},
	})
}
