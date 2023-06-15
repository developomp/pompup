package workflows

import "github.com/developomp/pompup/internal/install"

func init() {
	register(&Workflow{
		Name: "LibreOffice",
		Desc: "FOSS Office Suite",
		Tags: []Tag{App},
		Setup: func() {
			install.Flatpak("org.libreoffice.LibreOffice")
		},
	})
}
