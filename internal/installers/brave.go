package installers

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

//go:embed assets/home/.local/share/applications/com.brave.Browser.desktop
var _braveDesktopEntry string

func init() {
	register(&Installer{
		Name: "Brave Browser",
		Desc: "Least worst browser",
		Setup: func() {
			wrapper.FlatpakOnce("com.brave.Browser")

			var braveDesktopEntryPath = wrapper.InHome(".local/share/applications/com.brave.Browser.desktop")
			wrapper.SudoWriteFile(braveDesktopEntryPath, _braveDesktopEntry)

			err := wrapper.Run("sudo", "chmod", "+x", braveDesktopEntryPath)
			if err != nil {
				pterm.Fatal.Printfln("Failed to make %s executable: %s", braveDesktopEntryPath, err)
			}
		},
		Reminders: []string{
			"Enable Brave sync",
		},
	})
}
