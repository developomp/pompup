package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "LibreOffice",
		Desc: "FOSS Office Suite",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.Flatpak("org.libreoffice.LibreOffice")
		},
	})
}
