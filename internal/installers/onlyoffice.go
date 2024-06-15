package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "OnlyOffice",
		Desc: "FOSS Office Suite",
		Setup: func() {
			wrapper.FlatpakOnce("org.onlyoffice.desktopeditors") // cspell:disable-line
		},
	})
}
