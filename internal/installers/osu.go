package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Installer{
		Name: "osu!",
		Desc: "osu!lazer and tablet driver",
		Setup: func() {
			if wrapper.IsAppImageInstalled("osu") {
				return
			}

			if !wrapper.PathExists(wrapper.InHome("Downloads/osu.AppImage")) {
				wrapper.DownloadFromGitHub("ppy", "osu", "osu\\.AppImage")
			}
		},
		Reminders: []string{
			"Install osu! via gearlever",
			"Install osu! skin from https://github.com/developomp/osu-pomp-skin",
		},
	})
}
