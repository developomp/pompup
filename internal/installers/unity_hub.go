package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Unity Hub",
		Desc: "Unity Installation manager",
		Setup: func() {
			wrapper.FlatpakOnce("com.unity.UnityHub")
		},
		Reminders: []string{
			"Change Unity Hub editors location",
		},
	})
}
