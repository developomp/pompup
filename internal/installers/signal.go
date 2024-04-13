package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Signal",
		Desc: "private messaging",
		Setup: func() {
			wrapper.FlatpakOnce("org.signal.Signal")
		},
	})
}
