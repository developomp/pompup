package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "Slack",
		Desc: "business communication",
		Tags: []Tag{Gui},
		Setup: func() {
			install.Flatpak("com.slack.Slack")
		},
	})
}
