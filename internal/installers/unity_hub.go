package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Unity Hub",
		Desc: "Unity Installation manager",
		Tags: []Tag{Dev, Gui},
		Setup: func() {
			wrapper.FlatpakOnce("com.unity.UnityHub")
		},
		Reminders: []string{
			"Change Unity Hub editors location",
		},
	})
}
