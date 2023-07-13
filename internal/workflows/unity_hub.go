package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "Unity Hub",
		Desc: "Unity Installation manager",
		Tags: []Tag{Dev, Gui},
		Setup: func() {
			install.Flatpak("com.unity.UnityHub")
		},
		Reminders: []string{
			"Change Unity Hub editors location",
		},
	})
}
