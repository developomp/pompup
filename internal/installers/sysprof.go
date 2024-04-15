package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Installer{
		Name: "Sysprof",
		Desc: "kernel-based performance profiler",
		Setup: func() {
			wrapper.ParuOnce("sysprof")
		},
	})
}
