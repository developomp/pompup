package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "Godot",
		Desc: "FOSS game engine",
		Tags: []Tag{Dev, Gui},
		Setup: func() {
			install.Flatpak("org.godotengine.Godot")
		},
	})
}
