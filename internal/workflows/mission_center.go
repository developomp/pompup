package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "Mission Center",
		Desc: "Task Manager but Linux",
		Tags: []Tag{System, Gui},
		Setup: func() {
			wrapper.Flatpak("io.missioncenter.MissionCenter")
		},
	})
}
