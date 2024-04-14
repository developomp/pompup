package installers

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
)

//go:embed assets/home/.var/app/dev.vencord.Vesktop/config/vesktop/settings/settings.json
var _vesktop_settings []byte

//go:embed assets/home/.var/app/dev.vencord.Vesktop/config/vesktop/themes/nordic.theme.css
var _vesktop_theme []byte

//go:embed assets/home/.var/app/dev.vencord.Vesktop/config/vesktop/settings.json
var _vesktop_launcher_settings []byte

func init() {
	register(&Installer{
		Name: "Discord",
		Desc: "gamer chat platform",
		Setup: func() {
			wrapper.FlatpakOnce("dev.vencord.Vesktop")

			wrapper.WriteFile(wrapper.InHome(".var/app/dev.vencord.Vesktop/config/vesktop/settings/settings.json"), _vesktop_settings)
			wrapper.WriteFile(wrapper.InHome(".var/app/dev.vencord.Vesktop/config/vesktop/nordic.theme.css"), _vesktop_theme)
			wrapper.WriteFile(wrapper.InHome(".var/app/dev.vencord.Vesktop/config/vesktop/settings.json"), _vesktop_launcher_settings)
		},
	})
}
