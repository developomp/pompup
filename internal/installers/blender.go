package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Installer{
		Name: "Blender",
		Desc: "FOSS 3D creation suite",
		Setup: func() {
			wrapper.FlatpakOnce("org.blender.Blender")
		},
	})
}
