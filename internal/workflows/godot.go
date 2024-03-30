package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "Godot",
		Desc: "FOSS game engine",
		Tags: []Tag{Dev, Gui},
		Setup: func() {
			wrapper.Flatpak("org.godotengine.Godot")
		},
	})
}
