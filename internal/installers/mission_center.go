package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Mission Center",
		Desc: "Task Manager but Linux",
		Setup: func() {
			wrapper.FlatpakOnce("io.missioncenter.MissionCenter")
		},
	})
}
