package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Slack",
		Desc: "business communication",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.Flatpak("com.slack.Slack")
		},
	})
}
