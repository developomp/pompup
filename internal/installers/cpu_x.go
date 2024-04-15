package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "CPU-X",
		Desc: "CPU-Z for Linux",
		Setup: func() {
			wrapper.FlatpakOnce("io.github.thetumultuousunicornofdarkness.cpu-x")
		},
	})
}
