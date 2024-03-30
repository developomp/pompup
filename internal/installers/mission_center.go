package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Mission Center",
		Desc: "Task Manager but Linux",
		Tags: []Tag{System, Gui},
		Setup: func() {
			wrapper.Flatpak("io.missioncenter.MissionCenter")
		},
	})
}
