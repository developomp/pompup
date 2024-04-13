package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "OpenTofu",
		Desc: "terraform but not evil",
		Setup: func() {
			wrapper.ParuOnce("opentofu-bin")
		},
	})
}
