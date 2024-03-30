package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "Libre Menu Editor",
		Desc: "Application Menu Editor",
		Tags: []Tag{Gui, System, Configurator},
		Setup: func() {
			wrapper.Flatpak("page.codeberg.libre_menu_editor.LibreMenuEditor")
		},
	})
}
