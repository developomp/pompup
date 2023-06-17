package workflows

import (
	"github.com/developomp/pompup/internal/install"
	"github.com/pterm/pterm"
)

func init() {
	register(&Workflow{
		Name: "Unity Hub",
		Desc: "Unity Installation manager",
		Tags: []Tag{Dev, Gui},
		PostSetup: func() {
			pterm.Info.Println("Change Unity Hub editors location")
		},
		Setup: func() {
			install.Flatpak("com.unity.UnityHub")
		},
	})
}
