package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "Discord",
		Desc: "discord and betterdiscord plugins",
		Tags: []Tag{Gui},
		Setup: func() {
			install.Flatpak("com.discordapp.Discord")
		},
		Reminders: []string{
			"Install vencord",
		},
	})
}
