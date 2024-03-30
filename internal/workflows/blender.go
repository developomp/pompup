package workflows

import (
	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Workflow{
		Name: "Blender",
		Desc: "FOSS 3D creation suite",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.Flatpak("org.blender.Blender")
		},
	})
}
