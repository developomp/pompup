package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "LibreOffice",
		Desc: "FOSS Office Suite",
		Setup: func() {
			wrapper.FlatpakOnce("org.libreoffice.LibreOffice")
		},
	})
}
