package workflows

import "github.com/developomp/pompup/internal/install"

func init() {
	register(&Workflow{
		Name: "LibreOffice",
		Desc: "FOSS Office Suite",
		Setup: func() {
			install.Flatpak("org.libreoffice.LibreOffice")
		},
	})
}
