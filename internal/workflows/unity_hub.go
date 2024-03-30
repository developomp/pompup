package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "Unity Hub",
		Desc: "Unity Installation manager",
		Tags: []Tag{Dev, Gui},
		Setup: func() {
			wrapper.Flatpak("com.unity.UnityHub")
		},
		Reminders: []string{
			"Change Unity Hub editors location",
		},
	})
}
