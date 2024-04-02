package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Libre Menu Editor",
		Desc: "Application Menu Editor",
		Tags: []Tag{Gui, System, Configurator},
		Setup: func() {
			wrapper.FlatpakOnce("page.codeberg.libre_menu_editor.LibreMenuEditor")
		},
	})
}
