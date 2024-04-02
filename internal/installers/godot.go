package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Godot",
		Desc: "FOSS game engine",
		Tags: []Tag{Dev, Gui},
		Setup: func() {
			wrapper.FlatpakOnce("org.godotengine.Godot")
		},
	})
}
