package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Slack",
		Desc: "business communication",
		Setup: func() {
			wrapper.FlatpakOnce("com.slack.Slack")
		},
	})
}
