package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Ventoy",
		Desc: "Bootable USB on steroid",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.ParuOnce("ventoy-bin")
		},
	})
}
