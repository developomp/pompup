package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "GPA",
		Desc: "GnuPG GUI",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.FlatpakOnce("org.gnupg.GPA")
		},
	})
}
