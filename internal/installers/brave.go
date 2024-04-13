package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Brave Browser",
		Desc: "Least worst browser",
		Setup: func() {
			wrapper.FlatpakOnce("com.brave.Browser")
		},
		Reminders: []string{
			"Enable Brave sync",
		},
	})
}
