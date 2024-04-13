package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "GPA",
		Desc: "GnuPG GUI",
		Setup: func() {
			wrapper.FlatpakOnce("org.gnupg.GPA")
		},
	})
}
