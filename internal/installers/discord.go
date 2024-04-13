package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Discord",
		Desc: "discord and betterdiscord plugins",
		Setup: func() {
			wrapper.FlatpakOnce("com.discordapp.Discord")
		},
		Reminders: []string{
			"Install vencord",
		},
	})
}
