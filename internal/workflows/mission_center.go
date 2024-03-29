package workflows

import "github.com/developomp/pompup/internal/install"

func init() {
	register(&Workflow{
		Name: "Mission Center",
		Desc: "Task Manager but Linux",
		Tags: []Tag{System, Gui},
		Setup: func() {
			install.Flatpak("io.missioncenter.MissionCenter")
		},
	})
}
