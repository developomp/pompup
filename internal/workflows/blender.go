package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "Blender",
		Desc: "FOSS 3D creation suite",
		Tags: []Tag{App},
		Setup: func() {
			install.Flatpak("org.blender.Blender")
		},
	})
}
