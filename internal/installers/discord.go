package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Discord",
		Desc: "discord and betterdiscord plugins",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.Flatpak("com.discordapp.Discord")
		},
		Reminders: []string{
			"Install vencord",
		},
	})
}
