package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Installer{
		Name: "Seabird",
		Desc: "Kubernetes IDE",
		Setup: func() {
			wrapper.FlatpakOnce("dev.skynomads.Seabird")
		},
	})
}
