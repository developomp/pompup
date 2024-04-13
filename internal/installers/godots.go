package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Godots",
		Desc: "Godot version manager",
		Setup: func() {
			wrapper.FlatpakOnce("io.github.MakovWait.Godots")
		},
		Reminders: []string{
			"Install Godot editors",
		},
	})
}
