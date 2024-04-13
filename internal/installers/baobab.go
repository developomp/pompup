package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Installer{
		Name: "Baobab",
		Desc: "GNOME disk usage analyzer",
		Setup: func() {
			wrapper.FlatpakOnce("org.gnome.baobab")
		},
	})
}
