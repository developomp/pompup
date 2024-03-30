package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "Slack",
		Desc: "business communication",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.Flatpak("com.slack.Slack")
		},
	})
}
