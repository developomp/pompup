package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Bustle",
		Desc: "D-bus activity viewer",
		Setup: func() {
			wrapper.FlatpakOnce("org.freedesktop.Bustle")
		},
	})
}
