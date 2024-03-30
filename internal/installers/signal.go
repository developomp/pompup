package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Signal",
		Desc: "private messaging",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.Flatpak("org.signal.Signal")
		},
	})
}
